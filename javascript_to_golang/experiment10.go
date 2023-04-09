package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findCompany(array []string, input int) {
	for i := 0; i < len(array); i++ {
		if i == input {
			fmt.Printf("Company Found: %s\n", array[input])
			break
		}
		fmt.Printf("Searching Company... %d\n", i+1)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	telco := []string{"Telkom", "Indosat", "XL", "Verizon", "AT&T", "Nippon", "Vodafone", "Orange", "KDDI", "Telefonica", "T-Mobile"}
	company := rand.Intn(11)
	findCompany(telco, company)
}
