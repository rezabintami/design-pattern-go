package factory

import (
	"abstract_factory/domain"
	"abstract_factory/usecase"
	"fmt"
)

type iSportsFactory interface {
	MakeShoe() domain.IShoe
	MakeShirt() domain.IShirt
}

func GetSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &usecase.Adidas{}, nil
	}

	if brand == "nike" {
		return &usecase.Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}