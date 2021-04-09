package businesscore

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	dal "github.com/yockyyuwono/golang2/dal"
	mdl "github.com/yockyyuwono/golang2/model"
)

func GetMsUserSingle(usercode string) (string, error) {

	sqlQuery := fmt.Sprintf("select UserCode, Passwords from MsUser where UserCode = '%s'", usercode)
	jsonData, err := dal.GetMsUserSingleDal((sqlQuery))

	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
func GetMsUserList() (string, error) {

	//sqlQuery := fmt.Sprintf("select UserCode, Passwords from MsUser where UserCode = '%s'", usercode)
	sqlQuery := "select RowId,UserCode,Passwords,BadgeNumber,Ssn,LastLogin,LoginLocation,DepartmentId,Email,"
	sqlQuery += "RoleId,CreatedBy,CreatedDate,LastUpdatedBy,LastUpdatedDate,Active from MsUser"

	jsonData, err := dal.GetMsUserListDal((sqlQuery))

	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
func AddNewUserSingle(userSave mdl.UserSave) (string, error) {

	//sqlQuery := fmt.Sprintf("select UserCode, Passwords from MsUser where UserCode = '%s'", userSave.UserCode)
	sqlQuery := "INSERT INTO MsUSer ("
	sqlQuery += "UserCode"
	sqlQuery += ",Passwords"
	sqlQuery += ",BadgeNumber"
	sqlQuery += ",Ssn"
	sqlQuery += ",LastLogin"
	sqlQuery += ",LoginLocation"
	sqlQuery += ",DepartmentId"
	sqlQuery += ",Email"
	sqlQuery += ",RoleId"
	sqlQuery += ",CreatedBy"
	sqlQuery += ",CreatedDate"
	sqlQuery += ",LastUpdatedBy"
	sqlQuery += ",LastUpdatedDate"
	sqlQuery += ",Active"
	sqlQuery += ") "
	sqlQuery += "VALUES ( "
	sqlQuery += "'" + userSave.UserCode + "'"
	sqlQuery += ",'" + userSave.Passwords + "'"
	sqlQuery += ",'" + userSave.BadgeNumber + "'"
	sqlQuery += ",'" + userSave.Ssn + "'"
	sqlQuery += ",'" + userSave.LastLogin + "'"
	sqlQuery += ",'" + userSave.LoginLocation + "'"
	sqlQuery += ",'" + userSave.DepartmentId + "'"
	sqlQuery += ",'" + userSave.Email + "'"
	sqlQuery += ",'" + userSave.RoleId + "'"
	sqlQuery += ",'" + userSave.CreatedBy + "'"
	sqlQuery += ",'" + userSave.CreatedDate + "'"
	sqlQuery += ",'" + userSave.LastUpdatedBy + "'"
	sqlQuery += ",'" + userSave.LastUpdatedDate + "'"
	sqlQuery += ",'" + userSave.Active + "'"
	sqlQuery += "); "

	jsonData, err := dal.GetMsUserSingleDal((sqlQuery))

	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

/*


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
