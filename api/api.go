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
	apiThread := APIThread {
		ApiReqChan: make(chan common.ApiRequest, 10),
		ApiResChan: make(chan common.ApiResponse, 10),
	}
	return apiThread
}


func (api *APIThread) Exec() {
	fmt.Printf("APITHREAD!!\n")

    // Set up web api
    http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/key/", api.httpHandler)

	go http.ListenAndServe(":54321", nil)

	for {
		res := <- api.ApiResChan
		api.handleResponse(res)
	}
}

func (api *APIThread) handleResponse(res common.ApiResponse) {
	// temporary
    fmt.Printf("Received RESPONSE (%s, %s)\n", res.Key, res.Value)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid URI", 400)
}

func (api *APIThread) httpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource
		key := getKeyFromRequest(w, r)

		req := common.NewApiRequest(GET, 1, key, "")
		api.ApiReqChan <- req		

	case "PUT":
		// Store new key-value
		key := getKeyFromRequest(w, r)

		if len(r.Form) > 1 {
			http.Error(w, "Received too many values", 400)
		} else if len(r.Form) < 1 {
			http.Error(w, "No value received", 400)
		}

		for value := range r.Form {
			req := common.NewApiRequest(PUT, 1, key, value)
			api.ApiReqChan <- req
		}

	case "DELETE":
		// Delete an existing key-value
		key := getKeyFromRequest(w, r)

		if len(r.Form) > 1 {
			http.Error(w, "Received too many values", 400)
		}

		for value := range r.Form {
			req := common.NewApiRequest(DELETE, 1, key, value)
			api.ApiReqChan <- req
		}

	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func getKeyFromRequest(w http.ResponseWriter, r *http.Request) string {
	r.ParseForm()
	key := strings.TrimLeft(r.URL.Path, "/key/")
	if len(key) == 0 {
		key = "/"
	}
	return key
}