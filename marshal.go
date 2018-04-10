package main

import (
	"encoding/json"
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

func (m Merchant) MarshalMerchant(merchant Orders) []byte {

	merchantByte, err := json.Marshal(merchant)

	if err != nil {
		fmt.Println(err)
	}

	return merchantByte
}

func (m Merchant) UnmarshalMerchant(merchantByte []byte) *Merchant {

	err := json.Unmarshal(priceByte, &p)

	if err != nil {
		panic(err)
	}

	return &m
}
