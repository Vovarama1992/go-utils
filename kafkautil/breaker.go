// kafkautil/breaker.go
package kafkautil

import (
	"errors"
	"sync"
	"time"
)

var ErrBreakerOpen = errors.New("breaker is open")

type Breaker struct {
	mu           sync.Mutex
	maxFailures  int
	cooldown     time.Duration
	failureCount int
	lastFailTime time.Time
}

type BreakerConfig struct {
	MaxFailures int           // сколько ошибок подряд допустимо
	Cooldown    time.Duration // интервал для уменьшения failureCount
}

func NewBreaker(cfg BreakerConfig) *Breaker {
	return &Breaker{
		maxFailures: cfg.MaxFailures,
		cooldown:    cfg.Cooldown,
	}
}

func (b *Breaker) Do(fn func() error) error {
	b.mu.Lock()

	// Проверка: прошло ли достаточно времени, чтобы уменьшить счётчик
	if b.failureCount > 0 && time.Since(b.lastFailTime) > b.cooldown {
		b.failureCount-- // дать маленькую надежду
		b.lastFailTime = time.Now()
	}

	// Если достигли лимита фейлов — блокируем
	if b.failureCount >= b.maxFailures {
		b.mu.Unlock()
		return ErrBreakerOpen
	}

	b.mu.Unlock()

	// Выполняем функцию
	err := fn()

	b.mu.Lock()
	defer b.mu.Unlock()

	if err != nil {
		b.failureCount++
		b.lastFailTime = time.Now()
		return err
	}

	// Успех — сбрасываем счётчик
	b.failureCount = 0
	return nil
}
