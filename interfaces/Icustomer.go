package interfaces

import (
	"bankDemo/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Icustomer interface {
	CreateCustomer(*models.Customer)(*mongo.InsertOneResult,error)
	GetCustomerById(int64) (*models.Customer, error)
	UpdateCustomerById(int64, *models.Customer) (*mongo.UpdateResult, error)
	DeleteCustomerById(int64) (*mongo.DeleteResult, error)
}