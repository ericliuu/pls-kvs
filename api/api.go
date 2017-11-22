package api

import (
	"fmt"
)

type APIThread struct {

}

func (api *APIThread) Exec() {
	fmt.Printf("APITHREAD!!\n")
}
