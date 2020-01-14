package main

import (
	"fmt"
	nsq "github.com/nsqio/go-nsq"
	"log"
)

func sendMsg(producer *nsq.Producer)  {
	// 生产者写入nsq,10条消息，topic = "test"
	topic := "test"
	for i:=0; i< 10; i++ {
		msg := fmt.Sprintf("msg:%d",i)
		if producer != nil && msg != "" {	//不能发布空串，否则会导致error
			err := producer.Publish(topic, []byte(msg))
			if err != nil {
				log.Println("publish msg err", err)
			}
			fmt.Println(msg)
		}
	}
	fmt.Println("publish msg success")
}

func main() {
	//定义并初始化nsq生产者
	//producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	producer, err := nsq.NewProducer("127.0.0.1:5000", nsq.NewConfig())
	if err != nil {
		log.Println("initial producer failed", err)
	}
	err = producer.Ping() //Ping测试
	if err != nil {
		log.Println("ping nsq failed", err)
		//关闭生产者
		producer.Stop()
		producer = nil
	}
	fmt.Println("ping nsq success ")
	//生产者写入10条消息，
	sendMsg(producer)
}
