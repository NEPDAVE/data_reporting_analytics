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

type Merchant struct {
	ID             string    `json:"id"`
	Plan           string    `json:"merchant_plan_name"`
	IsQSR          string    `json:"is_qsr"`
	IsDemo         string    `json:"is_demo"`
	MCCCode        string    `json:"mcc_code"`
	MerchantType   string    `json"merchant_type"`
	CloverCategory string    `json:"clover_category"`
	CreatedDate    time.Time `json:"created_date"`
	//CreatedDate    string `json:"created_date"`
}

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

		//convert string to timestamp
		//layout := "2006-01-02T15:04:05.000Z"
		layout := "2006-01-02 15:04:05"
		str := line[7]
		time, err := time.Parse(layout, str)

		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(t)

		merchants = append(merchants, Merchant{
			ID:             line[0],
			Plan:           line[1],
			IsQSR:          line[2],
			IsDemo:         line[3],
			MCCCode:        line[4],
			MerchantType:   line[5],
			CloverCategory: line[6],
			CreatedDate:    time,
			//CreatedDate:    line[7],
		})
	}

	merchantsJson, _ := json.Marshal(merchants)
	//fmt.Println(string(merchantsJson))
	return merchantsJson
}

// func MerchantCount(merchantsJson []byte) {
// 	for _, merchant := range merchantsJson {
// 		fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
//
// 	}
// }

func main() {
	//file path could be os arg for when new files become available
	merchantsJson := MarshalMerchants("merchants.csv")
	fmt.Println(string(merchantsJson))
}
