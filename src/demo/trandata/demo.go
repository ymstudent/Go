package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
)

type rowType struct {
	id 			int 				`db:"id"`
	client_id 	sql.NullInt64		`db:"client_id"`
	device_id 	sql.NullInt64		`db:"device_id"`
	dm_id 		sql.NullInt64		`db:"dm_id"`
	pay_type 	sql.NullInt64		`db:"pay_type"`
	pay_total 	sql.NullFloat64		`db:"pay_total"`
	no 			sql.NullString		`db:"no"`
	pay_time 	sql.NullInt64		`db:"pay_time"`
	pay_data 	sql.NullString		`db:"pay_data"`
}

func check(e error)  {
	if e != nil {
		panic(e)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/o2o")
	check(err)

	dataChan := make(chan rowType, runtime.NumCPU())
	//每个goroutine一天
	for i := 1; i < 15; i++ {
		st := 1511280000
		et := 1511280000 + i*86400
		go getData(db, st, et, dataChan)
	}

	for j := 0; j < 20; j++ {

	}
}

func getData(db *sql.DB, st int, et int, dataChan chan rowType)  {
	rows, err := db.Query(
		"select * from t_o2o_dm_pay where pay_time between ? and ? limit 10000", st, et,
	)
	check(err)
	defer db.Close()
	for rows.Next() {
		var row rowType
		err := rows.Scan(&row.id, &row.client_id, &row.device_id, &row.dm_id, &row.pay_type, &row.pay_total, &row.no, &row.pay_time, &row.pay_data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		dataChan <- row
	}
}
