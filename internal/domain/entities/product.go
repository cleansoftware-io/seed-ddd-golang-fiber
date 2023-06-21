package entities

type Product struct {
	active      bool
	name        string
	price       Price
	description Description
	imageList   ImageList
	buyLinks    BuyLinks
}

func (p *Product) SetActive(active bool) {
	p.active = active
}

func (p *Product) SetName(name string) {
	p.name = name
}

func (p *Product) SetPrice(price Price) {
	p.price = price
}

func (p *Product) SetDescription(description Description) {
	p.description = description
}

func (p *Product) SetImageList(imageList ImageList) {
	p.imageList = imageList
}

func (p *Product) SetBuyLinks(buyLinks BuyLinks) {
	p.buyLinks = buyLinks
}

func (p *Product) GetActive() bool {
	return p.active
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetPrice() Price {
	return p.price
}

func (p *Product) GetDescription() Description {
	return p.description
}

func (p *Product) GetImageList() ImageList {
	return p.imageList
}

func (p *Product) GetBuyLinks() BuyLinks {
	return p.buyLinks
}

func NewProduct() *Product {
	return &Product{}
}
