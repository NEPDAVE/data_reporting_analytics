package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	//"time"
)

type Merchant struct {
	ID             string `json:"id"`
	Plan           string `json:"merchant_plan_name"`
	IsQSR          string `json:"is_qsr"`
	IsDemo         string `json:"is_demo"`
	MCCCode        string `json:"mcc_code"`
	MerchantType   string `json"merchant_type"`
	CloverCategory string `json:"clover_category"`
	//CreatedDate    time.Time `json:"created_date"`
	CreatedDate string `json:"created_date"`
}

//opens and loads csv into memory to prepare it for work
func MarshalMerchants(merchantCSV string) []byte {
	//csvFile, _ := os.Open("merchants.csv")
	csvFile, _ := os.Open(merchantCSV)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var merchants []Merchant
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//commenting this out because we ran into an issue unmarshaling the timestamp
		//convert string to timestamp
		//layout := "2006-01-02T15:04:05.000Z"
		// layout := "2006-01-02 15:04:05"
		// str := line[7]
		// time, err := time.Parse(layout, str)
		//
		// if err != nil {
		// 	fmt.Println(err)
		// }
		//fmt.Println(t)

		merchants = append(merchants, Merchant{
			ID:             line[0],
			Plan:           line[1],
			IsQSR:          line[2],
			IsDemo:         line[3],
			MCCCode:        line[4],
			MerchantType:   line[5],
			CloverCategory: line[6],
			//CreatedDate:    time,
			CreatedDate: line[7],
		})
	}

	merchantsJson, _ := json.Marshal(merchants)
	//fmt.Println(string(merchantsJson))
	return merchantsJson
}

func main() {
	//file path could be os arg for when new files become available
	merchantsJson := MarshalMerchants("merchants.csv")
	fmt.Println(reflect.TypeOf(merchantsJson))
	// for _, merchant := range merchantsJson {
	// 	fmt.Println("#########")
	// 	fmt.Println(reflect.TypeOf(merchant))
	// 	//fmt.Println(merchant.Plan)
	// 	if err := json.Unmarshal(merchant, &Merchant); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(Merchant)
	// 	fmt.Println("#########")
	// }
	// for _, merchantByte := range merchantsJson{
	//   if err := json.Unmarshal(merchantByte, &Merchant{}); err != nil {
	// 			panic(err)
	// 		}
	// }
	var merchants []Merchant
	if err := json.Unmarshal(merchantsJson, &merchants); err != nil {
		panic(err)
	}
  for _, merchant := range merchants {
    fmt.Println(merchant)
  }
	//fmt.Println(string(merchantsJson))
	fmt.Println("THE END")
}
