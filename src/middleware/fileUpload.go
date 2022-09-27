package middleware

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/zazhedho/gorental/src/helpers"
)

func FileUpload(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			helpers.New(err.Error(), 400, true)
			return
		}
		defer file.Close()

		fileName := time.Now().Format("2006-01-02_15:04:05") + "_" + fileHeader.Filename
		fileDestination, err := os.Create("src/images/" + fileName)
		if err != nil {
			helpers.New(err.Error(), 400, true)
			return
		}

		_, err = io.Copy(fileDestination, file)
		if err != nil {
			helpers.New(err.Error(), 400, true)
			return
		}

		ctx := context.WithValue(r.Context(), "imgName", "src/images/"+fileName)
		next.ServeHTTP(w, r.WithContext(ctx))

	}
}
