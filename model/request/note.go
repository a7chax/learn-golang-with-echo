package model_request

type Note struct {
	Title   string `json:"title" validate="required, min=5, max=100"`
	Content string `json:"content" validate="required, min=5, max=1000"`
}
