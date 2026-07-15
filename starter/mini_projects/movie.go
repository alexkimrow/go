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

func purchaseTicket(customer *Customer, ticket *Ticket) {
	customer.WalletBalance = customer.WalletBalance - ticket.Price
	customer.Ticket = append(customer.Ticket, ticket.Movie, ticket.Seat)
}

func applyDiscount(ticket *Ticket) {
	ticket.Price = ticket.Price * 0.8
}

func main() {

	ticket := Ticket{
		Movie: "Dune",
		Price: 20.00,
		Seat:  "B12",
	}

	customer := Customer{
		Name:          "Rolex",
		WalletBalance: 50,
	}

	fmt.Println("Before discount: ")

	fmt.Println("Movie: ", ticket.Movie)
	fmt.Println("Price: ", ticket.Price)
	fmt.Println("Seat: ", ticket.Seat)

	applyDiscount(&ticket)
	purchaseTicket(&customer, &ticket)

	fmt.Println("After discount: ")

	fmt.Println("Movie: ", ticket.Movie)
	fmt.Println("Price: ", ticket.Price)
	fmt.Println("Seat: ", ticket.Seat)

	fmt.Println("Customer: ", customer)
}
