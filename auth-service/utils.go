package main

//https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, err IError) {
	resp := IHttpResponse{
		body:       err,
		statusCode: code,
	}
	respondWithJSON(w, resp)
}

func respondWithJSON(w http.ResponseWriter, resp IHttpResponse) {
	response, _ := json.Marshal(resp.body)

	for headerName, headerValue := range resp.headers {
		w.Header().Set(headerName, headerValue)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.statusCode)
	w.Write(response)
}
