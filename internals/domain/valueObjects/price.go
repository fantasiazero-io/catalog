package valueObjects

type Price struct {
	Value    float32
	Currency Currency
}

func NewPrice(value float32, currency Currency) Price {
	return Price{
		Value:    value,
		Currency: currency,
	}
}

func (p Price) EqualTo(other Price) bool {
	return p.Value == other.Value && p.Currency == other.Currency
}

type Currency int

const (
	Barry Currency = iota + 1
)
