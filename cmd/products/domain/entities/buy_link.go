package entities

type BuyLink struct {
	url    string
	active bool
}

func (b *BuyLink) SetUrl(url string) {
	b.url = url
}

func (b *BuyLink) SetActive(active bool) {
	b.active = active
}

func (b *BuyLink) GetUrl() string {
	return b.url
}

func (b *BuyLink) GetActive() bool {
	return b.active
}

func NewBuyLink() *BuyLink {
	return &BuyLink{}
}
