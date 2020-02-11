package crud

import (
	"context"
	"net/http"

	"github.com/nealwolff/provoWorkshop/client"
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
