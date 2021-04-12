package dal

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	hlp "github.com/yockyyuwono/golang2/helper"
)

func GetUserAuth_dal(sqlQuery string) (bool, error) {
	mycon, _ := hlp.LoadConfiguration()

	connString := hlp.GetConnectionStringConfig()
	db, err = sql.Open(mycon.DbDriver, connString)

	fmt.Println(string(sqlQuery))
	rows, err := db.Query(sqlQuery)
	//hlp.checkErr(err)
	if err != nil {
		db.Close()
		return false, err
	}
	iCnt := checkCount(rows)
	defer rows.Close()

	if iCnt < 1 {
		return false, nil
	}
	return true, err
	//fmt.Println(string(jsonData))

}
func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		//checkErr(err)
		if err != nil {
			panic(err)
		}
	}
	return count
}
