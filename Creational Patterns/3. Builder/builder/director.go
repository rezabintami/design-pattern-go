package builder

import (
	"builder/domain"
	"builder/usecase"
)

type director struct {
	builder domain.IBuilder
}

func NewDirector(b domain.IBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) SetBuilder(b domain.IBuilder) {
	d.builder = b
}

func (d *director) BuildHouse() usecase.House {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}