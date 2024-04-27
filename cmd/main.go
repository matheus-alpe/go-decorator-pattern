package main

import (
	"fmt"

	"github.com/matheus-alpe/go-decorator-pattern/internal/shop"
)

func main() {
	coffee := &shop.SimpleCoffee{}

	coffeeWithMilk := &shop.Milk{Coffee: coffee}
	fmt.Println(coffeeWithMilk.Description(), "-", coffeeWithMilk.Cost())

	coffeeWithCaramel := &shop.Caramel{Coffee: coffee}
	fmt.Println(coffeeWithCaramel.Description(), "-", coffeeWithCaramel.Cost())

	coffeeWithMilkAndCaramel := &shop.Milk{Coffee: &shop.Caramel{Coffee: coffee}}
	fmt.Println(coffeeWithMilkAndCaramel.Description(), "-", coffeeWithMilkAndCaramel.Cost())
}
