package storage

import (
	"context"

	"go.elastic.co/apm/v2"
)

type ROMemStorage struct {
	rates []ExchangeRate
}

func (r ROMemStorage) GetByCurrency(ctx context.Context, from string, to string) *ExchangeRate {
	span, ctx := apm.StartSpan(ctx, "GetByCurrency", "storage")
	defer span.End()

	for _, rate := range r.rates {
		if rate.From == from && rate.To == to {
			return &rate
		}
	}

	return nil
}

func NewROMemStorage(rates []ExchangeRate) ROMemStorage {
	return ROMemStorage{
		rates: rates,
	}
}
