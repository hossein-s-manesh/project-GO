package main

import (
	"fmt"
	"log"
	"myapp/config"

	"github.com/labstack/echo/v4"
)



func main()  {

	//get config
   
	err:=config.GetConfig()
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println(config.AppConfig.Server)
	

	//init server

	e:=echo.New()



	
	//routing


	e.GET("/icon",func(c echo.Context) error {

		fmt.Println("welcom to icon")

		return nil
	})



	//middleware
	
	
	
	
	
	//start server

	e.Start(":"+config.AppConfig.Server.Port)


}