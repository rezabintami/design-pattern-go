package usecase

import "factory_method/domain"

type ak47 struct {
	gun
}

func NewAk47() domain.IGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}