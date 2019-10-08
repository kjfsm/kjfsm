package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Record struct {
	Args Args `json:"args"`
}

type Args []interface{}

func main() {
	tRecord := &Record{
		Args: []interface{}{"string", "12", 13, time.Now()},
	}
	tByte, tError := json.Marshal(tRecord)
	if tError != nil {
		fmt.Println(tError)
	}
	fmt.Println(string(tByte))

	var tNewRecord Record
	for _, tArg := range tRecord.Args {
		tString, tOK := tArg.(string)
		if !tOK {
			tNewRecord.Args = append(tNewRecord.Args, tArg)
			continue
		}
		tInt, tError := strconv.Atoi(tString)
		if tError != nil {
			tNewRecord.Args = append(tNewRecord.Args, tArg)
			continue
		}
		tNewRecord.Args = append(tNewRecord.Args, tInt)
	}

	tNewByte, tError := json.Marshal(tNewRecord)
	if tError != nil {
		fmt.Println(tError)
	}
	fmt.Println(string(tNewByte))

	// tRequest := req.Post(
	// 	req.URL("example.com", "logs"),
	// 	req.BodyJSON([]*Record{tRecord}),
	// )
}
