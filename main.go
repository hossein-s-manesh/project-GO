package main

import (
	"context"
	"fmt"
	"log"
	"myapp/config"
	"myapp/database"
	"myapp/routing"
	

	"go.mongodb.org/mongo-driver/bson"

	"github.com/labstack/echo/v4"
)



func main()  {

	// get config
	
	db,err:=database.Connect()
	if err !=nil{
		log.Fatalln(err)
	}else{

		userCollections:=db.GetUserList()
		var ros bson.M
		userCollections.FindOne(context.TODO(),bson.D{{"name","hossein"}}).Decode(&ros)
		fmt.Println(ros)
		
	}


	server:=echo.New()

	fmt.Println(config.AppConfig.Server.Port)

	
	//routing


	routing.SetRout(server)
	//middleware



	
	//start server

	server.Start(":"+config.AppConfig.Server.Port)


}