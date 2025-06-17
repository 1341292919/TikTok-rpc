package mq

import (
	"TikTok-rpc/pkg/kafka"
	"sync/atomic"
)

type KafkaAdapter struct {
	mq   *kafka.Kafka
	done atomic.Bool // topic并发的标记
}

func NewKafkaAdapter(mq *kafka.Kafka) *KafkaAdapter {
	return &KafkaAdapter{
		mq: mq,
	}
}
