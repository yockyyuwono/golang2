package dal

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	hlp "github.com/yockyyuwono/golang2/helper"
)

var db *sql.DB
var ctx context.Context
var err error

/*
var server = "192.168.1.227"
var database = "BelajarDb"
var userdb = "yocky"
var passworddb = "y0cky2012"
var encrypt = "disable"
//var port = 1434
*/

func GetFromDatabaseSingleDal(sqlQuery string) (string, error) {
	mycon, _ := hlp.LoadConfiguration()

	//connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;encrypt=%s", server, database, userdb, passworddb, encrypt)
	//connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;encrypt=%s", mycon.HostDatabase, mycon.DatabaseName, mycon.UserDatabase, mycon.PasswordDatabase, mycon.EncryptDatabase)

	connString := hlp.GetConnectionStringConfig()
	db, err = sql.Open(mycon.DbDriver, connString)

	//sqlQuery := fmt.Sprintf("select UserCode, Passwords from MsUser where UserCode = '%s'", usercode)
	fmt.Println(string(sqlQuery))
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	db.Close()
	if err != nil {
		db.Close()
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}

func GetFromDatabaseListDal(sqlQuery string) (string, error) {
	//mycon, _ := hlp.LoadConfiguration()
	//connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;encrypt=%s", mycon.HostDatabase, mycon.DatabaseName, mycon.UserDatabase, mycon.PasswordDatabase, mycon.EncryptDatabase)
	//db, err = sql.Open("sqlserver", connString)

	connString := hlp.GetConnectionStringConfig()
	db, err = sql.Open("sqlserver", connString)

	//sqlQuery := "select RowId,UserCode,Passwords,BadgeNumber,Ssn,LastLogin,LoginLocation,DepartmentId,Email,RoleId,CreatedBy,CreatedDate,LastUpdatedBy,LastUpdatedDate,Active from MsUser"
	//sqlQuery := "select RowId,UserCode,Passwords,BadgeNumber,Ssn,LastLogin,LoginLocation,DepartmentId,Email,"
	//sqlQuery += "RoleId,CreatedBy,CreatedDate,LastUpdatedBy,LastUpdatedDate,Active from MsUser"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(jsonData))
	return string(jsonData), nil
}

func SaveToDatabase(sqlQuery string) (string, error) {
	mycon, _ := hlp.LoadConfiguration()

	//connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;encrypt=%s", server, database, userdb, passworddb, encrypt)
	//connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;encrypt=%s", mycon.HostDatabase, mycon.DatabaseName, mycon.UserDatabase, mycon.PasswordDatabase, mycon.EncryptDatabase)

	connString := hlp.GetConnectionStringConfig()
	db, err = sql.Open(mycon.DbDriver, connString)

	//sqlQuery := fmt.Sprintf("select UserCode, Passwords from MsUser where UserCode = '%s'", usercode)
	fmt.Println(string(sqlQuery))
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	db.Close()
	if err != nil {
		db.Close()
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}

/*
func GetMsUserSingleJson(usercode string) (string, error) {
	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;encrypt=%s", server, database, userdb, passworddb, encrypt)
	db, err = sql.Open("sqlserver", connString)

	sqlQuery := fmt.Sprintf("select UserCode, Passwords from MsUser where UserCode = '%s'", usercode)
	fmt.Println(string(sqlQuery))
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}
*/
