package main

import (
	"factory_method/domain"
	"factory_method/factory"
	"fmt"
)

//! Important :
// Many designs start by using Factory Method
// (less complicated and more customizable via subclasses)
// and evolve toward Abstract Factory, Prototype, or
// Builder (more flexible, but more complicated).

func main() {
    ak47, _ := factory.GetGun("ak47")
    musket, _ := factory.GetGun("musket")

    printDetails(ak47)
    printDetails(musket)
}

func printDetails(g domain.IGun) {
    fmt.Printf("Gun: %s", g.GetName())
    fmt.Println()
    fmt.Printf("Power: %d", g.GetPower())
    fmt.Println()
}