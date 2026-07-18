package main

import "fmt"

type Ticket struct {
	Movie string
	Price float64
	Seat  string
}

type Customer struct {
	Name          string
	WalletBalance float64
	Ticket        []string
}

type Snack struct {
	Name string
	Price float64
}

func purchaseTicket(customer *Customer, ticket *Ticket) {
	customer.WalletBalance = customer.WalletBalance - ticket.Price
	customer.Ticket = append(customer.Ticket, ticket.Movie, ticket.Seat)
}

func applyDiscount(ticket *Ticket) {
	ticket.Price = ticket.Price * 0.8
}

// method
func (ticket *Ticket) ApplyDiscount() {
	ticket.Price = ticket.Price * 0.8
}


func GetPrice(t Purchasable) float64 {
	return t.GetPrice() 
} 

// 1. interface
type Purchasable interface {
	GetPrice() float64
}

// 2. Add matching methods to structs
func (t Ticket) GetPrice() float64 {
	return t.Price 
}

func (s Snack) GetPrice() float64 {
	return s.Price 
}

// 3. Use the interface in a function 
func ShowPrice(p Purchasable) {
	fmt.Println(p.GetPrice())
}


// func (t Ticket) GetPrice() float64
// func (s Snack) GetPrice() float64

func main() {

	ticket := Ticket{
		Movie: "Dune",
		Price: 20.00,
		Seat:  "B12",
	}

	// customer := Customer{
	// 	Name:          "Rolex",
	// 	WalletBalance: 50,
	// }

	snack := Snack{
		Name: "Popcorn",
		Price: 8.00,
	}

	// fmt.Println("Before discount: ")

	// fmt.Println("Movie: ", ticket.Movie)
	// fmt.Println("Price: ", ticket.Price)
	// fmt.Println("Seat: ", ticket.Seat)

	// applyDiscount(&ticket)
	// purchaseTicket(&customer, &ticket)

	// fmt.Println("After discount: ")

	// fmt.Println("Movie: ", ticket.Movie)
	// fmt.Println("Price: ", ticket.Price)
	// fmt.Println("Seat: ", ticket.Seat)

	// fmt.Println("Customer: ", customer)

	fmt.Println(ticket.GetPrice())
	fmt.Println(snack.GetPrice())
}
