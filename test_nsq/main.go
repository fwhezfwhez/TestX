package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"time"
	"fmt"
)
//handler
type ConsumerT struct{}
//处理消息
func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

var producer *nsq.Producer
var con1 *nsq.Consumer
var con2 *nsq.Consumer
var nsqdAddrTCP = "localhost:4150" //tcp  used to publish topics and messages
var nsqdAddrHTTP = "localhost:4151" //http
var nsqAdminAddr = "localhost:4171" //http
var nsqlookupdAddr = "localhost:4161" //http used to receive messages and topics by consumers
func init(){
	log.SetFlags(log.Llongfile|log.LstdFlags)
	//init a producer
	var er error
	config := nsq.NewConfig()
	producer,er =nsq.NewProducer(nsqdAddrTCP,config)
	if er!=nil{
		log.Println(er.Error())
		return
	}
	er=producer.Ping()
	if er!=nil{
		log.Println(er.Error())
		return
	}


	//init consumer
	config.LookupdPollInterval = 5*time.Second
	con1,er = nsq.NewConsumer("go-nsq_testcase","channel_1",config)
	if er!=nil{
		log.Println(er.Error())
		return
	}
	con1.SetLogger(nil, 0)        //屏蔽系统日志
	con1.AddHandler(&ConsumerT{}) // 添加消费者接口

	con2,er = nsq.NewConsumer("go-nsq_testcase","channel_1",config)
	if er!=nil{
		log.Println(er.Error())
		return
	}
	con2.SetLogger(nil, 0)        //屏蔽系统日志
	con2.AddHandler(&ConsumerT{}) // 添加消费者接口


}
func main() {
	go func(){
		if err := con1.ConnectToNSQLookupd(nsqlookupdAddr); err != nil {
			panic(err)
		}
		if err := con2.ConnectToNSQLookupd(nsqlookupdAddr); err != nil {
			panic(err)
		}
	}()	// publish topic and 2 messages
	er:=producer.Publish("go-nsq_testcase",[]byte("hello,everyone"))
	if er!=nil{
		log.Println(er.Error())
		return
	}
	er=producer.Publish("go-nsq_testcase",[]byte("hello,everyone2"))
	if er!=nil{
		log.Println(er.Error())
		return
	}

   time.Sleep(5*time.Second)

}
