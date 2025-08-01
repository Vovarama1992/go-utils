package kafkautil

import (
	"time"
)

type RetryConfig struct {
	Attempts int
	Delay    time.Duration
}

type Retry struct {
	attempts int
	delay    time.Duration
}

func NewRetry(cfg RetryConfig) *Retry {
	return &Retry{
		attempts: cfg.Attempts,
		delay:    cfg.Delay,
	}
}

func (r *Retry) Do(fn func() error) error {
	var err error
	for i := 0; i < r.attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		time.Sleep(r.delay)
	}
	return err
}
