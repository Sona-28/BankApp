package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct{
	Account_ID int64 `json:"account_id" bson:"account_id"`
	Customer_ID primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	Account_type string `json:"account_type" bson:"account_type"`
	Branch string `json:"branch" bson:"branch"`
}