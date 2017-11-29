package kvs

import (
	"fmt"
	"time"
	"github.com/ericliuu/pls-kvs/api"
)

type KVSThread struct {
	ApiThreadPtr *api.APIThread
}

func (kvs *KVSThread) Exec() {
	fmt.Printf("KVSTHREAD!!\n")

	time.Sleep(100 * time.Millisecond)
}
