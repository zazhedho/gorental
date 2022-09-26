package helpers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/zazhedho/gorental/src/database/orm/models"
)

// Local directory
func Upload(data *models.Vehicle, w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		New(err.Error(), 400, true)
		return
	}

	imgName := time.Now().Format("2006-01-02-15:04:05") + "-" + fileHeader.Filename
	fileDestination, err := os.Create("src/images/" + imgName)
	if err != nil {
		New(err.Error(), 400, true)
		return
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		New(err.Error(), 400, true)
		return
	}

	data.Image = imgName
}
