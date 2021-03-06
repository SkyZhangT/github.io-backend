package database

import (
	"fmt"

	"app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type DBInterface interface {
	Insert(models.Item) (models.Item, error)
	Delete(string) (int64, error)
	Get(string) (models.Item, error)
	GetPage(int64) ([]models.Item, error)
	Search(interface{}) ([]models.Item, error)
	Update(string, bson.D) (int64, error)
}


func (c *PostDB) Insert(data models.Item) (models.Item, error){
	item := models.Item{}

	res, err := c.Col.InsertOne(c.Ctx, data)
	if err != nil{
		fmt.Println(err)
		return item, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}

func (c *PostDB) Update(id string, filter bson.D) (int64, error){
	_id, _ := primitive.ObjectIDFromHex(id)

	_, err := c.Col.UpdateOne(
		c.Ctx,
		bson.M{"_id": _id},
		filter,
	)

	res, _ := c.Get(id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return res.Likes, nil
}

func (c *PostDB) Delete(id string) (int64, error){
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

func (c *PostDB) Get(id string) (models.Item, error){
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

func (c *PostDB) GetPage(pageNumber int64) ([]models.Item, error){
	items := []models.Item{}

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"time": -1})
	findOptions.SetLimit(10)
	findOptions.SetSkip(pageNumber*10)

	cursor, err := c.Col.Find(c.Ctx, bson.D{}, findOptions)
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

func (c *PostDB) Search(filter interface{}) ([]models.Item, error){
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

