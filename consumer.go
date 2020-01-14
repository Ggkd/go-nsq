package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

// 消费者
type Consumer struct{}

//处理消息
func (c Consumer) HandleMessage(message *nsq.Message) error {
	fmt.Printf("addr: %s,  msg : %s\n", message.NSQDAddress, message.Body)
	return nil
}

func main() {
	//初始化消费者
	config := nsq.NewConfig()
	config.LookupdPollInterval = time.Second	//设置重连时间
	consumer, err := nsq.NewConsumer("test", "test-channel", config)
	if err != nil {
		fmt.Println("initial consumer err:", err)		//新建一个消费者
	}
	c := new(Consumer)
	consumer.SetLogger(nil, 0)		//屏蔽系统日志
	consumer.AddHandler(c)		//添加消费者接口

	//建立nsqlookupd连接
	//err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	//if err != nil {
	//	fmt.Println("conn nsqlookupd err : ", err)
	//}

	//建立一个nsqd连接
	//err = consumer.ConnectToNSQD("127.0.0.1:4150")
	//if err != nil {
	//	fmt.Println("conn nsqd err :", err)
	//}

	//建立多个nsqd连接
	err = consumer.ConnectToNSQDs([]string{"127.0.0.1:4000", "127.0.0.1:5000"})
	if err != nil {
		fmt.Println("conn nsqd err :", err)
	}

	for {
		time.Sleep(time.Second)
	}
}