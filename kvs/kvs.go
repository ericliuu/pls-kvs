package kvs

import (
	"fmt"
	"time"
	"github.com/ericliuu/pls-kvs/api"
)

type KVSThread struct {
	ApiThreadPtr *api.APIThread
}

func NewKVSThread(apiThread *api.APIThread) KVSThread {
	kvsThread := KVSThread{
       ApiThreadPtr: apiThread,
	}
    return kvsThread
}

func (kvs *KVSThread) Exec() {
	fmt.Printf("KVSTHREAD!!\n")

	time.Sleep(100 * time.Millisecond)

    req := <-kvs.ApiThreadPtr.ApiReqChan
    fmt.Printf("%s: %s\n", req.Key, req.Value)
}
