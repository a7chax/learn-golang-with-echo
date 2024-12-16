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

type BaseResponsePagination[T any] struct {
	Message   string   `json:"message"`
	Data      *[]T     `json:"data"`
	IsSuccess bool     `json:"isSuccess"`
	Metadata  Metadata `json:"metadata"`
}

type BaseResponsePaginationNoData struct {
	Message   string   `json:"message"`
	IsSuccess bool     `json:"isSuccess"`
	Metadata  Metadata `json:"metadata"`
}
