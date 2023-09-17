package db

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swarajkumarsingh/ziplink/conf"
	"github.com/swarajkumarsingh/ziplink/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var collection *mongo.Collection
func Init() {
  clientOptions := options.Client().ApplyURI(conf.ConnectionString)

  client, err := mongo.Connect(context.TODO(), clientOptions)

  if err != nil {
    panic(err)
  }

  fmt.Println("Connected to DB successfully")
  collection = client.Database(conf.DbName).Collection(conf.ColName)
}

func InsertUrl(c *gin.Context, model model.UrlModel) {
  if err := c.BindJSON(&model); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  inserted, err := collection.InsertOne(context.Background(), model)

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  fmt.Println("Inserted 1 movie in DB with id: ", inserted.InsertedID)
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
  // convert string to _id
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

func GetAllUrlsList() {
  cur, err := collection.Find(context.Background(), bson.D{{}})

  if err != nil {
    panic(err)
  }

  var urlList []primitive.M

  for cur.Next(context.Background()) {
    var model bson.M
    err := cur.Decode(&model)
    if err != nil {
      panic(err)
    }

    urlList = append(urlList, primitive.M(model))
  }

  defer cur.Close(context.Background())
}
