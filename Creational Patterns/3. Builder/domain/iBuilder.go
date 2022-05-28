package domain

import "builder/usecase"

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() usecase.House
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return &usecase.NormalBuilder{}
	}

	if builderType == "igloo" {
		return &usecase.IglooBuilder{}
	}
	return nil
}