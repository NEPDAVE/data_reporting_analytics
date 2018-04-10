package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Case struct {
	CaseNumber    string    `json:"case_number"`
	Type          string    `json:"type"`
	SubType       string    `json:"request_reason_subtype"`
	SubTypeDetail string    `json:"reason_subtype_detail"`
	ID            string    `json:"id"`
	SerialNumber  string    `json"serial_numbers"`
	CreatedDate   time.Time `json:"created_date"`
}

//opens csv and sanitizes/marshals csv data into []byte
func MarshalCases(caseCSV string) []byte {
	csvFile, _ := os.Open(caseCSV)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var cases []Case
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//convert string to timestamp
		layout := "2006-01-02 15:04:05"
		str := line[6]
		time, err := time.Parse(layout, str)

		if err != nil {
			fmt.Println(err)
		}

		cases = append(cases, Case{
			CaseNumber:    line[0],
			Type:          line[1],
			SubType:       line[2],
			SubTypeDetail: line[3],
			ID:            line[4],
			SerialNumber:  line[5],
			CreatedDate:   time,
		})
	}

	casesJson, _ := json.Marshal(cases)
	fmt.Println(string(casesJson))
	return casesJson
}

//takes sanitized and marshaled merchantsJson data and unmarshals it into []Merchant
func UnmarshalCases(casesJson []byte) []Case {
	//merchants is the object that the []byte gets marshaled into
	var cases []Case
	if err := json.Unmarshal(casesJson, &cases); err != nil {
		panic(err)
	}
	return cases
}
