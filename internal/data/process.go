package data

import (
	"context"

	"github.com/xorima/hub-sphere/internal/data/paginator"
)

func ProcessDoNothing[T any]() paginator.Process[T] {
	return func(ctx context.Context, item T) error {
		return nil
	}
}
