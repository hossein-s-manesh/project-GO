package controller

import (
	
	"fmt"
	
	"net/http"

	"github.com/labstack/echo/v4"
)



func Myfunc2(c echo.Context) error {


	idString:=c.Param("id")
	fmt.Println(idString)
	return c.String(http.StatusOK,"helllllo",)
}


type UserSearchInput struct{
	Name string `query:"name"`
	Family string `query:"family"`
	Age int `query:"age"`
}





func Myfunc3(c echo.Context) error {

	idString:=new(UserSearchInput)
	err:=c.Bind(idString)
	if err !=nil{
		return err
	}
	fmt.Println(idString)
	


	
	return c.String(http.StatusOK,"helllllo",)
}


