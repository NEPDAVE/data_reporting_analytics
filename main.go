package main

import (
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "dra_challenge"
)

func main() {
	//file path could be os arg for when new files become available

	//######################
	//ENTERING MERCHANT DATA TO DB
	//######################

	// //sanitizing data and converting it to native go types
	// merchantsJson := MarshalMerchants("merchants.csv")
	//
	// //creating slice object we can iterate through for writing to database
	// merchants := UnmarshalMerchants(merchantsJson)
	//
	// //passing merchant slice to database write function
	// WriteToMerchantsTable(merchants)

	// //######################
	// //ENTERING DEVICE DATA TO DB
	// //######################
	//
	// //sanitizing data and converting it to native go types
	// devicesJson := MarshalDevices("devices.csv")
	//
	// //creating slice object we can iterate through for writing to database
	// devices := UnmarshalDevices(devicesJson)
	//
	// //passing device slice to database write function
	// WriteToDevicesTable(devices)

	//######################
	//ENTERING CASE DATA TO DB
	//######################

	// //sanitizing data and converting it to native go types
	casesJson := MarshalCases("cases.csv")

	//creating slice object we can iterate through for writing to database
	cases := UnmarshalCases(casesJson)

	//passing device slice to database write function
	WriteToCasesTable(cases)

	//fmt.Println(merchants)
	fmt.Println("THE END")
}
