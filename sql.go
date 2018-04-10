package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "dra_challenge"
)

func WriteToMerchantsTable() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

// CREATE TABLE merchants (
// dra_challenge(#   id text PRIMARY KEY,
// dra_challenge(#   merchant_plan_name text,
// dra_challenge(#   is_qsr text,
// dra_challenge(#   is_demo text,
// dra_challenge(#   mcc_code text,
// dra_challenge(#   merchant_type text,
// dra_challenge(#   clover_category text,
// dra_challenge(#   created_date timestamp
// dra_challenge(# );
