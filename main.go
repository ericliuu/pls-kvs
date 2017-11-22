package main

import (
	"fmt"
	"github.com/ericliuu/pls-kvs/api"
	"github.com/ericliuu/pls-kvs/kvs"
)

func main() {
	fmt.Printf("hello world\n")

	apiThread := api.APIThread{

	}
	
	kvsThread := kvs.KVSThread{

	}

	go apiThread.Exec()
	
	kvsThread.Exec()
}
