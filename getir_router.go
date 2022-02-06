package main

import (
	"getir/internal/service"
	"log"
	"net/http"
	"os"
)

func recordServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "Method Not Allowed"}`))
	case "POST":
		service.GetRecords(w, r)
	case "PUT":
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "Method Not Allowed" }`))
	case "DELETE":
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "Method Not Allowed"}`))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "Method Not Allowed"}`))
	}
}

func inMemoryServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		service.GetInMemoryKey(w, r)
	case "POST":
		service.SetInMemoryKey(w, r)
	case "PUT":
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "Method Not Allowed"}`))
	case "DELETE":
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "Method Not Allowed"}`))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "Method Not Allowed"}`))
	}
}

func main() {
	http.HandleFunc("/", recordServer)
	http.HandleFunc("/in-memory", inMemoryServer)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
