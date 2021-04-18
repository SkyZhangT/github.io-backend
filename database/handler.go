package database

import (
	"context"

	"github.io-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type DBInterface interface {
	Insert(models.Item) (models.Item, error)
	Update(string, interface{}) (int64, error)
	Delete(string) (int64, error)
	Get(string) (models.Item, error)
	NextTen(int64) ([]models.Item, error)
	Search(interface{}) ([]models.Item, error)
}

type PostsClient struct{
	Ctx context.Context
	Col *mongo.Collection
}

func (c *PostsClient) Insert(data models.Item) (models.Item, error){
	item := models.Item{}

	res, err := c.Col.InsertOne(c.Ctx, data)
	if err != nil{
		return item, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}

func (c *PostsClient) Delete(id string) (int64, error){
	var count int64 = 0

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return count, err
	}

	res, err := c.Col.DeleteOne(c.Ctx, bson.M{"_id": _id})
	if err != nil{
		return count, err
	}
	return res.DeletedCount, nil
}

func (c *PostsClient) Get(id string) (models.Item, error){
	item := models.Item{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return item, err
	}

	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&item)
	if err != nil{
		return item, err
	}

	return item, nil
}

func (c *PostsClient) NextTen(offset int64) ([]models.Item, error){
	items := []models.Item{}

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"time": -1})
	findOptions.SetLimit(10)
	findOptions.SetSkip(offset)

	cursor, err := c.Col.Find(c.Ctx, findOptions)
	if err != nil{
		return items, err
	}

	for cursor.Next(c.Ctx){
		row := models.Item{}
		cursor.Decode(&row)
		items = append(items, row)
	}

	return items, nil
}

func (c *PostsClient) Search(filter interface{}) ([]models.Item, error){
	items := []models.Item{}
	if filter == nil{
		filter = bson.M{}
	}

	cursor, err := c.Col.Find(c.Ctx, filter)
	if err != nil{
		return items, err
	}

	for cursor.Next(c.Ctx){
		row := models.Item{}
		cursor.Decode(&row)
		items = append(items, row)
	}

	return items, nil
}