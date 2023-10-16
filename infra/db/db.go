package db

import (
	"fmt"
	"context"
	"errors"

	"gopkg.in/mgo.v2/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/swarajkumarsingh/ziplink/conf"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/swarajkumarsingh/ziplink/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/swarajkumarsingh/ziplink/functions/general"
)

var collection *mongo.Collection

var isConnectDB = false
func Init() {

  if isConnectDB {
    return
  }

  clientOptions := options.Client().ApplyURI(conf.MONGO_CONNECTION_URL)
  client, err := mongo.Connect(context.TODO(), clientOptions)
  if err != nil {
    isConnectDB = false
    fmt.Printf("Error while connecting to DB, Error: %s", err.Error())
    panic(err)
  }
  
  isConnectDB = true

  collection = client.Database(conf.DbName).Collection(conf.ColName)

  indexModel := mongo.IndexModel{
    Keys: bson.M{"shortUrl": 1},
  }

  _, err = collection.Indexes().CreateOne(context.Background(), indexModel)
  if err != nil {
    fmt.Println("Unable to add indexes", err)
  }

  fmt.Println("Connected to DB successfully")
}

func InsertUrl(model model.UrlModel) (string, error) {

  if !general.IsValidURL(model.LongUrl) {
    return "Invalid url", errors.New("invalid url")
  }

  fmt.Println(model)

  inserted, err := collection.InsertOne(context.TODO(), model)

  if err != nil {
    return "error while inserting document in DB", err
  }

  fmt.Println("Inserted link in DB with id: ", inserted.InsertedID)
  return "Success", nil
}

func UpdateUrl(urlId string) {
  // convert string to _id
  id, err := primitive.ObjectIDFromHex(urlId)
  if err != nil {
    panic(err)
  }

  filter := bson.M{"_id": id}
  update := bson.M{"$set": bson.M{"watched": true}}
  result, err := collection.UpdateOne(context.Background(), filter, update)

  if err != nil {
    panic(err)
  }

  fmt.Println("modified count: ", result.ModifiedCount)
}

func DeleteUrl(urlId string) {
  id, err := primitive.ObjectIDFromHex(urlId)
  if err != nil {
    panic(err)
  }

  filter := bson.M{"_id": id}
  result, err := collection.DeleteOne(context.Background(), filter)

  if err != nil {
    panic(err)
  }

  fmt.Println("deleted count: ", result.DeletedCount)
}

func DeleteAllUrl() {
  filter := bson.D{{}}
  result, err := collection.DeleteMany(context.Background(), filter, nil)

  if err != nil {
    panic(err)
  }

  fmt.Println("deleted count: ", result.DeletedCount)
}

func FindOne(shortId string) (model.UrlModel, error) {
  filter := bson.M{"shortId": shortId}
  var result model.UrlModel
  err := collection.FindOne(context.TODO(), filter).Decode(&result)

  if err != nil {
    if err == mongo.ErrNoDocuments {
      return result, err
    }
    return result, err
  }

  return result, nil
}
