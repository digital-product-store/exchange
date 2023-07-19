package storage

type ExchangeRate struct {
	From string
	To   string
	Rate float32
}

func (e ExchangeRate) Total(amount float32) float32 {
	return amount * e.Rate
}
