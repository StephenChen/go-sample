package kafka_demo

import (
	"kafka_demo/common"
	"kafka_demo/consumer"
	"kafka_demo/producer"
)

func main() {
	topic := "chen_kafka_test"
	consm := consumer.NewConsumer(common.Group, []string{topic}, &consumer.EventHandler{})
	defer consm.Stop()
	go consm.Consume() // 异步消费

	pro := producer.NewEventProducer(topic)

	for i := 0; i < 100; i++ {
		msg := &common.KafkaMsg{
			ID:     uint64(i),
			Detail: "chen jiushiwan",
		}
		pro.Producer(msg)
	}
	select {}
}
