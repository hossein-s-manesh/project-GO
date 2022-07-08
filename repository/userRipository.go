package repository

import (
	"context"
	"log"
	"myapp/database"
	"myapp/model/user"

	"go.mongodb.org/mongo-driver/bson"
	// "fmt"
)

type UserRepositor interface{
	GetUserList() ([]user.Jasoner,error )
}

type userRepositor struct{
	db database.Db
}


func NewUserRipository()UserRepositor{
	db,err:=database.Connect()
	if err !=nil{
		log.Fatalln(err)
	}
	return userRepositor{
		db: db,
	}
}


func (userRepositor userRepositor) GetUserList() ([]user.Jasoner,error ) {
	userCollections:=userRepositor.db.GetUserList()
	ctx:=context.TODO()
	cursor,err:=userCollections.Find(ctx,bson.D{})
	if err !=nil{
		return nil,err
	}
	var user []user.Jasoner
	err=cursor.All(ctx,&user)
	if err !=nil{
		return nil,err
	}
	
	return user,nil
}
