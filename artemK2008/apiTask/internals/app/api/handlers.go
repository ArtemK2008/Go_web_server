package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *API) grab(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	var data Data
	err := json.NewDecoder(req.Body).Decode(&data)
	fmt.Println(data)
	if err != nil {
		msg := Message{
			StatusCode: 400,
			Message:    "Privvided JSON invalid",
		}
		json.NewEncoder(writer).Encode(msg)
		writer.WriteHeader(400)
		return
	}
	writer.WriteHeader(200)
	fmt.Println(data)
	json.NewEncoder(writer).Encode(data)
	api.data = data
}

func (api *API) solve(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	fmt.Println(api.data)
	rootsCount := countRoots(api.data.A, api.data.B, api.data.C)
	msg := Message{
		StatusCode: 200,
		Message:    fmt.Sprintf("Number of roots is %d", rootsCount),
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(msg)
}
