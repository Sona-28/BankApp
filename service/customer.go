package service

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Cust struct{
	ctx context.Context
	mongoCollection *mongo.Collection
}

func InitCustomer(collection *mongo.Collection, ctx context.Context) interfaces.Icustomer{
	return &Cust{ctx,collection}
}
func(c *Cust) CreateCustomer(user *models.Customer)(*mongo.InsertOneResult,error){
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"account_id": 1}, // 1 for ascending, -1 for descending
		Options: options.Index().SetUnique(true),
	}
	_, err := c.mongoCollection.Indexes().CreateOne(c.ctx, indexModel)
	if err != nil {
		log.Fatal(err)
	}
	user.Customer_ID = primitive.NewObjectID()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password),8)
	user.Password = string(hashedPassword)
	res,err := c.mongoCollection.InsertOne(c.ctx, &user)
	if err!=nil{
		if mongo.IsDuplicateKeyError(err){
			log.Fatal("Duplicate key error")
		}
		return nil,err
	}
	// pwd:="78787878"
	// if err1:=bcrypt.CompareHashAndPassword([]byte(pwd),[]byte(user.Password));err1!=nil{
	// 	fmt.Println(err1)
	// }else{
	// 	fmt.Println("yes")
	// }
	return res,nil
}

func(c *Cust) GetCustomer() ([]*models.Customer, error) {
	filter := bson.D{}
	result, err :=  c.mongoCollection.Find(c.ctx,filter)
	if err!=nil{
		// fmt.Println(err.Error())
		return nil,err
	}else{
		var customers []*models.Customer
		for result.Next(c.ctx){
			post := &models.Customer{}
			err := result.Decode(post)
			if err!=nil{
				return nil,err
			}
			customers = append(customers, post)
		}
		if err:=result.Err(); err!=nil{
			return nil, err
		}
		if len(customers) == 0{
			return []*models.Customer{},nil
		}
		return customers,nil
	}
}