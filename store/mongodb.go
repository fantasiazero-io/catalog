package store

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connestionString = "mongodb://localhost:27017"

func createConnection(ctx context.Context) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(connestionString))

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func closeConnection(client *mongo.Client, ctx context.Context) {
	error := client.Disconnect(ctx)
	if error != nil {
		log.Fatal(error)
	}
}

func GetProducts() []Product {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := createConnection(ctx)
	cursor, err := client.Database("catalog").Collection("products").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	var products []Product
	for cursor.Next(ctx) {
		var product Product
		if err = cursor.Decode(&product); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	closeConnection(client, ctx)
	return products
}

func CreateProduct(id string, title string, image string, order int) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := createConnection(ctx)
	product := Product{
		Id:    id,
		Title: title,
		Image: image,
		Order: order,
	}
	_, err := client.Database("catalog").Collection("products").InsertOne(ctx, product)
	if err != nil {
		log.Fatal(err)
	}
	closeConnection(client, ctx)
	return err
}

type Product struct {
	Id    string
	Title string
	Image string
	Order int
}
