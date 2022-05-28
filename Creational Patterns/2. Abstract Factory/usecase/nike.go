package usecase

import "abstract_factory/domain"

type Nike struct {
}

type NikeShirt struct {
	domain.Shirt
}

type NikeShoe struct {
	domain.Shoe
}

func (n *Nike) MakeShoe() domain.IShoe {
	return &NikeShoe{
		Shoe: domain.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShirt() domain.IShirt {
	return &NikeShirt{
		Shirt: domain.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}