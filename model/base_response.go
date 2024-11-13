package model

type BaseResponse[T any] struct {
	Message   string `json:"message"`
	Data      *T     `json:"data"`
	IsSuccess bool   `json:"isSuccess"`
}

type BaseResponseNoData struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
}
