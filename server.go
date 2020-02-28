package main

import (
  "github.com/labstack/echo"
	"gachaApi/handler"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
	e := echo.New()
  e.GET("/user/get", handler.UserGet)//ログイン認証情報からユーザー情報を取ってくる
	e.POST("/user/create", handler.UserCreate)//ユーザー情報の生成
	e.PUT("/user/update", handler.UserUpdate)//ユーザー情報のアップデート
  e.Logger.Fatal(e.Start(":1323"))
}
