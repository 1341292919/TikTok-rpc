package kafka

import (
	"TikTok-rpc/pkg/base/client"
	"TikTok-rpc/pkg/constants"
	"context"
	"errors"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/samber/lo"
	kafkago "github.com/segmentio/kafka-go"
	"io"
)

type Kafka struct {
	readers      []*kafkago.Reader
	writers      map[string]*kafkago.Writer
	consumeChans map[string]chan *Message
}

type Message struct {
	K, V []byte
}

func NewKafkaInstance() *Kafka {
	return &Kafka{
		readers:      make([]*kafkago.Reader, 0),
		writers:      make(map[string]*kafkago.Writer, 0),
		consumeChans: make(map[string]chan *Message),
	}
}

func (k *Kafka) SetWriter(topic string, asyncWrite ...bool) error {
	async := constants.DefaultKafkaProductorSyncWrite
	if asyncWrite != nil {
		async = asyncWrite[0]
	}

	w, err := client.GetNewWriter(topic, async)
	if err != nil {
		return err
	}

	k.writers[topic] = w
	return nil
}

// Send 发送消息到指定的 topic
func (k *Kafka) Send(ctx context.Context, topic string, messages []*Message) []error {
	if k.writers[topic] == nil {
		if err := k.SetWriter(topic); err != nil {
			return []error{err}
		}
	}

	return k.send(ctx, topic, messages)
}

func (k *Kafka) send(ctx context.Context, topic string, messages []*Message) []error {
	msgs := lo.Map(messages, func(item *Message, index int) kafkago.Message {
		return kafkago.Message{
			Key:   item.K,
			Value: item.V,
		}
	})

	err := k.writers[topic].WriteMessages(ctx, msgs...)
	switch e := err.(type) { //nolint
	case nil:
		return nil
	case kafkago.WriteErrors:
		return e
	default:
		return []error{err}
	}
}

func (k *Kafka) Consume(ctx context.Context, topic string, consumerNum int, groupID string, chanCap ...int) <-chan *Message {
	if k.consumeChans[topic] != nil {
		return k.consumeChans[topic]
	}
	chCap := constants.DefaultConsumerChanCap
	if chanCap != nil {
		chCap = chanCap[0]
	}
	ch := make(chan *Message, chCap)
	k.consumeChans[topic] = ch
	for i := 0; i < consumerNum; i++ {
		readers := client.GetNewReader(topic, groupID)
		k.readers = append(k.readers, readers)
		go k.consume(ctx, topic, readers)
	}
	return ch
}

func (k *Kafka) consume(ctx context.Context, topic string, r *kafkago.Reader) {
	ch := k.consumeChans[topic]
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			logger.Errorf("read message from kafka reader failed,err: %v", err.Error())
			return
		}

		ch <- &Message{K: msg.Key, V: msg.Value}
	}
}
