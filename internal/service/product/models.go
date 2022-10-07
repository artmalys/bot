package product

import "fmt"

var allProducts = []Product{
	{Title: "Balloon"},
	{Title: "Postcard"},
	{Title: "Toy"},
	{Title: "Flower"},
}

type Product struct {
	id    int
	Title string
	Price float64
}

func (p *Product) TextPrice() string {
	return fmt.Sprintf("%.2f", p.Price)
}
