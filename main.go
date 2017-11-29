package main

import (
	"fmt"
	"github.com/ericliuu/pls-kvs/api"
	"github.com/ericliuu/pls-kvs/kvs"
)

func main() {
	fmt.Printf("hello world\n")

	apiThread := api.NewAPIThread()

    kvsThread := kvs.NewKVSThread(&apiThread)

	go apiThread.Exec()

    kvsThread.Exec()
}
