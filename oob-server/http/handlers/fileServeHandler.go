package handlers

import (
	"log"
	"net/http"

	"github.com/ind-exe/pulse/models"
)

func FileServeHandler(w http.ResponseWriter, r *http.Request, fileModel models.ServeFileModel) {
	w.Header().Set("Content-Type", fileModel.ContentType)
	http.ServeFile(w ,r, fileModel.Path)
	log.Printf("FileServe: the file %s served on %s", fileModel.Path, r.Host+r.URL.String())
}