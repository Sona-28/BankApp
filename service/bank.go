package service

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Bank1 struct {
	ctx  context.Context
	coll *mongo.Collection
}


func InitBank(collection *mongo.Collection, ctx context.Context) interfaces.IBank {
	return &Bank1{ctx, collection}
}

func (b *Bank1) CreateBank(bank *models.Bank) (*mongo.InsertOneResult, error) {
	indexModel := []mongo.IndexModel{
		{
			Keys:    bson.M{"bank_id": 1}, // 1 for ascending, -1 for descending
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := b.coll.Indexes().CreateMany(b.ctx, indexModel)
	if err != nil {
		return nil, err
	}
	result, err := b.coll.InsertOne(b.ctx, bank)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *Bank1) GetBankById(id int64) (*models.Bank, error) {
	filter := bson.M{"bank_id": id}
	var bank models.Bank
	err := b.coll.FindOne(b.ctx, filter).Decode(&bank)
	if err != nil {
		return nil, err
	}
	return &bank, nil
}

func (b *Bank1) UpdateBankById(id int64, banks *models.Bank) (*mongo.UpdateResult, error) {
	iv := bson.M{"bank_id": id}
	fv := bson.M{"$set": &banks}
	result, err := b.coll.UpdateOne(b.ctx, iv, fv)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *Bank1) DeleteBankById(id int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"bank_id": id}
	result, err := b.coll.DeleteOne(b.ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
