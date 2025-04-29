package main

import (
	"fmt"
	"once/internal/digger"
)

func main() {

	conf := digger.Config{SourceDirectory: `c:\tmp\007`}

	d, err := digger.New(conf, []string{})
	if err != nil {
		panic(err)
	}

	orderNumbers, err := d.FindNewOrderNumbers()
	if err != nil {
		return
	}

	for _, number := range orderNumbers {
		fmt.Println(number)
	}

}
