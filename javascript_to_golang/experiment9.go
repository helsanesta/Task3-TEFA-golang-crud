package main

import "fmt"

var results [][]string

func arrange(array []string, memory []string) [][]string {
	if memory == nil {
		memory = make([]string, 0)
	}
	for i := 0; i < len(array); i++ {
		current := []string{array[i]}
		if len(array) == 1 {
			results = append(results, append(memory, current...))
		} else {
			subarray := make([]string, len(array)-1)
			copy(subarray, array[:i])
			copy(subarray[i:], array[i+1:])
			newMemory := append(memory, current...)
			arrange(subarray, newMemory)
		}
	}
	return results
}

func main() {
	candidates := []string{"Baswedan", "Subianto", "Maharani", "Ganjar"}
	chairs := arrange(candidates, nil)
	for _, chair := range chairs {
		fmt.Println(chair)
	}
}
