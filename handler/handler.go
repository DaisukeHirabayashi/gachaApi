package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

//UserGet 情報を取ってくる
func UserGet(c echo.Context) error {
	email := c.QueryParam("email")
	password := c.QueryParam("password")
	db, err := sqlConnect()
	if err != nil {
		return c.String(http.StatusInternalServerError,"サーバーエラー")
	}
	defer db.Close()

	result := []*Users{}
	error := db.Where(&Users{Email: email, Password: password}).Find(&result).Error
	if error != nil || len(result) == 0 {
		fmt.Println(error)
	}
	for _, user := range result {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.ID
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return c.String(http.StatusNotFound,"emailかパスワードが違います。")
}

//UserCreate 情報作成
func UserCreate(c echo.Context) error {
	u := new(Users)
	if err := c.Bind(u); err != nil {
		return err
	}
	name := u.Name
	email := u.Email
	password := u.Password
	db, err := sqlConnect()
	if err != nil {
		return c.String(http.StatusInternalServerError,"サーバーエラー")
	}
	defer db.Close()

	error := db.Create(&Users{
		Name:     name,
		Email:    email,
		Password: password,
	}).Error
	if error != nil {
		return c.String(http.StatusInternalServerError,"サーバーエラー")
	}
	return c.String(http.StatusOK, "データ追加成功")
}

//UserUpdate ユーザー情報のアップデート
func UserUpdate(c echo.Context) error {
	u := new(Users)
	if err := c.Bind(u); err != nil {
		return err
	}
	id := u.ID
	name := u.Name
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	error := db.Model(Users{}).Where("id = ?", id).Update(&Users{
		Name: name,
	}).Error
	if error != nil {
		return c.String(http.StatusInternalServerError,"サーバーエラー")
	}
	return c.String(http.StatusOK, "データアップデート成功")
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
	Email    string `json:"email"`
	Password string `json:"password"`
}
