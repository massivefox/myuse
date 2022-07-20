package github.com/massivefox/myuse
import (
	"context"
	"crud/config"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	credential := options.Credential{
		Username: "root",
		Password: "root",
	}

	connectionURI := fmt.Sprintln(config.ConnectURI)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI).SetAuth(credential))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.ConnectTimeout*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}
	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}
