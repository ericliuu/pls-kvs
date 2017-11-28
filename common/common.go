package common

import (
)

type Request struct {
	ReqType int
	Id int
	Key string
	Value string
}

type Return struct {
	RetType int
	Id int
	Key string
	Value string
}