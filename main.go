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

//enters merchant data into db
fun EnterMerchantData(merchantsCSV string ) {
  //sanitizing data and converting it to native go types
	merchantsJson := MarshalMerchants(merchantsCSV)

	//creating slice object we can iterate through for writing to database
	merchants := UnmarshalMerchants(merchantsJson)

	//passing merchant slice to database write function
	WriteToMerchantsTable(merchants)
}

//enters device data into db
func EnterDeviceData(devicesCSV string) {
  //sanitizing data and converting it to native go types
	devicesJson := MarshalDevices(devicesCSV)

	//creating slice object we can iterate through for writing to database
	devices := UnmarshalDevices(devicesJson)

	//passing device slice to database write function
	WriteToDevicesTable(devices)
}

//enters case data into db
func EnterCaseData(casesCSV string) {
	// //sanitizing data and converting it to native go types
	casesJson := MarshalCases(casesCSV)

	//creating slice object we can iterate through for writing to database
	cases := UnmarshalCases(casesJson)

	//passing device slice to database write function
	WriteToCasesTable(cases)
}

func main() {
	//file path for Enter*Data could be os arg for when new files become available

  EnterMerchantData("merchant.csv")
	EnterDeviceData("devices.csv")
	EnterCaseDatq("cases.csv")

	fmt.Println("THE END")
}
