package service

import (
	"encoding/json"
	memory2 "getir/internal/models/memory"
	"getir/internal/store"
	"net/http"
	"time"
)

var redisClient = store.Initialize()

func SetInMemoryKey(w http.ResponseWriter, r *http.Request) {
	var inMemoryRequest memory2.InMemoryRequest
	// decode inMemoryRequest or return error
	err := json.NewDecoder(r.Body).Decode(&inMemoryRequest)
	if err != nil {
		getError(err, w, http.StatusBadRequest)
		return
	}

	err = redisClient.SetKey(inMemoryRequest.Key, inMemoryRequest.Value, time.Hour*24*30)
	if err != nil {
		getError(err, w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Key is saved"}`))
}

func GetInMemoryKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	var value string
	err := redisClient.GetKey(key, &value)
	if err != nil {
		getError(err, w, http.StatusNotFound)
		return
	}
	inMemoryResponse := memory2.InMemoryResponse{
		Key:   key,
		Value: value,
	}

	message, _ := json.Marshal(inMemoryResponse)

	w.WriteHeader(http.StatusAccepted)
	w.Write(message)
}
