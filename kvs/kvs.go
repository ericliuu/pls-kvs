package kvs

import (
	"fmt"
	"time"
)

type KVSThread struct {

}

func (kvs *KVSThread) Exec() {
	fmt.Printf("KVSTHREAD!!\n")

	time.Sleep(100 * time.Millisecond)
}
