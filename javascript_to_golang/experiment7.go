package main

import "fmt"

func logPairs(array1 []string, array2 []string, words string) {
	counter := 0
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			counter++
			fmt.Println(words, counter, array1[i], "dan", array2[j])
		}
	}
}

func main() {
	foods := []string{"Sate", "Bakso", "Dimsum", "Rames"}
	drinks := []string{"Jeruk", "Teh", "Kelapa", "Cendol"}
	logPairs(foods, drinks, "Menu")
}
