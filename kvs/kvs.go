package kvs

import (
	"fmt"
	"time"
	"github.com/ericliuu/pls-kvs/api"
	"github.com/ericliuu/pls-kvs/common"
)

type KVSThread struct {
	ApiThreadPtr *api.APIThread
    KVStore *Node
}

// constructor
func NewKVSThread(apiThread *api.APIThread) KVSThread {
	kvsThread := KVSThread {
        ApiThreadPtr: apiThread,
        KVStore: NewNode(),
	}
    return kvsThread
}

// returns value for a given key
func (kvs *KVSThread) handleRequestGet(req common.ApiRequest) { 

    // checks if key is not in KVS
    code := common.OK
    if ! kvs.KVStore.inStore(req.Key) {
        code = common.NOTFOUND
    }

    response := common.ApiResponse {
        ReqType: api.GET,
        Id: 1,
        Key: req.Key,
        Value: kvs.KVStore.getValue(req.Key),
        ResCode: code,
    }
    //fmt.Printf("Received: GET (%s, %s)\n", response.Key, response.Value)
    kvs.ApiThreadPtr.ApiResChan <- response
}

// adds key and value to store
func (kvs *KVSThread) handleRequestPut(req common.ApiRequest) {
    kvs.KVStore.putValue(req.Key, req.Value)
    //fmt.Printf("Received: PUT (%s, %s)\n", req.Key, req.Value)
}

// deletes a key and value from store
func (kvs *KVSThread) handleRequestDelete(req common.ApiRequest) {
    // checks if key is not in KVS
    code := common.OK
    if ! kvs.KVStore.inStore(req.Key) {
        code = common.NOTFOUND
    }

    response := common.ApiResponse {
        ReqType: api.DELETE,
        Id: 1,
        Key: req.Key,
        Value: kvs.KVStore.getValue(req.Key),
        ResCode: code,
    }
    kvs.KVStore.delValue(req.Key)
    kvs.ApiThreadPtr.ApiResChan <- response
}

// calls method that performs request
func (kvs *KVSThread) handleRequest(req common.ApiRequest) {
    switch req.ReqType {
    case api.GET:
        kvs.handleRequestGet(req)
    case api.PUT:
        kvs.handleRequestPut(req)
    case api.DELETE:
        kvs.handleRequestDelete(req)
    default:
        fmt.Printf("Something went wrong !!!!!!!!!!!!")
    }
}

func (kvs *KVSThread) Exec() {
	fmt.Printf("KVS thread started...\n")

	time.Sleep(100 * time.Millisecond)

    for {
        //select {
        req := <-kvs.ApiThreadPtr.ApiReqChan
        kvs.handleRequest(req)

    }
}
