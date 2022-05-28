package main

import (
	"abstract_factory/domain"
	"abstract_factory/factory"
	"fmt"
)

//! Important:
// Abstract Factory is a creational design pattern that lets you produce
// families of related objects without specifying their concrete classes

func main() {
    adidasFactory, _ := factory.GetSportsFactory("adidas")
    nikeFactory, _ := factory.GetSportsFactory("nike")

    nikeShoe := nikeFactory.MakeShoe()
    nikeShirt := nikeFactory.MakeShirt()

    adidasShoe := adidasFactory.MakeShoe()
    adidasShirt := adidasFactory.MakeShirt()

    printShoeDetails(nikeShoe)
    printShirtDetails(nikeShirt)

    printShoeDetails(adidasShoe)
    printShirtDetails(adidasShirt)
}

func printShoeDetails(s domain.IShoe) {
    fmt.Printf("Logo: %s", s.GetLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.GetSize())
    fmt.Println()
}

func printShirtDetails(s domain.IShirt) {
    fmt.Printf("Logo: %s", s.GetLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.GetSize())
    fmt.Println()
}