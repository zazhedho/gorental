package middleware

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func FileUpload(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "aplication/json")

		var pathfile string

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, headers := range r.MultipartForm.File {
			for _, hdr := range headers {
				infile, err := hdr.Open()
				if err != nil {
					fmt.Fprint(w, err.Error())
				}

				pathfile = "src/images/" + hdr.Filename
				outfile, err := os.Create(pathfile)
				if err != nil {
					fmt.Fprint(w, err.Error())
				}

				defer outfile.Close()
				io.Copy(outfile, infile)
			}
		}

		log.Println("Upload middleware Pass")

		// share contex to controller
		ctx := context.WithValue(r.Context(), "file", pathfile)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
