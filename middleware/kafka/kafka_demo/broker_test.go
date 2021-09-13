package kafka_demo

import (
	"fmt"
	"kafka_demo/common"
	"testing"
)

func TestNewAsyncProducer(t *testing.T) {
	cli := common.NewAsyncProducer()
	fmt.Println(cli)
}
