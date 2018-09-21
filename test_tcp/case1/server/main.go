package main

import (
	"fmt"
	"net"
	"os"
	"encoding/json"
	"bufio"
	"hash/crc32"
	"io"
)
//数据包的类型
const (
	HEART_BEAT_PACKET = 0x00
	REPORT_PACKET = 0x01
)

var (
	server = "127.0.0.1:8044"
)
//这里是包的结构体，其实是可以不需要的
type Packet struct {
	PacketType      byte
	PacketContent     []byte
}
//心跳包，这里用了json来序列化，也可以用github上的gogo/protobuf包
//具体见(https://github.com/gogo/protobuf)
type HeartPacket struct {
	Version     string`json:"version"`
	Timestamp   int64`json:"timestamp"`
}
//正式上传的数据包
type ReportPacket struct {
	Content   string`json:"content"`
	Rand         int`json:"rand"`
	Timestamp   int64`json:"timestamp"`
}
//与服务器相关的资源都放在这里面
type TcpServer struct {
	listener       *net.TCPListener
	hawkServer  *net.TCPAddr
}

func main() {
	//类似于初始化套接字，绑定端口
	hawkServer, err := net.ResolveTCPAddr("tcp", server)
	checkErr(err)
	//侦听
	listen, err := net.ListenTCP("tcp", hawkServer)
	checkErr(err)
	//记得关闭
	defer listen.Close()
	tcpServer := &TcpServer{
		listener:listen,
		hawkServer:hawkServer,
	}
	fmt.Println("start server successful......")
	//开始接收请求
	for {
		conn, err := tcpServer.listener.Accept()
		fmt.Println("accept tcp client %s",conn.RemoteAddr().String())
		checkErr(err)
		// 每次建立一个连接就放到单独的协程内做处理
		go Handle(conn)
	}
}
//处理函数，这是一个状态机
//根据数据包来做解析
//数据包的格式为|0xFF|0xFF|len(高)|len(低)|Data|CRC高16位|0xFF|0xFE
//其中len为data的长度，实际长度为len(高)*256+len(低)
//CRC为32位CRC，取了最高16位共2Bytes
//0xFF|0xFF和0xFF|0xFE类似于前导码
func Handle(conn net.Conn) {
	// close connection before exit
	defer conn.Close()
	// 状态机状态
	state := 0x00
	// 数据包长度
	length := uint16(0)
	// crc校验和
	crc16 := uint16(0)
	var recvBuffer []byte
	// 游标
	cursor := uint16(0)
	bufferReader := bufio.NewReader(conn)
	//状态机处理数据
	for {
		recvByte,err := bufferReader.ReadByte()
		if err != nil {
			//这里因为做了心跳，所以就没有加deadline时间，如果客户端断开连接
			//这里ReadByte方法返回一个io.EOF的错误，具体可考虑文档
			if err == io.EOF {
				fmt.Printf("client %s is close!\n",conn.RemoteAddr().String())
			}
			//在这里直接退出goroutine，关闭由defer操作完成
			return
		}
		//进入状态机，根据不同的状态来处理
		switch state {
		case 0x00:
			if recvByte == 0xFF {
				state = 0x01
				//初始化状态机
				recvBuffer = nil
				length = 0
				crc16 = 0
			}else{
				state = 0x00
			}
			break
		case 0x01:
			if recvByte == 0xFF {
				state = 0x02
			}else{
				state = 0x00
			}
			break
		case 0x02:
			length += uint16(recvByte) * 256
			state = 0x03
			break
		case 0x03:
			length += uint16(recvByte)
			// 一次申请缓存，初始化游标，准备读数据
			recvBuffer = make([]byte,length)
			cursor = 0
			state = 0x04
			break
		case 0x04:
			//不断地在这个状态下读数据，直到满足长度为止
			recvBuffer[cursor] = recvByte
			cursor++
			if(cursor == length){
				state = 0x05
			}
			break
		case 0x05:
			crc16 += uint16(recvByte) * 256
			state = 0x06
			break
		case 0x06:
			crc16 += uint16(recvByte)
			state = 0x07
			break
		case 0x07:
			if recvByte == 0xFF {
				state = 0x08
			}else{
				state = 0x00
			}
		case 0x08:
			if recvByte == 0xFE {
				//执行数据包校验
				if (crc32.ChecksumIEEE(recvBuffer) >> 16) & 0xFFFF == uint32(crc16) {
					var packet Packet
					//把拿到的数据反序列化出来
					json.Unmarshal(recvBuffer,&packet)
					//新开协程处理数据
					go processRecvData(&packet,conn)
				}else{
					fmt.Println("丢弃数据!")
				}
			}
			//状态机归位,接收下一个包
			state = 0x00
		}
	}
}

//在这里处理收到的包，就和一般的逻辑一样了，根据类型进行不同的处理，因人而异
//我这里处理了心跳和一个上报数据包
//服务器往客户端的数据包很简单地以\n换行结束了，偷了一个懒:)，正常情况下也可根据自己的协议来封装好
//然后在客户端写一个状态来处理
func processRecvData(packet *Packet,conn net.Conn)  {
	switch packet.PacketType {
	case HEART_BEAT_PACKET:
		var beatPacket HeartPacket
		json.Unmarshal(packet.PacketContent,&beatPacket)
		fmt.Printf("recieve heat beat from [%s] ,data is [%v]\n",conn.RemoteAddr().String(),beatPacket)
		conn.Write([]byte("heartBeat\n"))
		return
	case REPORT_PACKET:
		var reportPacket ReportPacket
		json.Unmarshal(packet.PacketContent,&reportPacket)
		fmt.Printf("recieve report data from [%s] ,data is [%v]\n",conn.RemoteAddr().String(),reportPacket)
		conn.Write([]byte("Report data has recive\n"))
		return
	}
}
//处理错误，根据实际情况选择这样处理，还是在函数调之后不同的地方不同处理
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//作者：getyouyou
//链接：https://www.jianshu.com/p/dbc62a879081
//來源：简书
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。