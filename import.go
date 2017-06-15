package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	stream, err := ioutil.ReadFile("test.csv")
	if err != nil {
		fmt.Println("file open error")
		return
	}

	reader := csv.NewReader(strings.NewReader(string(stream)))
	reader.Comma = ';'

	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println("file read error")
		return
	}

	fmt.Println(data[1][0])
}
