package mq

import (
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/kafka"
	"context"
)

func (ka *KafkaAdapter) ConsumeLikeMessage(ctx context.Context) <-chan *kafka.Message {
	msgCh := ka.mq.Consume(ctx,
		constants.KafkaInteractLikeTopic,
		constants.KafkaInteractConsumerNum,
		constants.KafkaInteractLikeGroupId,
		constants.DefaultConsumerChanCap)
	return msgCh
}

func (ka *KafkaAdapter) ConsumeCommentMessage(ctx context.Context) <-chan *kafka.Message {
	msgCh := ka.mq.Consume(ctx,
		constants.KafkaInteractCommentTopic,
		constants.KafkaInteractConsumerNum,
		constants.KafkaInteractLikeGroupId,
		constants.DefaultConsumerChanCap)
	return msgCh
}
