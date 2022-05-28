package main

import (
	"builder/builder"
	"builder/domain"
	"fmt"
)

//! Important :
// Builder is a creational design pattern that lets you construct complex
// objects step by step. The pattern allows you to produce different types and
// representations of an object using the same construction code.

func main() {
    normalBuilder := domain.GetBuilder("normal")
    iglooBuilder := domain.GetBuilder("igloo")

    director := builder.NewDirector(normalBuilder)
    normalHouse := director.BuildHouse()

    fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
    fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
    fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

    director.SetBuilder(iglooBuilder)
    iglooHouse := director.BuildHouse()

    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)

}