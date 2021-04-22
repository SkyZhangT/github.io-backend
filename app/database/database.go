package database

import (
	"context"
	"fmt"
	"log"

	"app/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostDB struct{
	Client mongo.Client
	DB 	mongo.Database
	Col mongo.Collection
	Ctx context.Context
}

func Initdb(ctx context.Context, conf config.MongoConfiguration) *PostDB{
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Server))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	galleryDatabase := client.Database(conf.Database)
	postCollection := galleryDatabase.Collection(conf.Collection)
	
	return &PostDB{Client: *client, DB: *galleryDatabase, Col: *postCollection, Ctx: ctx}
}

 
func (p *PostDB) Printdb(){
	databases, err := p.Client.ListDatabaseNames(p.Ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}

func (p *PostDB) Close(){
	err := p.Client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
