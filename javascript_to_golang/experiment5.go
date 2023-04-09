package main

import (
	"fmt"
	"time"
)

func findAddress(addresses []string) {
	tx := time.Now()
	fmt.Println("The Default Address is", addresses[0])
	ty := time.Now()
	fmt.Println("The performance is", ty.Sub(tx).Seconds(), "ms")
}

func main() {
	address := "DKI Jakarta"
	var addresses []string
	for i := 0; i < 1000; i++ {
		addresses = append(addresses, address)
	}
	findAddress(addresses)
}
