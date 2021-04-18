package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.io-backend/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type PostDB struct{
	client mongo.Client
	db 	mongo.Database
	col mongo.Collection
}



func Initdb(conf config.MongoConfiguration) *PostDB{
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Server))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	database := client.Database(conf.Database)
	collection := database.Collection(conf.Collection)


	return &PostDB{client: *client, db: *database, col: *collection}
}

 
func (p *PostDB) Printdb(){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	databases, err := p.client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}

func (p *PostDB) Close(){
	err := p.client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
