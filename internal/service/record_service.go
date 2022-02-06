package service

import (
	"context"
	"encoding/json"
	"getir/internal/models/record"
	"getir/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

var collection = store.ConnectDB()

const RFC3339FullDate = "2006-01-02"

func GetRecords(w http.ResponseWriter, r *http.Request) {

	var recordRequest record.Request

	// decode recordRequest or return error
	err := json.NewDecoder(r.Body).Decode(&recordRequest)
	if err != nil {
		getError(err, w, http.StatusBadRequest)
		return
	}
	formattedStartDate, err := time.Parse(RFC3339FullDate, recordRequest.StartDate)
	if err != nil {
		getError(err, w, http.StatusBadRequest)
		return
	}

	formattedEndDate, err := time.Parse(RFC3339FullDate, recordRequest.EndDate)
	if err != nil {
		getError(err, w, http.StatusBadRequest)
		return
	}
	// we created Record array
	var records []record.Record

	dateFilter := bson.D{
		{"$match", bson.D{
			{"createdAt", bson.D{
				{"$gte", formattedStartDate},
				{"$lte", formattedEndDate},
			}},
		}},
	}

	sumStage := bson.D{
		{"$project", bson.D{
			{"key", 1},
			{"createdAt", 1},
			{"_id", 0},
			{"totalCount", bson.D{{"$sum", "$counts"}}},
		}},
	}

	totalCountFilter := bson.D{
		{"$match", bson.D{
			{"totalCount", bson.D{
				{"$lte", recordRequest.MaxCount},
				{"$gte", recordRequest.MinCount},
			}},
		}},
	}

	stagePipeline := mongo.Pipeline{dateFilter, sumStage, totalCountFilter}
	cur, err := collection.Aggregate(context.TODO(), stagePipeline)

	if err != nil {
		getError(err, w, http.StatusInternalServerError)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			getError(err, w, http.StatusInternalServerError)
			return
		}
	}(cur, context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var record record.Record
		// & character returns the memory address of the following variable.
		err := cur.Decode(&record) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		records = append(records, record)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	getRecordResponse(records, w)
}

func getRecordResponse(records []record.Record, w http.ResponseWriter) {
	recordResponse := record.Response{
		Code:    0,
		Msg:     "Success",
		Records: records,
	}

	message, _ := json.Marshal(recordResponse)

	w.WriteHeader(http.StatusOK)
	w.Write(message)
}
