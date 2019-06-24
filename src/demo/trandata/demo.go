package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type rowType struct {
	id        int             `db:"id"`
	client_id sql.NullString  `db:"client_id"`
	device_id sql.NullInt64   `db:"device_id"`
	dm_id     sql.NullString  `db:"dm_id"`
	pay_type  sql.NullString  `db:"pay_type"`
	pay_total sql.NullString  `db:"pay_total"`
	no        sql.NullString  `db:"no"`
	pay_time  sql.NullString  `db:"pay_time"`
	pay_data  sql.NullString  `db:"pay_data"`
}


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
	dbMap :=  make(map[string]*sql.DB)

	dbMap["o2o"], err = connectMySQL("", "", "", "3306", "o2o")
	check(err)

	dbMap["o2o_dms_00"], err = connectMySQL("", "", "", "3306", "o2o_dms_00")
	check(err)

	dbMap["o2o_dms_01"], err = connectMySQL("", "", "", "3306", "o2o_dms_01")
	check(err)

	dbMap["o2o_dms_02"], err = connectMySQL("", "", "", "3306", "o2o_dms_02")
	check(err)

	dbMap["o2o_dms_03"], err = connectMySQL("", "", "", "3306", "o2o_dms_03")
	check(err)

	dbMap["o2o_dms_04"], err = connectMySQL("", "", "", "3306", "o2o_dms_04")
	check(err)

	dbMap["o2o_dms_05"], err = connectMySQL("", "", "", "3306", "o2o_dms_05")
	check(err)

	dbMap["o2o_dms_06"], err = connectMySQL("", "", "", "3306", "o2o_dms_06")
	check(err)

	dbMap["o2o_dms_07"], err = connectMySQL("", "", "", "3306", "o2o_dms_07")
	check(err)

	dbMap["o2o_dms_08"], err = connectMySQL("", "", "", "3306", "o2o_dms_08")
	check(err)

	dbMap["o2o_dms_09"], err = connectMySQL("", "", "", "3306", "o2o_dms_09")
	check(err)

	defer func() {
		for dmsDB := range dbMap {
			err := dbMap[dmsDB].Close()
			if err != nil {
				fmt.Printf("关闭%s失败, error : %s", dmsDB, err)
			}
		}
	}()


	dataChan := make(chan rowType, runtime.NumCPU())

	//读取30天数据
	go func() {
		for i := 1; i < 2; i++ {
			et := 1530374400 + i*86400
			st := et - 86400
			err := countData(dbMap["o2o"], st, et, dataChan)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("spend %.2fs \n", time.Since(start).Seconds())
		}
		close(dataChan)
		fmt.Println("读取数据完成")
	}()

	//处理30天数据
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range dataChan {
				//1、将数据转化为SQL
				//2、插入备份库
				//3、删除原表数据
				dbName, tableName := getTable(row.device_id)
				insertSql, deleteSql := getSql(row, tableName)
				err := tranData(dbMap[dbName], dbMap["o2o"], insertSql, deleteSql)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}
	wg.Wait()

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func tranData(dbDms *sql.DB, dbO2o *sql.DB, insertSql string, deleteSql string) (err error) {
	_, err = dbDms.Exec(insertSql)
	if err != nil {
		err = fmt.Errorf("数据插入失败, sql: %s, error: %s\r\n", insertSql, err)
		return
	}
	_, err = dbO2o.Exec(deleteSql)
	if err != nil {
		err = fmt.Errorf("数据删除失败, sql: %s, error: %s\r\n", deleteSql, err)
	}
	return
}

func countData(db *sql.DB, st int, et int, dataChan chan rowType) (err error) {
	var count int
	err = db.QueryRow(
		"select count(*) from t_o2o_dm_pay_tmp where pay_time between ? and ?", st, et,
	).Scan(&count)
	if err != nil {
		err = fmt.Errorf("读取数据总数出错：%s", err)
		return
	}
	fmt.Println(count)
	//控制并发
	limit := make(chan struct{}, runtime.NumCPU())
	for i := 0; i < count; i += 10000 {
		limit <- struct{}{}
		go func(i int) {
			defer func() {
				<-limit
			}()
			err = getData(db, st, et, i, dataChan)
			if err != nil {
				fmt.Printf("getData error: %s\r\n", err)
			}
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
		err := rows.Close()
		if err != nil {
			fmt.Printf("getData 数据库关闭失败, st: %d, et: %d, offset: %d\r\n", st, et, offset)
		}
	}()
	for rows.Next() {
		var row rowType
		err := rows.Scan(&row.id, &row.client_id, &row.device_id, &row.dm_id, &row.pay_type, &row.pay_total, &row.no, &row.pay_time, &row.pay_data)
		if err != nil {
			fmt.Printf("scan error: %s\r\n", err)
			continue
		}
		dataChan <- row
	}
	return
}

func getTable(deviceId sql.NullInt64) (dbName, tableName string) {
	dbName = "o2o_dms_"
	tableName = "t_o2o_dm_pay_tmp_"
	if deviceId.Valid  {
		deviceStr := strconv.FormatInt(deviceId.Int64, 10)
		if deviceId.Int64 < 10 {
			dbName += "00"
			tableName += "0"+ deviceStr
		} else {
			dbName += "0"+ deviceStr[len(deviceStr) -3 : len(deviceStr) -2]
			tableName += deviceStr[len(deviceStr) -2:]
		}
	} else {
		dbName += "00"
		tableName += "00"
	}
	return
}

func getSql(row rowType, tableName string) (insertSql, deleteSql string) {
	insertSql = fmt.Sprintf("insert into %s (id, client_id, device_id, dm_id, pay_type, pay_total, no, pay_time, pay_data) " +
		"values (%d, %s, %d, '%s', '%s', '%s', '%s', '%s', '%s')",
		tableName, row.id, row.client_id.String, row.device_id.Int64, row.dm_id.String, row.pay_type.String, row.pay_total.String,
		row.no.String, row.pay_time.String, row.pay_data.String)
	deleteSql = fmt.Sprintf("delete from t_o2o_dm_pay_tmp where id=%d", row.id)
	return
}
