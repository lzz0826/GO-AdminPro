package mongodb

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient *mongo.Client

func init() {
	// 连接 MongoDB
	url := viper.GetString("mongodb.url")
	clientOptions := options.Client().ApplyURI(url)

	// 创建自定义日志记录器
	mongoLogger := &MongoLogger{}

	// 配置日志记录选项
	loggerOption := options.
		Logger().
		SetSink(mongoLogger).
		SetMaxDocumentLength(25).
		SetComponentLevel(options.LogComponentCommand, options.LogLevelInfo)

	clientOptions.SetLoggerOptions(loggerOption)

	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("test").Collection("testCollection")
	_, err = collection.InsertOne(context.Background(), bson.M{"a": "foo", "b": "bar"})

}
