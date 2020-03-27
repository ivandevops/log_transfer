package main

import (
	"strings"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

/*
var (
	wg sync.WaitGroup
)
*/

type KafkaClient struct {
	client sarama.Consumer
	addr   string
	topic  string
	wg     sync.WaitGroup
}

var (
	kafkaClient *KafkaClient
)

func initKafka(addr string, topic string) (err error) {

	kafkaClient = &KafkaClient{}

	consumer, err := sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	kafkaClient.client = consumer
	kafkaClient.addr = addr
	kafkaClient.topic = topic
	return
}
