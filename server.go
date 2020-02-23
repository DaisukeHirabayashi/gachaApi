package main

import (
	"net/http"
	"fmt"
  "github.com/labstack/echo"
	 _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
)

func main() {
	e := echo.New()
  e.GET("/user/get", userGet)//ログイン認証情報からユーザー情報を取ってくる
	e.POST("/user/create", userCreate)//ユーザー情報の生成
	
  e.Logger.Fatal(e.Start(":1323"))
}

func userGet(c echo.Context) error {
    email := c.QueryParam("email")
    password := c.QueryParam("password")
		db, err := sqlConnect()
		if err != nil {
				panic(err.Error())
		}
		defer db.Close()

		result := []*Users{}
		error := db.Where(&Users{Email: email, Password: password}).Find(&result).Error
		if error != nil || len(result) == 0 {
				fmt.Println(error)
		}
		for _, user := range result {
				return c.String(http.StatusOK, "name:"+user.Name)
		}
    return c.String(http.StatusOK, "emailかパスワードが違います。")
}

func userCreate(c echo.Context) error {
    name := c.FormValue("name")
		email := c.FormValue("email")
    password := c.FormValue("password")
		db, err := sqlConnect()
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    error := db.Create(&Users{
        Name:     name,
        Email:      email,
        Password:  password,
    }).Error
    if error != nil {
        return(err)
    }
		return c.String(http.StatusOK, "データ追加成功")
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
    DBMS := "mysql"
    USER := "root"
    PASS := "root"
    PROTOCOL := "tcp(localhost:3306)"
    DBNAME := "gacha"
    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}
// Users ユーザー情報のテーブル情報
type Users struct {
    ID       int
    Name     string `json:"name"`
    Email  string `json:"email"`
    Password string `json:"password"`
}
