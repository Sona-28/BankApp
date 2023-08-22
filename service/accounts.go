package service

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Acc struct{
	ctx context.Context
	mongoCollection *mongo.Collection
}

func InitAccount(collection *mongo.Collection, ctx context.Context) interfaces.Iaccount{
	return &Acc{ctx,collection}
}

func (a *Acc)CreateAccount(account *models.Account)(*mongo.InsertOneResult,error){
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"account_id": 1}, // 1 for ascending, -1 for descending
		Options: options.Index().SetUnique(true),
	}
	_, err := a.mongoCollection.Indexes().CreateOne(a.ctx, indexModel)
	if err != nil {
		return nil,err
	}
	result,err := a.mongoCollection.InsertOne(a.ctx,account)
	if(err!=nil){
		return nil,err
	}
	return result,nil
}