package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map

	// Store a key-value pair
	sm.Store("foo", "bar")

	// Load a value
	value, ok := sm.Load("foo")
	if ok {
		fmt.Println("Loaded value:", value)
	}

	// Store another key-value pair
	sm.Store("baz", 42)

	// Delete a key
	sm.Delete("foo")

	// LoadOrStore: loads if present, otherwise stores and returns the new value
	actual, loaded := sm.LoadOrStore("baz", 100)
	fmt.Println("Actual:", actual, "Loaded:", loaded)

	// Iterate over all key-value pairs
	sm.Range(func(key, value interface{}) bool {
		fmt.Println("Key:", key, "Value:", value)
		return true // return false to stop iteration
	})
}
