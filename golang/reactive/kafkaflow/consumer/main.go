package main

import (
	"github.com/Shopify/sarama"
	"github.com/enesanbar/workspace/golang/reactive/kafkaflow"
	flow "github.com/trustmaster/goflow"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("example", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	net := kafkaflow.NewUpperApp()

	in := make(chan string)
	net.SetInPort("In", in)

	wait := flow.Run(net)
	defer func() {
		close(in)
		<-wait
	}()

	for {
		msg := <-partitionConsumer.Messages()
		in <- string(msg.Value)
	}

}
