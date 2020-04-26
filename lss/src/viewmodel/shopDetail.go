package viewmodel

import (
	"fileserver/src/model"
)

type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

// NewShopDetail creates new instance for shop details.
func NewShopDetail(products []model.Product) ShopDetail {
	result := ShopDetail{
		Title:    "Lemonade Stand Supply",
		Active:   "shop",
		Products: []Product{},
	}
	for _, p := range products {
		result.Products = append(result.Products, productToVM(p))
	}
	return result
}
