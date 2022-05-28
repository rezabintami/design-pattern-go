package usecase

import "factory_method/domain"

type musket struct {
	gun
}

func NewMusket() domain.IGun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}