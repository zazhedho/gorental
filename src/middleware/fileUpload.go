package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/zazhedho/gorental/src/helpers"
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

// func FileUpload(role string, next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-type", "application/json")

// 		var pathfile string

// 		err := r.ParseForm()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		if err := r.ParseMultipartForm(32 << 20); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		for _, fheaders := range r.MultipartForm.File {
// 			for _, hdr := range fheaders {
// 				infile, err := hdr.Open()
// 				if err != nil {
// 					fmt.Fprint(w, err.Error())
// 				}

// 				pathfile = "/tmp/" + hdr.Filename
// 				outfile, err := os.Create(pathfile)
// 				if err != nil {
// 					fmt.Fprint(w, err.Error())
// 				}

// 				defer outfile.Close()
// 				io.Copy(outfile, infile)
// 			}
// 		}

// 		log.Println("Upload Middleware Pass")
// 		// share context to controller
// 		ctx := context.WithValue(r.Context(), "file", pathfile)

// 		// Serve the next handler
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func Cloudinary(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//**parse multipart/form-data
		if err := r.ParseMultipartForm(20 << 20); err != nil {
			helpers.New(err.Error(), 400, true).Send(w)
			return
		}

		file, handlerFile, err := r.FormFile("image")

		defer file.Close()

		if err != nil {
			helpers.New(err.Error(), 400, true).Send(w)
			return
		}

		cntx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//**file validation
		checkType := handlerFile.Header.Get("Content-Type") == "image/jpeg" || handlerFile.Header.Get("Content-Type") == "image/jpg" || handlerFile.Header.Get("Content-Type") == "image/png"

		if !checkType {
			helpers.New(err.Error(), 400, true).Send(w)
			return
		}

		name := strings.ReplaceAll(strings.ReplaceAll(time.Now().Format(time.ANSIC), ":", "-")+"-"+handlerFile.Filename, " ", "_")

		cld, errs := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUD_KEY"), os.Getenv("CLOUD_SEC"))

		if errs != nil {
			helpers.New(errs.Error(), 400, true).Send(w)
			return
		}

		upload, err2 := cld.Upload.Upload(cntx, file, uploader.UploadParams{Folder: "image", PublicID: name})

		if err2 != nil {
			helpers.New(err2.Error(), 400, true).Send(w)
			return
		}

		helpers.New("success", 400, true).Send(w)

		ctx := context.WithValue(r.Context(), "imageName", upload.SecureURL)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
