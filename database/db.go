package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


type Db struct{
	clinet *mongo.Client
}



func Connect()( Db,error) {

	ctx:=context.TODO()

	clinet,err:=mongo.Connect(ctx,options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil{
		return Db{},err
	}


	err=clinet.Ping(ctx,readpref.Primary())
	if err !=nil{
		return Db{},err
	}

	
	
	return Db{clinet: clinet},nil
}



func(db Db) GetUserList()  *mongo.Collection {
	userCollections:=db.clinet.Database("hsm").Collection("users")

	return userCollections
}