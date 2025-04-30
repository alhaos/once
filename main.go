package main

import (
	"once/internal/digger"
	"os"
	"runtime/pprof"
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

	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}

	_, err = d.FindNewOrderNumbers()
	if err != nil {
		panic(err)
	}

	defer pprof.StopCPUProfile()

}
