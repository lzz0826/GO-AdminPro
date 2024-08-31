package test

import (
	"AdminPro/common/mongodb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
)

func TestMongoDB(t *testing.T) {
	mongoClient := mongodb.MongoClient
	collection := mongoClient.Database("test").Collection("testCollection")
	res, err := collection.InsertOne(context.Background(), bson.M{"a": "foo", "b": "bar"})
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(res)
}

func TestMongoDB2(t *testing.T) {
	// 获取 MongoDB 客户端
	client := mongodb.MongoClient
	if client == nil {
		log.Fatal("MongoClient is not initialized")
	}

	// 获取集合
	collection := client.Database("test").Collection("testCollection")

	// 插入一个文档
	newDoc := bson.M{"a": "foo", "b": "bar"}
	insertResult, err := collection.InsertOne(context.TODO(), newDoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)

	// 查找一个文档
	var result bson.M
	err = collection.FindOne(context.TODO(), bson.M{"a": "foo"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found document: %v\n", result)

	// 更新一个文档
	updateResult, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"a": "foo"},
		bson.M{"$set": bson.M{"b": "updated"}},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// 删除一个文档
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"a": "foo"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents\n", deleteResult.DeletedCount)
}
