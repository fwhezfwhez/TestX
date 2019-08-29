package main

import (
	"os"
	"fmt"
	"net"
	"time"
	"math/rand"
	"encoding/json"
	"bufio"
	"hash/crc32"
	//"sync"
)
//数据包类型
const (
	HEART_BEAT_PACKET = 0x00
	REPORT_PACKET = 0x01
)
//默认的服务器地址
var (
	server = "127.0.0.1:8044"
)
//数据包
type Packet struct {
	PacketType      byte
	PacketContent   []byte
}
//心跳包
type HeartPacket struct {
	Version     string`json:"version"`
	Timestamp   int64`json:"timestamp"`
}
//数据包
type ReportPacket struct {
	Content   string`json:"content"`
	Rand         int`json:"rand"`
	Timestamp   int64`json:"timestamp"`
}

//客户端对象
type TcpClient struct {
	connection     *net.TCPConn
	hawkServer  *net.TCPAddr
	stopChan       chan struct{}
}

func main()  {
	//拿到服务器地址信息
	hawkServer,err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Printf("hawk server [%s] resolve error: [%s]",server,err.Error())
		os.Exit(1)
	}
	//连接服务器
	connection,err := net.DialTCP("tcp",nil,hawkServer)
	if err != nil {
		fmt.Printf("connect to hawk server error: [%s]",err.Error())
		os.Exit(1)
	}
	client := &TcpClient{
		connection:connection,
		hawkServer:hawkServer,
		stopChan:make(chan struct{}),
	}
	//启动接收
	go client.receivePackets()

	//发送心跳的goroutine
	go func() {
		heartBeatTick := time.Tick(2 * time.Second)
		for{
			select {
			case <-heartBeatTick:
				client.sendHeartPacket()
			case <-client.stopChan:
				return
			}
		}
	}()

	//测试用的，开300个goroutine每秒发送一个包
	for i:=0;i<300;i++ {
		go func() {
			sendTimer := time.After(1 * time.Second)
			for{
				select {
				case <-sendTimer:
					client.sendReportPacket()
					sendTimer = time.After(1 * time.Second)
				case <-client.stopChan:
					return
				}
			}
		}()
	}
	//等待退出
	<-client.stopChan
}

// 接收数据包
func (client *TcpClient)receivePackets()  {
	reader := bufio.NewReader(client.connection)
	for {
		//承接上面说的服务器端的偷懒，我这里读也只是以\n为界限来读区分包
		msg, err := reader.ReadString('\n')
		if err != nil {
			//在这里也请处理如果服务器关闭时的异常
			close(client.stopChan)
			break
		}
		fmt.Print(msg)
	}
}
//发送数据包
//仔细看代码其实这里做了两次json的序列化，有一次其实是不需要的
func (client *TcpClient)sendReportPacket()  {
	reportPacket := ReportPacket{
		Content:getRandString(),
		Timestamp:time.Now().Unix(),
		Rand:rand.Int(),
	}
	packetBytes,err := json.Marshal(reportPacket)
	if err!=nil{
		fmt.Println(err.Error())
	}
	//这一次其实可以不需要，在封包的地方把类型和数据传进去即可
	packet := Packet{
		PacketType:REPORT_PACKET,
		PacketContent:packetBytes,
	}
	sendBytes,err := json.Marshal(packet)
	if err!=nil{
		fmt.Println(err.Error())
	}
	//发送
	client.connection.Write(EnPackSendData(sendBytes))
	fmt.Println("Send metric data success!")
}

//使用的协议与服务器端保持一致
func EnPackSendData(sendBytes []byte) []byte {
	packetLength := len(sendBytes) + 8
	result := make([]byte,packetLength)
	result[0] = 0xFF
	result[1] = 0xFF
	result[2] = byte(uint16(len(sendBytes)) >> 8)
	result[3] = byte(uint16(len(sendBytes)) & 0xFF)
	copy(result[4:],sendBytes)
	sendCrc := crc32.ChecksumIEEE(sendBytes)
	result[packetLength-4] = byte(sendCrc >> 24)
	result[packetLength-3] = byte(sendCrc >> 16 & 0xFF)
	result[packetLength-2] = 0xFF
	result[packetLength-1] = 0xFE
	fmt.Println(result)
	return result
}
//发送心跳包，与发送数据包一样
func (client *TcpClient)sendHeartPacket() {
	heartPacket := HeartPacket{
		Version:"1.0",
		Timestamp:time.Now().Unix(),
	}
	packetBytes,err := json.Marshal(heartPacket)
	if err!=nil{
		fmt.Println(err.Error())
	}
	packet := Packet{
		PacketType:HEART_BEAT_PACKET,
		PacketContent:packetBytes,
	}
	sendBytes,err := json.Marshal(packet)
	if err!=nil{
		fmt.Println(err.Error())
	}
	client.connection.Write(EnPackSendData(sendBytes))
	fmt.Println("Send heartbeat data success!")
}
//拿一串随机字符
func getRandString()string  {
	length := rand.Intn(50)
	strBytes := make([]byte,length)
	for i:=0;i<length;i++ {
		strBytes[i] = byte(rand.Intn(26) + 97)
	}
	return string(strBytes)
}

//作者：getyouyou
//链接：https://www.jianshu.com/p/dbc62a879081
//來源：简书
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。