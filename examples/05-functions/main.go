package main

import (
	"fmt"

	"github.com/mcnull/go-provide"
)

// Define a struct to hold data

type myData struct {
	txt string
}

// Define a factory function to create an instance of myData

func myDataFactory() *myData {
	return &myData{txt: "Hello, World!"}
}

// Define a function type, which we'll use to register a factory function

type myFunc func()

// Define a factory function to create an instance of myFunc

func myFuncFactory(d *myData) myFunc {
	// d is injected by the container
	return func() {
		fmt.Printf("%s\n", d.txt)
	}
}

func main() {
	// Register the factory functions
	provide.Set[*myData](myDataFactory)
	provide.Set[myFunc](myFuncFactory)

	// Get the instance of myFunc
	myFunc, err := provide.Get[myFunc]()

	if err != nil {
		panic(err)
	}

	myFunc() // Output: Hello, World!
}
