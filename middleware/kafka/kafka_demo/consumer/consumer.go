package consumer

import (
	"context"
	"github.com/Shopify/sarama"
	"kafka_demo/common"
	"log"
)

type EventConsumer struct {
	handler sarama.ConsumerGroupHandler
	group   sarama.ConsumerGroup
	topics  []string
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewConsumer(group string, topics []string, handler *EventHandler) *EventConsumer {
	gp := common.NewConsumerGroup(group)
	ctx, cancel := context.WithCancel(context.Background())
	return &EventConsumer{
		handler: handler,
		group:   gp,
		topics:  topics,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (e *EventConsumer) Consume() {
	for {
		select {
		case <-e.ctx.Done():
			e.group.Close()
			log.Println("EventConsumer ctx done")
			return
		default:
			if err := e.group.Consume(e.ctx, e.topics, e.handler); err != nil {
				log.Println("EventConsumer Consume failed err is ", err.Error())
			}
		}
	}
}

func (e *EventConsumer) Stop() {
	e.cancel()
}
