package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type ImportData struct {
	id    int16
	name  string
	value float64
	text1 string
	text2 string
}

func main() {
	stream, err := ioutil.ReadFile("test.csv")
	if err != nil {
		panic("file open error")
	}
	reader := csv.NewReader(strings.NewReader(string(stream)))

	var line ImportData
	for {
		err := Unmarshal(reader, &line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(line)
	}
}

func Unmarshal(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			f.SetString(record[i])
		case "int":
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			return &UnsupportedType{f.Type().String()}
		}
	}
	return nil
}

type FieldMismatch struct {
	expected, found int
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

type UnsupportedType struct {
	Type string
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type
}

func importTest() {
	defer errorHanlder()
	stream, err := ioutil.ReadFile("test.csv")
	// stream, err := ioutil.ReadFile("test_empty.csv")
	if err != nil {
		panic("file open error")
	}

	reader := csv.NewReader(strings.NewReader(string(stream)))
	reader.Comma = ';'

	data, err := reader.ReadAll()
	if err != nil {
		panic("file read error")
	}

	if data == nil {
		panic("no data in file")
	}

	line := ImportData{}
	// line.id, err := strconv.Atoi(data[0][0])
	// i, err := strconv.Atoi(data[0][0])
	if i, err := strconv.Atoi(data[1][0]); err != nil {
		panic("conversation error")
	} else {
		line.id = int16(i)
	}

	f, err := strconv.ParseFloat(data[1][2], 64)
	if err != nil {
		panic("conversation error")
	}
	line.value = float64(f)

	line.name, line.text1, line.text2 = data[1][1], data[1][3], data[1][4]
	fmt.Println(data)
	fmt.Println(line)
}

func errorHanlder() {
	if recov := recover(); recov != nil {
		fmt.Println(recov)
	}
}
