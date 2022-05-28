package mp

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func MongoClient() *mongo.Client {
	clientOptions := options.Client().
		ApplyURI(Config.MongoDB.URI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func RecordSendMsg(client *mongo.Client, user string, receive string, reply string) {
	type ChatRecord struct {
		User     string
		Receive  string
		Reply    string
		CreateAt time.Time
	}
	fmt.Printf("[RecordSendMsg]%s -> %s\n", receive, reply)
	collection := client.Database("vqchat").Collection("chatrecord")
	_, err := collection.InsertOne(context.TODO(), ChatRecord{
		User:     user,
		Receive:  receive,
		Reply:    reply,
		CreateAt: time.Now(),
	})
	if err != nil {
		log.Fatal(err)
	}
}
