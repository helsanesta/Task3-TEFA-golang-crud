package main

import (
	"fmt"
	"time"
)

func findCompany(array []string) {
	tx := time.Now()
	for i := 0; i < len(array); i++ {
		if array[i] == "telkom" {
			fmt.Println("Company Found!")
		}
	}
	ty := time.Now()
	fmt.Printf("The performance is %v ms", (ty.Sub(tx).Seconds()))
}

func main() {
	company := "telkom"
	var largeCompanyName []string
	for i := 0; i < 10; i++ {
		largeCompanyName = append(largeCompanyName, company)
	}
	findCompany(largeCompanyName)
}
