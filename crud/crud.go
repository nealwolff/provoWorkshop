package crud

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/nealwolff/provoWorkshop/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Insert will insert a document into the database
func Insert(collection string, data interface{}, w http.ResponseWriter) (*mongo.InsertOneResult, error) {
	col := client.GetCollection(collection)

	result, err := col.InsertOne(context.TODO(), data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	return result, err

}

//GetOne gets a single document from the database
func GetOne(collection, ID string, w http.ResponseWriter) (ret []byte, err error) {

	IDobj, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	col := client.GetCollection(collection)

	filter := bson.M{
		"_id": IDobj,
	}

	rawData := bson.M{}

	err = col.FindOne(context.TODO(), filter).Decode(&rawData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ret, _ = json.Marshal(rawData)

	return

}
