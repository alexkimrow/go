package main

import "fmt"

func main() {

	username := "guest"
	fmt.Println("Before :", username)

	var pointerUsername *string

	changeUsername(&username)

	fmt.Println("After :", username)

}

func changeUsername(username *string) {

	*username = "rolex"
}
