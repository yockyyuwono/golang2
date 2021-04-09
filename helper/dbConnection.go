package helper

import "database/sql"

func GetConnectionStringConfig() string {
	mycon, _ := LoadConfiguration()
	consString := "server=" + mycon.HostDatabase + "; database=" + mycon.DatabaseName + ";user id=" + mycon.UserDatabase + ";password=" + mycon.PasswordDatabase + ";encrypt=" + mycon.EncryptDatabase
	return consString

}

//func dbConn() (db *sql.DB,error) {
func dbConn() (*sql.DB, error) {
	mycon, _ := LoadConfiguration()

	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	//connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;encrypt=%s", mycon.HostDatabase, mycon.DatabaseName, mycon.UserDatabase, mycon.PasswordDatabase, mycon.EncryptDatabase)
	db, err := sql.Open(mycon.DbDriver, "server="+mycon.HostDatabase+"; database="+mycon.DatabaseName+";user id="+mycon.UserDatabase+";password="+mycon.PasswordDatabase+";encrypt="+mycon.EncryptDatabase)
	if err != nil {
		//	panic(err.Error())
		//var errString string = err.Error()
		return db, err
	}

	return db, nil
}

func rtnstring() string {
	balikan := "asdasdsd"
	return balikan
}

/*
import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

var db *sql.DB
var ctx context.Context
var err error

type Configuration struct {
	server     string
	database   string
	userdb     string
	passworddb string
	encrypt    string
}

func GetDbConnection() (*sql.DB, error) {

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration)

	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;encrypt=%s", configuration.server, configuration.database, configuration.userdb, configuration.passworddb, configuration.encrypt)
	db, err = sql.Open("sqlserver", connString)

	if err != nil {
		return nil, err
	}
	return db, nil
}

*/
