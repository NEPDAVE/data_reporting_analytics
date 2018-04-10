package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func WriteToCasesTable(cases []Case) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//case is a reserved word so im putting a _ infront
	for _, _case := range cases {

		sqlStatement := `
  INSERT INTO cases (case_number, type, request_reason_subtype,
    reason_subtype_detail, id, serial_numbers, created_date)
  VALUES ($1, $2, $3, $4, $5, $6, $7)
  RETURNING id`

		uuid := 0
		err = db.QueryRow(sqlStatement, _case.CaseNumber, _case.Type, _case.SubType,
			_case.SubTypeDetail, _case.ID, _case.SerialNumber, _case.CreatedDate).Scan(&uuid)
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
// CREATE TABLE cases (
// uuid serial PRIMARY KEY,
// case_number text,
// type text,
// request_reason_subtype text,
// reason_subtype_detail text,
// id text,
// serial_numbers text,
// created_date timestamp
// );
