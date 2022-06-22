package routing

import (
	"fmt"
	"net/http"
	"myapp/controller"

	"github.com/labstack/echo/v4"
)



func SetRout(e *echo.Echo)error {

	MayFunc:=func(c echo.Context) error {
		return c.String(http.StatusOK,"helllllo")
	}
	

	g:=e.Group("users")

	g.GET("/get/15",MayFunc)

	g.GET("/get/:id/good",func(c echo.Context) error {
		return c.String(http.StatusOK,"helllllo")
	})

	g.GET("/get/*",controller.Myfunc2)
	
	g.GET("/get/foods/id:",func(c echo.Context) error {
		return c.String(http.StatusOK,"helllllo")
	}).Name="hoo"
	
	g.GET("/get/:id/*/foods/:id/*",func(c echo.Context) error {
		return c.String(http.StatusOK,"helllllo")
	}).Name="hello"
	
		for i,rout:=range e.Routes(){
			fmt.Println(i,rout.Method,rout.Path,rout.Name)
		}
	
	fmt.Println(e.URI(controller.Myfunc2,19))
	fmt.Println(e.URI(MayFunc,15))
	fmt.Println(e.Reverse("hello",18,13,14))
	fmt.Println(e.Reverse("hoo",34))

	
	return nil
}
