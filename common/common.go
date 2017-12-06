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
	RespType int
	Id int
	Key string
	Value string
}

func NewApiRequest(reqType int, id int, key string, value string) ApiRequest {
    apiReq := ApiRequest {
        ReqType: reqType,
        Id: id,
        Key: key,
        Value: value,
    }
    return apiReq
}

func NewApiResponse(resType int, id int, key string, value string) ApiResponse {
    apiRes := ApiResponse {
        RespType: resType,
        Id: id,
        Key: key,
        Value: value,
    }
    return apiRes
}
