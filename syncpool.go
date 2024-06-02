package main

import (
	"fmt"
	"sync"
)

func main() {
	var numPool = sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	// Get an object from the pool
	num := numPool.Get().(int)
	fmt.Println("Got from pool:", num)

	// Put an object back into the pool
	numPool.Put(42)

	// Get the object back from the pool
	num = numPool.Get().(int)
	fmt.Println("Got from pool:", num)

	// Put the object back into the pool again
	numPool.Put(100)
  numPool.Put(101)

	// Get the object back from the pool
	num = numPool.Get().(int)
	fmt.Println("Got from pool:", num)
  num = numPool.Get().(int)
	fmt.Println("Got from pool:", num)
}
