package persistence

import (
	aggregates "catalog/internals/domain/aggregates"
	valueObjects "catalog/internals/domain/valueObjects"
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

func GetCharacters() []aggregates.Character {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := createConnection(ctx)
	cursor, err := client.Database("catalog").Collection("characters").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	var characters []aggregates.Character
	for cursor.Next(ctx) {
		var character aggregates.Character
		if err = cursor.Decode(&character); err != nil {
			log.Fatal(err)
		}
		characters = append(characters, character)
	}
	closeConnection(client, ctx)
	return characters
}

func CreateCharacter(id string, mangaId string, localizations []aggregates.CharacterLocalization,
	class aggregates.Class, personalityId string, leadershipId string,
	rarity valueObjects.Rarity, level int, isPackable bool) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := createConnection(ctx)
	character := aggregates.NewCharacter(id, mangaId, localizations, class, personalityId, leadershipId, rarity, level, isPackable)
	_, err := client.Database("catalog").Collection("characters").InsertOne(ctx, character)
	if err != nil {
		log.Fatal(err)
	}
	closeConnection(client, ctx)
	return err
}
