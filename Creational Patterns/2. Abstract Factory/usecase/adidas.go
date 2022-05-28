package usecase

import "abstract_factory/domain"

type Adidas struct {
}

type AdidasShirt struct {
	domain.Shirt
}

type AdidasShoe struct {
	domain.Shoe
}

func (a *Adidas) MakeShoe() domain.IShoe {
	return &AdidasShoe{
		Shoe: domain.Shoe{
			Logo: "adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() domain.IShirt {
	return &AdidasShirt{
		Shirt: domain.Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}