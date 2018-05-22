package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
)

func main() {
	// 初始化配置
	config := sarama.NewConfig()
	// Ack
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 分区负载均衡
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true



	// 初始化生产者
	client, err := sarama.NewSyncProducer([]string{"124.16.84.192:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()
	i := 1
	// 发消息
	for   {
		// 生产者消息
		msg := &sarama.ProducerMessage{}
		// 消息内容
		msg.Topic = "nginx_log"
		msg.Value = sarama.StringEncoder("this is a good test, my message is good "+strconv.Itoa(i))


		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		i++
	}

}

