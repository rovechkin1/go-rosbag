package main

import (
	"os"
	"runtime/pprof"

	"github.com/k0kubun/pp"
	"github.com/lherman-cs/go-rosbag"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cpu, err := os.Create("cpu.out")
	must(err)
	defer cpu.Close()

	must(pprof.StartCPUProfile(cpu))
	defer pprof.StopCPUProfile()
	f, err := os.Open("example.bag")
	must(err)
	defer f.Close()

	decoder := rosbag.NewDecoder(f)
	var record rosbag.Record

	for {
		op, err := decoder.Read(&record)
		must(err)

		if op == rosbag.OpMessageData {
			v := make(map[string]interface{})
			must(record.UnmarshallTo(v))
			pp.Println(v)
		}
	}
}
