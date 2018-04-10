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

type Device struct {
	ID           string    `json:"id"`
	DeviceEvent  string    `json:"device_event"`
	SerialNumber string    `json:"serial_number"`
	DeviceTypeID string    `json:"device_type_id"`
	MerchantID   string    `json:"merchant_id"`
	Timestamp    time.Time `json"timestamp"`
}

//opens csv and sanitizes/marshals csv data into []byte
func MarshalDevices(devicesCSV string) []byte {
	csvFile, _ := os.Open(devicesCSV)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var devices []Device
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//convert string to timestamp
		layout := "2006-01-02 15:04:05"
		str := line[5]
		time, err := time.Parse(layout, str)

		if err != nil {
			fmt.Println(err)
		}

		devices = append(devices, Device{
			ID:           line[0],
			DeviceEvent:  line[1],
			SerialNumber: line[2],
			DeviceTypeID: line[3],
			MerchantID:   line[4],
			Timestamp:    time,
		})
	}

	devicesJson, _ := json.Marshal(devices)
	//fmt.Println(string(devicesJson))
	return devicesJson
}

//takes sanitized and marshaled merchantsJson data and unmarshals it into []Merchant
func UnmarshalDevices(devicesJson []byte) []Device {
	//merchants is the object that the []byte gets marshaled into
	var devices []Device
	if err := json.Unmarshal(devicesJson, &devices); err != nil {
		panic(err)
	}
	return devices
}
