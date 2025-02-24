package models

type ServeFileModel struct {
	Path string	`json:"path"`
	ContentType string `json:"content_type"`
}