package controller

import (

	// "fmt"

	// "myapp/model/user"
	"myapp/servis"

	"net/http"

	"github.com/labstack/echo/v4"
)



func Myfunc2(c echo.Context) error {

	userServis :=servis.NewUserServis()
	listjson,err:=userServis.GetServis()
	if err !=nil{
		println(err)
	}

	return c.JSON(http.StatusOK,listjson)
}


