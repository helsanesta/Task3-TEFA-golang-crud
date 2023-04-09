package main

import (
	"fmt"
	"strconv"
)

func generateData(n int) []string {
	baseNumber := "082"
	customers := make([]string, n)
	for i := 0; i < n; i++ {
		mobileNumber := baseNumber + strconv.Itoa(i+1)
		customers[i] = mobileNumber
	}
	return customers
}

func sendPromoDiscount(data1 []string, data2 []string) {
	for _, customer := range data1 {
		fmt.Printf("Sending Promo to %v\n", customer)
	}
	fmt.Printf("It's finished sending Promo to %v customers\n", len(data1))

	for _, customer := range data2 {
		fmt.Printf("Sending Discount to %v\n", customer)
	}
	fmt.Printf("It's finished sending Discount to %v customers\n", len(data2))
}

func main() {
	dataPromo := generateData(100000000)
	dataDiscount := generateData(1000)
	sendPromoDiscount(dataPromo, dataDiscount)
}
