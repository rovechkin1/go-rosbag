package main

import (
	"fmt"
	"os"

	"github.com/lherman-cs/go-rosbag"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("example.bag")
	must(err)
	defer f.Close()

	msgCnt := 0
	decoder := rosbag.NewDecoder(f)
	for {
		record, err := decoder.Read()
		if err != nil && err.Error() == "EOF" {
			fmt.Printf("EOF\n")
			return
		} else {
			must(err)
		}

		switch record := record.(type) {

		case *rosbag.RecordMessageData:
			msgCnt+=1
			data := make(map[string]interface{})
			err = record.ViewAs(data)
			must(err)
			s,_ := record.String()
			fmt.Printf("%v, cnt: %v\n",s, msgCnt)

		default:
			s,_ := record.(rosbag.Record).String()
			fmt.Printf("%v\n",s)
		}

		record.Close()
	}
}
