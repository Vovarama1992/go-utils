package pgutil

import (
	"time"

	"github.com/sony/gobreaker"
)

type BreakerConfig struct {
	Name             string
	OpenTimeout      time.Duration
	FailureThreshold uint32
	MaxRequests      uint32
}

func NewBreaker(cfg BreakerConfig) *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        cfg.Name,
		MaxRequests: cfg.MaxRequests,
		Interval:    time.Minute,
		Timeout:     cfg.OpenTimeout,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= cfg.FailureThreshold
		},
	})
}
