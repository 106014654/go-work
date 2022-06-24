package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"time"
)

type KafkaConfigInfo struct {
	IsOpen     uint32 `xml:is_open`
	KafkaAddr  string `xml:kafka_addr`
	FlushSize  int    `xml:producer_flushsize`
	Zoookeeper string `xml:zookeeper`
}

type LogInfo struct {
	Time uint32
	Msg  string
}

type KafkaClient struct {
	Producer sarama.AsyncProducer
	Config   *KafkaConfigInfo
	Msg      chan *LogInfo
}

func NewKafkaClient() *KafkaClient {
	c := &KafkaClient{

	}

	if c.IsOpen() {
		config := sarama.NewConfig()
		config.Net.DialTimeout = 3 * time.Second
		config.Producer.Return.Successes = true
		config.Producer.Partitioner = sarama.NewManualPartitioner
		addr := c.Config.KafkaAddr

		producer, err := sarama.NewAsyncProducer(strings.Split(addr, ","), config)
		if err != nil {
			c.Config.IsOpen = 0
		} else {
			c.Producer = producer
		}
	}
	return c
}

func (this *KafkaClient) IsOpen() bool {
	if this.Config == nil {
		return false
	}
	return this.Config.IsOpen != 0
}

func (this *KafkaClient) write(info *LogInfo) {
	if this.IsOpen() {
		select {
		case this.Msg <- info:
			fmt.Print("kafka write success")
		default:
			fmt.Print("kafka write fail")
		}
	}
}

func (this *KafkaClient) Run() {
	if !this.IsOpen() {
		return
	}

	defer func() {
		if err := this.Producer.Close(); err != nil {
			fmt.Print("producer close error")
		}
	}()

	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Print("kafka producer state xxx")
		case info := <-this.Msg:
			fmt.Printf("%v", info)
		case err := <-this.Producer.Errors():
			fmt.Printf("kafka error:%v",err)
			return
		case <-this.Producer.Successes():
			fmt.Print("kafka producer success")
		}
	}
}
