package crud

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nealwolff/provoWorkshop/client"
	"github.com/nealwolff/provoWorkshop/types"
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

//GetAll gets all the objects in a database
func GetAll(collection string, w http.ResponseWriter) (*mongo.Cursor, error) {

	col := client.GetCollection(collection)

	cursor, err := col.Find(context.TODO(), bson.M{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	return cursor, err

}

//GetAllUsers gets all the users from the user collection
func GetAllUsers(w http.ResponseWriter) ([]types.User, error) {
	var users []types.User

	cursor, err := GetAll("users", w)

	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return users, err
	}

	if len(users) < 1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No users found"))
	}
	return users, err
}

//Update updates a doument in the database
func Update(collection, ID string, data interface{}, w http.ResponseWriter) (*mongo.UpdateResult, error) {
	IDobj, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return nil, err
	}

	col := client.GetCollection(collection)

	filter := bson.M{
		"_id": IDobj,
	}

	opts := options.Replace().SetUpsert(true)

	result, err := col.ReplaceOne(context.TODO(), filter, data, opts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	return result, err
}
