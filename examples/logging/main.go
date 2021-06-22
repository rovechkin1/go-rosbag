package main

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/lherman-cs/go-rosbag"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("/Users/ruslan/Downloads/Sample-Data/Sample-Data.bag")
	must(err)
	defer f.Close()

	msgCnt := 0
	decoder := rosbag.NewDecoder(f)
	for {
		record, err := decoder.Read()
		if err != nil && err.Error() == "EOF" {
			pp.Printf("EOF\n")
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
			pp.Printf("Message: %v %v\n",record.ConnectionHeader().Topic, msgCnt)

		default:
			s,_ := record.(rosbag.Record).String()
			pp.Printf("%v\n",s)
		}

		record.Close()
	}
}
