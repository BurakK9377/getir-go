package main

import (
	"bytes"
	"encoding/json"
	"getir/internal/models/record"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodNotAllowedForRecordServerForGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	recordServer(rr, req)
	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
	assert.JSONEq(t, `{"message": "Method Not Allowed"}`, string(resBody))
}

func TestBadRequestForRecordServerForPostWithWrongDateFormat(t *testing.T) {
	data, err := json.Marshal(record.Request{
		StartDate: "10-1992-10",
		EndDate:   "2022-12-10",
		MinCount:  10,
		MaxCount:  1000,
	})
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	req := httptest.NewRequest(http.MethodPost, "/", reader)
	rr := httptest.NewRecorder()

	recordServer(rr, req)
	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestMethodNotAllowedForRecordServerForPut(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rr := httptest.NewRecorder()

	recordServer(rr, req)
	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
	assert.JSONEq(t, `{"message": "Method Not Allowed"}`, string(resBody))
}

func TestMethodNotAllowedForRecordServerForDelete(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rr := httptest.NewRecorder()

	recordServer(rr, req)
	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
	assert.JSONEq(t, `{"message": "Method Not Allowed"}`, string(resBody))
}
