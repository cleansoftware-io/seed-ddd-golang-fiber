package entities

type ProductList struct {
	products []Product
}

func (p *ProductList) SetProducts(products []Product) {
	p.products = products
}

func (p *ProductList) GetProducts() []Product {
	return p.products
}

func NewProductList() *ProductList {
	return &ProductList{}
}

func (p *ProductList) AddProduct(product Product) {
	p.products = append(p.products, product)
}

func (p *ProductList) RemoveProductByName(product Product) {
	products := p.GetProducts()
	for i, p := range products {
		if p.GetName() == product.GetName() {
			products = append(products[:i], products[i+1:]...)
		}
	}
}
