package api

import (
	"fmt"
	"github.com/ericliuu/pls-kvs/common"
)

func (api *APIThread) handleResponse(res common.ApiResponse) {
   	switch res.ReqType {
   	case GET:
   		displayGetResponse(res)
   	case PUT:
   		displayPutResponse(res)
   	case DELETE:
   		displayDeleteResponse(res)
	default:
		fmt.Printf("Invalid request method received from key-value store\n")
   	}
}

func displayGetResponse(res common.ApiResponse) {
	switch res.ResCode {
	case common.OK:
		fmt.Printf("Found pair: (%s, %s)\n", res.Key, res.Value)
	case common.NOTFOUND:
		fmt.Printf("No value found for key: %s\n", res.Key)
	case common.FAIL:
		fmt.Printf("Get operation failed for key: %s\n", res.Key)
	default:
		fmt.Printf("Error occured when handling request\n")
	}
}

func displayPutResponse(res common.ApiResponse) {
	switch res.ResCode {
	case common.OK:
		fmt.Printf("Added pair: (%s, %s)\n", res.Key, res.Value)
	case common.NOTFOUND: fallthrough;
	case common.FAIL:
		fmt.Printf("Put operation failed for pair: (%s, %s)\n", res.Key, res.Value)
	default:
		fmt.Printf("Error occured when handling request\n")
	}
}

func displayDeleteResponse(res common.ApiResponse) {
	switch res.ResCode {
	case common.OK:
		fmt.Printf("Deleted pair: (%s, %s)\n", res.Key, res.Value)
	case common.NOTFOUND:
		fmt.Printf("Pair not found: (%s, %s)\n", res.Key, res.Value)
	case common.FAIL:
		fmt.Printf("Delete operation failed for pair: (%s, %s)\n", res.Key, res.Value)
	default:
		fmt.Printf("Error occured when handling request\n")
	}
}
