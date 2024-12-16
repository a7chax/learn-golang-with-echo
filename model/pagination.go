package model

type Pagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type Metadata struct {
	TotalSize  int `json:"totalSize"`
	TotalPages int `json:"totalPages"`
	Page       int `json:"page"`
	Size       int `json:"size"`
}
