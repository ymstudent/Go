package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func connectMySQL(dbUser, dbPass, dbHost, dbPort, dbName string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		err = fmt.Errorf("\nconnection mysql error: %s", err)
		return
	}
	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("\nopen mysql error: %s", err)
		return
	}
	return
}

func main() {
	start := time.Now()
	var err error
	dbMap := make(map[string]*sql.DB)

	dbMap["o2o"], err = connectMySQL("", "", "", "3306", "o2o")
	check(err)

	defer func() {
		for dmsDB := range dbMap {
			err := dbMap[dmsDB].Close()
			if err != nil {
				fmt.Printf("关闭%s失败, error : %s", dmsDB, err)
			}
		}
	}()

	for i := 1; i < 32; i++ {
		et := 1538323200 + i*86400
		st := et - 86400
		err := countData(dbMap["o2o"], st, et)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("spend %.2fs \n", time.Since(start).Seconds())
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func countData(db *sql.DB, st int, et int) (err error) {
	var minId, maxId int
	err = db.QueryRow(
		"select max(id), min(id) from t_o2o_dm_pay_tmp where pay_time between ? and ?", st, et,
	).Scan(&maxId, &minId)
	if err != nil {
		err = fmt.Errorf("读取ID总数出错：%s", err)
		return
	}
	fmt.Println(minId, maxId)
	//控制并发
	limit := make(chan struct{}, runtime.NumCPU())
	for i := minId; i < maxId; i += 100 {
		limit <- struct{}{}
		go func(i int) {
			defer func() {
				<-limit
			}()
			err = deleteData(db, i)
			if err != nil {
				fmt.Printf("deleteData error: %s\r\n", err)
			}
		}(i)
	}
	for i := 0; i < cap(limit); i++ {
		limit <- struct{}{}
	}
	close(limit)
	return
}

func deleteData(db *sql.DB, start int) (err error) {
	_, err = db.Exec(
		"delete from t_o2o_dm_pay_tmp where id between ? and ?", start, start+99,
	)
	if err != nil {
		err = fmt.Errorf("删除数据出错, error:%s", err)
		return
	}
	return
}
