package main

import (
	"gachaApi/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.POST("/user/create", handler.UserCreate) //ユーザー情報の生成

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/user/get", handler.UserGet()) //ログイン認証情報からユーザー情報を取ってくる
	//e.PUT("/user/update", handler.UserUpdate) //ユーザー情報のアップデート

	r := e.Group("/user/update")
	r.Use(middleware.JWT([]byte("secret")))
	r.PUT("", handler.UserUpdate()) //ユーザー情報のアップデート

	e.Logger.Fatal(e.Start(":1323"))
}
