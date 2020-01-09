package nsq

import (
	"fmt"
	nsq "github.com/nsqio/go-nsq"
)

func main() {
	var producer *nsq.Producer
	 _, _ = nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	_ = producer.Ping()
	fmt.Println()
}