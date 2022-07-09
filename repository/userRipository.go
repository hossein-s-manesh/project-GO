package repository

import (
	"context"
	"log"
	"myapp/database"
	"myapp/model/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "fmt"
)

type UserRepositor interface{
	GetUserList() ([]user.Jasoner,error )
	GetById(id string) (user.Jasoner,error )
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



func (userRepositor userRepositor) GetById(id string) (user.Jasoner,error ) {

	objectId,err:=primitive.ObjectIDFromHex(id)
	if err !=nil{
		return user.Jasoner{},err 
	}

	var work user.Jasoner

	userCollections:=userRepositor.db.GetUserList()
	userCollections.FindOne(context.TODO(),bson.D{
		{"_id",objectId},
	}).Decode(&work)
	return work,nil
}
