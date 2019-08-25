package controller

import (
	"../service"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
)


func CheckUserInformation(ctx iris.Context){
   tablename :="go_user"
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
   QueryResult,err :=service.MysqlCheckID(tablename,password,username)
   checkErr(err)
	if QueryResult == sql.ErrNoRows {
		ctx.ViewData("name","loginError")
	} else {
		ctx.ViewData("name","Success")
		ctx.View("base.html")
	}

}




func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
