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