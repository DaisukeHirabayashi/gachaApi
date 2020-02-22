package main

import (
	"net/http"
  "github.com/labstack/echo"

)

func main() {
	e := echo.New()
  e.GET("/user/get", userGet)
  e.Logger.Fatal(e.Start(":1323"))
}
func userGet(c echo.Context) error {
    //emailとパスワードから今後はuseridを取ってくる予定。
    email := c.QueryParam("email")
    password := c.QueryParam("password")
    return c.String(http.StatusOK, "email:"+email+", password:"+password)
}
