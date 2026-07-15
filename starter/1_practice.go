package main

import "fmt"

func main() {
	fmt.Println("Receipt for Coffee")

	// TODO

	var coffeePrice float64 = 3.50
	var muffinPrice float64 = 2.75
	var teaPrice float64 = 2.80
	var quantity int

	var subtotal float64
	items := []string{"coffee", "muffin", "tea"}

	for i := 0; i < len(items); i++ {
		fmt.Println("How many " + items[i] + "s?: ")
		fmt.Scanln(&quantity)

		if items[i] == "coffee" {
			subtotal += coffeePrice * float64(quantity)
		}

		if items[i] == "muffin" {
			subtotal += muffinPrice * float64(quantity)
		}

		if items[i] == "tea" {
			subtotal += teaPrice * float64(quantity)
		}

	}

	fmt.Println("Your total is: ", subtotal)
}
