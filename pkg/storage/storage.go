package storage

import "context"

type Storage interface {
	GetByCurrency(ctx context.Context, from string, to string) *ExchangeRate
}
