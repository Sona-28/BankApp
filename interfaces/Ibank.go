package interfaces

import (
	"bankDemo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type IBank interface{
	CreateBank(bank *models.Bank)(*mongo.InsertOneResult,error)
	GetBankById(id int64) (*models.Bank, error)
	UpdateBankById(id int64, bank *models.Bank) (*mongo.UpdateResult, error)
	DeleteBankById(id int64) (*mongo.DeleteResult, error)
}