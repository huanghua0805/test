package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"../config"

)


var MysqlConfig = config.MysqlAddr

//检查身份信息
func MysqlCheckID(tablename string,password string,username string)(queryResult error,err error){
	db, err := sql.Open("mysql",MysqlConfig )
	checkErr(err)

	result := db.QueryRow("SELECT * FROM"+tablename+" WHERE username="+username).Scan(&password)

	db.Close()
	return result,err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
