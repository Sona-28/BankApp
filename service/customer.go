package service

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Cust struct{
	ctx context.Context
	mongoCollection *mongo.Collection
}

func InitCustomer(collection *mongo.Collection, ctx context.Context) interfaces.Icustomer{
	return &Cust{ctx,collection}
}
func(c *Cust) CreateCustomer(user *models.Customer)(*mongo.InsertOneResult,error){
	user.Customer_ID = primitive.NewObjectID()
	res,err := c.mongoCollection.InsertOne(c.ctx, &user)
	if err!=nil{
		return nil,err
	}
	return res,nil
}