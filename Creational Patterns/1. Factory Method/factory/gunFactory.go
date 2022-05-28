package factory

import (
	"factory_method/domain"
	"factory_method/usecase"

	"fmt"
)

func GetGun(gunType string) (domain.IGun, error) {
    if gunType == "ak47" {
        return usecase.NewAk47(), nil
    }
    if gunType == "musket" {
        return usecase.NewMusket(), nil
    }
    return nil, fmt.Errorf("wrong gun type passed")
}