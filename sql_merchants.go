package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func WriteToMerchantsTable(merchants []Merchant) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	for _, merchant := range merchants {

		sqlStatement := `
  INSERT INTO merchants (id, merchant_plan_name, is_qsr, is_demo, mcc_code,
		merchant_type, clover_category, created_date)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
  RETURNING id`

		uuid := 0
		err = db.QueryRow(sqlStatement, merchant.ID, merchant.Plan, merchant.IsQSR,
			merchant.IsDemo, merchant.MCCCode, merchant.MerchantType,
			merchant.CloverCategory, merchant.CreatedDate).Scan(&uuid)
		if err != nil {
			//panic(err)
			log.Println(err)
		}
		fmt.Println("New record UUID is:", uuid)

	}
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println("Successfully connected!")
}

//Extra copy of table schema
// CREATE TABLE merchants (
// uuid serial PRIMARY KEY,
// id text,
// merchant_plan_name text,
// is_qsr text,
// is_demo text,
// mcc_code text,
// merchant_type text,
// clover_category text,
// created_date timestamp
// );
