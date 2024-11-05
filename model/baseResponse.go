package model

type BaseResponse[T any] struct {
	Message   string `json:"message"`
	Data      *T     `json:"data"`
	IsSuccess bool   `json:"isSuccess"`
}
