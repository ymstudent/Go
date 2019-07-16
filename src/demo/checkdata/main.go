package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"sync/atomic"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func connectMySQL(dbUser, dbPass, dbHost, dbPort, dbName string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		err = fmt.Errorf("\nconnection mysql %s error: %s", dbName, err)
		return
	}
	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("\nopen mysql %s error: %s", dbName, err)
		return
	}
	return
}

func main() {
	var err error
	dbMap := make(map[string]*sql.DB)
	tableChan := make(chan string)

	dbMap["o2o"], err = connectMySQL("", "", "", "3306", "ymblog")
	check(err)
	dbMap["o2oNew"], err = connectMySQL("", "", "", "3306", "ymblog")
	check(err)

	defer func() {
		for db := range dbMap {
			err := dbMap[db].Close()
			if err != nil {
				fmt.Printf("关闭%s失败, error : %s", db, err)
			}
		}
	}()

	defer func() {
		for db := range dbMap {
			err := dbMap[db].Close()
			if err != nil {
				fmt.Printf("关闭%s失败, error : %s", db, err)
			}
		}
	}()

	go selectAllTable(dbMap["o2o"], tableChan)

	var tableCount int64
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for tableName := range tableChan {
				atomic.AddInt64(&tableCount, 1)
				err := checkId(dbMap, tableName)
				if err != nil {
					//fmt.Println(err)
					err := checkCount(dbMap, tableName)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("all table has checked, total %d", tableCount)
}

//查询出数据库中所有的表名
func selectAllTable(db *sql.DB, tableChan chan<- string) {
	rows, err := db.Query("show tables")
	if err != nil {
		err = fmt.Errorf("读取数据出错, error:%s", err)
		return
	}

	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			fmt.Printf("scan error: %s\r\n", err)
			continue
		}
		tableChan <- table
	}
	close(tableChan)
}

//确认新旧库中各表最大ID是否相同
func checkId(dbMap map[string]*sql.DB, tableName string) (err error) {
	var o2oMaxId, o2oNewMaxId int64
	err = dbMap["o2o"].QueryRow("select max(id) from " + tableName).Scan(&o2oMaxId)
	if err != nil {
		err = fmt.Errorf("读取o2o.%s最大ID出错：%s", tableName, err)
		return
	}

	err = dbMap["o2oNew"].QueryRow("select max(id) from " + tableName).Scan(&o2oNewMaxId)
	if err != nil {
		err = fmt.Errorf("读取o2o.%s最大ID出错：%s", tableName, err)
		return
	}

	if o2oMaxId == o2oNewMaxId {
		fmt.Printf("%s maxId is the same\n", tableName)
	} else {
		fmt.Printf("%s maxId are different, o2o is %d, o2oNew is %d\n", tableName, o2oMaxId, o2oNewMaxId)
	}
	return
}

//确认新旧库中各表count是否相同
func checkCount(dbMap map[string]*sql.DB, tableName string) (err error) {
	var o2oCount, o2oNewCount int64
	err = dbMap["o2o"].QueryRow("select count(*) from " + tableName).Scan(&o2oCount)
	if err != nil {
		err = fmt.Errorf("统计o2o.%s数据出错：%s", tableName, err)
		return
	}
	err = dbMap["o2oNew"].QueryRow("select count(*) from " + tableName).Scan(&o2oNewCount)
	if err != nil {
		err = fmt.Errorf("统计o2o.%s数据出错：%s", tableName, err)
		return
	}

	if o2oCount == o2oNewCount {
		fmt.Printf("%s countNum is the same\n", tableName)
	} else {
		fmt.Printf("%s countNum are different, o2o is %d, o2oNew is %d\n", tableName, o2oCount, o2oNewCount)
	}
	return
}