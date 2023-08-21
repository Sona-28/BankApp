package interfaces

import (
	"bankDemo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Icustomer interface {
	CreateCustomer (*models.Customer)(*mongo.InsertOneResult,error)
	GetCustomer() ([]*models.Customer, error) 
}