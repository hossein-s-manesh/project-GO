package main

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/routing"
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

	server:=echo.New()



	
	//routing


	routing.SetRout(server)
	//middleware
	
	
	
	
	
	//start server

	server.Start(":"+config.AppConfig.Server.Port)


}