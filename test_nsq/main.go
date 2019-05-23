package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"runtime"
	"sync"
	"time"
)

//handler
type ConsumerT struct{}

var consumeMessageNumber int
var l sync.RWMutex
//处理消息
func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))

	l.Lock()
	defer l.Unlock()
	consumeMessageNumber ++
	return nil
}

var producer *nsq.Producer
var consumers []*nsq.Consumer

var conf *nsq.Config
var nsqdAddrTCP = "localhost:4150"    // tcp  used to publish topics and messages
var nsqdAddrHTTP = "localhost:4151"   // http publish topics,not exampled
var nsqAdminAddr = "localhost:4171"   // backend ui
var nsqlookupdAddr = "localhost:4161" // help consumer find topics
func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	//init a producer
	var er error
	conf = nsq.NewConfig()
	producer, er = nsq.NewProducer(nsqdAddrTCP, conf)
	if er != nil {
		log.Println(er.Error())
		return
	}
	er = producer.Ping()
	if er != nil {
		log.Println(er.Error())
		return
	}
	//
	////init consumer
	conf.LookupdPollInterval = 5 * time.Second
}
func main() {
	defer func(){
		producer.Stop()
		for i,_:=range consumers {
			consumers[i].Stop()
		}
	}()
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 创建100个消费者
	consumers = make([]*nsq.Consumer, 100)
	var e error
	for i, _ := range consumers {
		consumers[i], e = nsq.NewConsumer("go-nsq_testcase", "channel_1", conf)
		if e != nil {
			log.Println(e.Error())
			return
		}
		consumers[i].SetLogger(nil, 0)
		consumers[i].AddHandler(&ConsumerT{}) // 添加消费者接口

		if e = consumers[i].ConnectToNSQLookupd(nsqlookupdAddr);e!=nil {
			log.Println(e)
			return
		}

	}

	// 并发发布3000个消息
	for i := 0; i < 3000; i++ {
		go func(i int) {
			er := producer.Publish("go-nsq_testcase", []byte(fmt.Sprintf("hello,everyone_%d", i)))
			if er != nil {
				log.Println(er.Error())
				return
			}
		}(i)
	}

	time.Sleep(20 * time.Second)
	fmt.Println(consumeMessageNumber)
	select {}
}
