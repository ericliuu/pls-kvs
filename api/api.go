package api

import (
	"fmt"
	"net/http"
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

    req := common.NewApiRequest(GET, 23, "foo", "bar")
    api.ApiReqChan <- req
	//setupService()

}

func setupService() {
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe(":54329", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {

}
