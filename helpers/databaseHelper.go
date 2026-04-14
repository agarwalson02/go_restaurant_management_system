package helpers

import "go.mongodb.org/mongo-driver/mongo"

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}