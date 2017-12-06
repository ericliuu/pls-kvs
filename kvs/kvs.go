package kvs

import (
	"fmt"
	"time"
	"github.com/ericliuu/pls-kvs/api"
	"github.com/ericliuu/pls-kvs/common"
)

type KVSThread struct {
	ApiThreadPtr *api.APIThread
    KVStore map[string]string
}

func NewKVSThread(apiThread *api.APIThread) KVSThread {
	kvsThread := KVSThread{
        ApiThreadPtr: apiThread,
        KVStore: make(map[string]string),
	}
    return kvsThread
}

// returns value for a given key
func (kvs *KVSThread) HandleRequestGet(req common.ApiRequest) {
    response := common.ApiResponse{
        RespType: api.GET,
        Id: 1,
        Key: req.Key,
        Value: kvs.KVStore[req.Key],
    }
    fmt.Printf("Received: GET (%s, %s)\n", response.Key, response.Value)
    kvs.ApiThreadPtr.ApiResChan <- response
}

// adds key and value to store
func (kvs *KVSThread) HandleRequestPut(req common.ApiRequest) {
    kvs.KVStore[req.Key] = req.Value
    fmt.Printf("Received: PUT (%s, %s)\n", req.Key, req.Value)
}

// deletes a key and value from store
func (kvs *KVSThread) HandleRequestDelete(req common.ApiRequest) {
    delete(kvs.KVStore, req.Key)
}
// calls method that performs request
func (kvs *KVSThread) HandleRequest(req common.ApiRequest) {
    switch req.ReqType {
    case api.GET:
        kvs.HandleRequestGet(req)
    case api.PUT:
        kvs.HandleRequestPut(req)
    case api.DELETE:
        kvs.HandleRequestDelete(req)
    }
}

func (kvs *KVSThread) Exec() {
	fmt.Printf("KVSTHREAD!!\n")

	time.Sleep(100 * time.Millisecond)

    for {
        //select {
        req := <-kvs.ApiThreadPtr.ApiReqChan
        kvs.HandleRequest(req)

    }
}
