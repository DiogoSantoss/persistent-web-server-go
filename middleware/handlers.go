// Package middleware provides handlers to deal with API requests
// and stores data in a database.
package middleware

import (
	"gitlab.com/psem/recruitment-software/diogosantoss/persistent-web-server/models"

	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

// Decoder to transform form data into struct
var decoder = schema.NewDecoder()

// get_data handler returns all data from database
func GetData(w http.ResponseWriter, r *http.Request) {

	var data = getDataDatabase()

	// Concat all data into a single struct
	var allData models.AggregateData
	for _, d := range data {
		allData.Latitude = append(allData.Latitude, d.Latitude)
		allData.Longitude = append(allData.Longitude, d.Longitude)
		allData.Time = append(allData.Time, d.Time)
		allData.Speed = append(allData.Speed, d.Speed)
	}

	res := models.Response {
		ListData: &allData,
		Message: "Data retrieved successfully",
	}

	// send response
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("Failed to encode response with error: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	
	// log response
	log.Printf("Sent %v entries from database\n", len(data))
}

// put_data handler receives data from POST request
func PutDataPost(w http.ResponseWriter, r *http.Request) {

	var data models.Data

	err := r.ParseForm()
	if err != nil {
		log.Printf("Failed to parse form with error: %v\n", err)
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	err = decoder.Decode(&data, r.PostForm)
	if err != nil {
		log.Printf("Failed to decode form with error: %v\n", err)
		http.Error(w, "Failed to decode parameters", http.StatusInternalServerError)
		return
	}

	addDataDatabase(data)

	res := models.Response {
		Data: &data,
		Message: "Data added successfully",
	}

	// send respsonse
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		log.Printf("Failed to encode response with error: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// log response
	log.Printf("New entry %v added to database\n", data)
}

// put_data handler receives data from GET request
func PutDataGet(w http.ResponseWriter, r *http.Request) {

	var data models.Data

	err := decoder.Decode(&data, r.URL.Query())
	if err != nil {
		log.Printf("Failed to decode parameters with error: %v\n", err)
		http.Error(w, "Failed to decode parameters", http.StatusInternalServerError)
		return
	}

	addDataDatabase(data)

	res := models.Response {
		Data: &data,
		Message: "Data added successfully",
	}

	// send respsonse
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		log.Printf("Failed to encode response with error: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// log response
	log.Printf("New entry %v added to database\n", data)
}

// Dummy handler to test middleware
func Dummy(w http.ResponseWriter, r *http.Request) {
	
	res := "Hello world"

	// send response
	err := json.NewEncoder(w).Encode(&res)
	if err != nil {
		log.Printf("Failed to encode response with error: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// log response
	log.Printf("Fake endpoint\n")
}