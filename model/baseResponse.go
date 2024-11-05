package model

type BaseResponse[data any] struct {
	Message string `json:"message"`
	Data    data   `json:"data"`
}
