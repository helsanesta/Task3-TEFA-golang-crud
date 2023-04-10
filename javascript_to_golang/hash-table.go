package main

import (
	"fmt"
)

type HashTable struct {
	data [][]string
}

func NewHashTable(size int) *HashTable {
	return &HashTable{make([][]string, size)}
}

func (h *HashTable) _hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])*i) % len(h.data)
	}
	return hash
}

func (h *HashTable) Set(key, value string) {
	address := h._hash(key)
	if h.data[address] == nil {
		h.data[address] = []string{}
	}
	h.data[address] = append(h.data[address], key, value)
}

func (h *HashTable) Get(key string) string {
	address := h._hash(key)
	currentBucket := h.data[address]
	if currentBucket != nil {
		for i := 0; i < len(currentBucket); i += 2 {
			if currentBucket[i] == key {
				return currentBucket[i+1]
			}
		}
	}
	return ""
}

func main() {
	myHashTable := NewHashTable(100)
	myHashTable.Set("082124606606", "Rony Setyawan")
	fmt.Println(myHashTable.Get("082124606606"))
}
