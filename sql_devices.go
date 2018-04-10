package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func WriteToDevicesTable(devices []Device) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	for _, device := range devices {

		sqlStatement := `
  INSERT INTO devices (id, device_event, serial_number, device_type_id,
		merchant_id, timestamp)
  VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING id`

		uuid := 0
		err = db.QueryRow(sqlStatement, device.ID, device.DeviceEvent, device.SerialNumber,
			device.DeviceTypeID, device.MerchantID, device.Timestamp).Scan(&uuid)
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

// CREATE TABLE devices (
// uuid serial PRIMARY KEY,
// id text,
// device_event text,
// serial_number text,
// device_type_id text,
// merchant_id text,
// timestamp timestamp
// );
