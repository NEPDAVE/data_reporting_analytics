package main

import (
	"fmt"
)

func main() {
	//file path could be os arg for when new files become available

	//sanitizing data and converting it to native go types
	merchantsJson := MarshalMerchants("merchants.csv")

	//creating slice object we can iterate through for writing to database
	merchants := UnmarshalMerchants(merchantsJson)

	//passing merchant slice to database write function
	WriteToMerchantsTable(merchants)

	// for _, merchant := range merchants {
	// 	fmt.Println(merchant.ID)
	// 	fmt.Println(merchant.Plan)
	// }
	//fmt.Println(merchants)
	fmt.Println("THE END")
}
