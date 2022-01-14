package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", "wangyiwei", "123", "192.168.83.150", 5672)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channel.Close()
	q, err := channel.QueueDeclare("test", false, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = channel.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("msg001"),
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("成功")

	//开始消费
	channel2, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channel2.Close()
	msgs, err := channel2.Consume(q.Name, "c1", false, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
	for msg := range msgs {
		fmt.Println(msg.DeliveryTag, string(msg.Body))
	}
}
