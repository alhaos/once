package main

import (
	"fmt"
	"once/internal/digger"
)

func main() {

	conf := digger.Config{SourceDirectory: `C:\tmp\007\MedgenInput`}

	d, err := digger.New(conf, []string{})
	if err != nil {
		panic(err)
	}

	orderNumbers, err := d.FindNewOrderNumbers()
	if err != nil {
		panic(err)
	}

	for _, number := range orderNumbers {
		fmt.Println(number)
	}

}
