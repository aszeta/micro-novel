package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aszeta/micro-novel/account/account"
	"github.com/aszeta/micro-novel/account/config"
	"github.com/go-kit/kit/log"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	config := config.ProvideConfig("app.yaml")
	fmt.Println(config.App.Name)
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", config.App.Port, "caller", log.DefaultCaller)
	}

	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	r := account.NewHttpServer(account.NewService(ctx, setDatabase(config), setRedis(config)), logger)
	logger.Log("msg", "HTTP", "addr", config.App.Port)
	logger.Log("err", http.ListenAndServe(config.App.Port, r))
}

func setDatabase(config *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func setRedis(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
