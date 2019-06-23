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

	//读取30天数据
	go func() {
		for i := 1; i < 31; i++ {
			st := 1511280000
			et := 1511280000 + i*86400
			err := countData(db, st, et, dataChan)
			fmt.Println(err)
		}
		close(dataChan)
	}()

	//处理30天数据
	for row := range dataChan{
		fmt.Println(row)
		//1、将数据转化为SQL
		//2、开始事务
		//3、插入备份库
		//4、删除原表数据
	}
}

func countData(db *sql.DB, st int, et int, dataChan chan rowType) (err error) {
	var count int
	err = db.QueryRow(
		"select count(*) from t_o2o_dm_pay_tmp  pay_time between ? and ?", st, et,
	).Scan(&count)
	if err != nil {
		err = fmt.Errorf("读取数据总数出错：%s", err)
		return
	}
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Printf("countData 数据库关闭失败, st: %d, et: %d", st, et)
		}
	}()
	//控制并发
	limit := make(chan struct{}, runtime.NumCPU())
	for i := 0; i < count; i += 10000 {
		limit <- struct{}{}
		go func(i int) {
			defer func() {
				<-limit
			}()
			err = getData(db, st, et, i, dataChan)
			fmt.Println(err)
		}(i)
	}
	for i := 0; i < cap(limit); i++ {
		limit <- struct{}{}
	}
	close(limit)
	return
}

func getData(db *sql.DB, st int, et int, offset int, dataChan chan rowType) (err error) {
	rows, err := db.Query(
		"select * from t_o2o_dm_pay_tmp where pay_time between ? and ? limit ?, ? ", st, et, offset, 10000,
	)
	if err != nil {
		err = fmt.Errorf("读取数据出错, offser: %d, error:%s", offset, err)
		return
	}
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Printf("getData 数据库关闭失败, st: %d, et: %d, offset: %d", st, et, offset)
		}
	}()
	for rows.Next() {
		var row rowType
		err := rows.Scan(&row.id, &row.client_id, &row.device_id, &row.dm_id, &row.pay_type, &row.pay_total, &row.no, &row.pay_time, &row.pay_data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		dataChan <- row
	}
	return
}
