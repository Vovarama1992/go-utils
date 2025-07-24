package ctxutil

import (
	"context"
	"time"
)

func WithTimeout(ctx context.Context, seconds int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, time.Duration(seconds)*time.Second)
}
