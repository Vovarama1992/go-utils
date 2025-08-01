package kafkautil

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaRetryBreaker struct {
	retry   *Retry
	breaker *Breaker
}

func NewKafkaRetryBreaker(retry *Retry, breaker *Breaker) *KafkaRetryBreaker {
	return &KafkaRetryBreaker{
		retry:   retry,
		breaker: breaker,
	}
}

// Обёртка: fn — это логика отправки сообщения
func (krb *KafkaRetryBreaker) Send(ctx context.Context, msg kafka.Message, fn func(context.Context, kafka.Message) error) error {
	return krb.retry.Do(func() error {
		return krb.breaker.Do(func() error {
			return fn(ctx, msg)
		})
	})
}
