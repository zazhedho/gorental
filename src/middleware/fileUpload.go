package middleware

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// func FileUpload(role string, next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		file, fileHeader, err := r.FormFile("image")
// 		if err != nil {
// 			helpers.New(err.Error(), 400, true)
// 			return
// 		}
// 		defer file.Close()

// 		fileName := "tmp/" + time.Now().Format("2006-01-02_15:04:05") + "_" + fileHeader.Filename
// 		fileDestination, err := os.Create(fileName)
// 		if err != nil {
// 			helpers.New(err.Error(), 400, true)
// 			return
// 		}

// 		_, err = io.Copy(fileDestination, file)
// 		if err != nil {
// 			helpers.New(err.Error(), 400, true)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), "imgName", fileName)
// 		next.ServeHTTP(w, r.WithContext(ctx))

// 	}
// }

func FileUpload(role string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		var pathfile string

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := r.ParseMultipartForm(32 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, fheaders := range r.MultipartForm.File {
			for _, hdr := range fheaders {
				infile, err := hdr.Open()
				if err != nil {
					fmt.Fprint(w, err.Error())
				}

				pathfile = "/tmp/" + hdr.Filename
				outfile, err := os.Create(pathfile)
				if err != nil {
					fmt.Fprint(w, err.Error())
				}

				defer outfile.Close()
				io.Copy(outfile, infile)
			}
		}

		log.Println("Upload Middleware Pass")
		// share context to controller
		ctx := context.WithValue(r.Context(), "file", pathfile)

		// Serve the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
