package entities

type Price struct {
	value    uint
	currency string
}

func (p *Price) SetValue(value uint) {
	p.value = value
}

func (p *Price) SetCurrency(currency string) {
	p.currency = currency
}

func (p *Price) GetValue() uint {
	return p.value
}

func (p *Price) GetCurrency() string {
	return p.currency
}

func NewPrice() *Price {
	return &Price{}
}
