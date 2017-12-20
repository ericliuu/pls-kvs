package common

import (
)

type ApiRequest struct {
	ReqType int
	Id int
	Key string
	Value string
}

type ApiResponse struct {
	ReqType int
	Id int
	Key string
	Value string
    ResCode int
}

const (
    FAIL int = -1
    NOTFOUND int = 0
    OK int = 1
)

func NewApiRequest(reqType int, id int, key string, value string) ApiRequest {
    apiReq := ApiRequest {
        ReqType: reqType,
        Id: id,
        Key: key,
        Value: value,
    }
    return apiReq
}

func NewApiResponse(resType int, id int, key string, value string, code int) ApiResponse {
    apiRes := ApiResponse {
        ReqType: resType,
        Id: id,
        Key: key,
        Value: value,
        ResCode: code,
    }
    return apiRes
}
