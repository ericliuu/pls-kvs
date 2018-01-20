package api

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/ericliuu/pls-kvs/common"
)

const (
	GET int = 0
	PUT int = 1
	DELETE  int = 2
)

type APIThread struct {
	ApiReqChan chan common.ApiRequest
	ApiResChan chan common.ApiResponse
}

func NewAPIThread() APIThread {
	// Create the request and response channels to communicate with the KVS thread
	apiThread := APIThread {
		ApiReqChan: make(chan common.ApiRequest, 10),
		ApiResChan: make(chan common.ApiResponse, 10),
	}
	return apiThread
}


func (api *APIThread) Exec() {
	fmt.Printf("API thread started...\n")

	// Set up web server
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/key/", api.httpHandler)

	go http.ListenAndServe(":54321", nil)

	for {
		select {
			case res := <- api.ApiResChan:
				api.handleResponse(res)
		}
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid URI", 400)
}

func (api *APIThread) httpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Retrieve the corresponding value for the key
		key := getKeyFromRequest(w, r)

		req := common.NewApiRequest(GET, 1, key, "")
		api.ApiReqChan <- req

	case "PUT":
		// Store new key-value
		key := getKeyFromRequest(w, r)

		numValues := 0

		for value := range r.Form {
			numValues++
			if numValues > 1 {
				http.Error(w, "Received too many values", 400)
				break
			}

			req := common.NewApiRequest(PUT, 1, key, value)
			api.ApiReqChan <- req
		}

	case "DELETE":
		// Delete an existing key-value
		key := getKeyFromRequest(w, r)

		req := common.NewApiRequest(DELETE, 1, key, "")
		api.ApiReqChan <- req

	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func getKeyFromRequest(w http.ResponseWriter, r *http.Request) string {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing request", 400)
	}
	
	// Retreive the key string from the URI	
	key := strings.TrimLeft(r.URL.Path, "/key/")
	if len(key) == 0 {
		key = "/"
	}
	return key
}
