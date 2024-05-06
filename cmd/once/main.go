package main

import (
	"fmt"
	"once/internal/facory"
)

func main() {

	f := facory.NewFactory()
	doc1, err := f.NewDocument("One")
	if err != nil {
		panic(err)
	}

	doc2, err := f.NewDocument("Two")
	if err != nil {
		panic(err)
	}

	fmt.Println("Name of doc is:", doc1.GetName())
	fmt.Println("Name of doc is:", doc2.GetName())
}
