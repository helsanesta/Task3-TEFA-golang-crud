package main

import (
	"fmt"
	"strconv"
)

func generateData(n int) []string {
	var baseNumber string = "082"
	customers := make([]string, n)
	for i := 0; i < n; i++ {
		mobileNumber := baseNumber + strconv.Itoa(i+1)
		customers[i] = mobileNumber
	}
	return customers
}

func sendPromoDiscount(customers []string) {
	for i := 0; i < len(customers); i++ {
		fmt.Printf("Sending Promo to %s\n", customers[i])
	}
	fmt.Printf("It's finished sending Promo to %d customers\n", len(customers))

	for i := 0; i < len(customers); i++ {
		fmt.Printf("Sending Discount to %s\n", customers[i])
	}
	fmt.Printf("It's finished sending Discount to %d customers\n", len(customers))
}

func main() {
	data := generateData(1000)
	sendPromoDiscount(data)
}
