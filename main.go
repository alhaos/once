package main

import (
	"fmt"
	"once/internal/digger"
	"os"
	"runtime/pprof"
	"time"
)

func main() {

	conf := digger.Config{SourceDirectory: `C:\tmp\007\MedgenInput`}

	d, err := digger.New(conf, []string{})
	if err != nil {
		panic(err)
	}

	f, err := os.Create("outCpuProf.pprof")
	if err != nil {
		panic(err)
	}
	t := time.Now()
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}

	_, err = d.FindNewOrderNumbers()
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(t).String())
	defer pprof.StopCPUProfile()

}
