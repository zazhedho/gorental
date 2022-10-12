package middleware

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/zazhedho/gorental/src/helpers"
)

/* Upload to local storage */
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

/* Upload to cloudinary */

func Cloudinary(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//**parse multipart/form-data
		if isErr := r.ParseMultipartForm(20 << 20); isErr != nil {
			helpers.New(isErr.Error(), 400, true).Send(w)
			return
		}

		file, handlerFile, err := r.FormFile("image")
		if err != nil {
			helpers.New(err.Error(), 400, true)
			return
		}
		defer file.Close()

		cntx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//**file validation
		checkType := handlerFile.Header.Get("Content-Type") == "image/jpeg" || handlerFile.Header.Get("Content-Type") == "image/jpg" || handlerFile.Header.Get("Content-Type") == "image/png"
		if !checkType {
			helpers.New(err.Error(), 400, true).Send(w)
			return
		}

		name := time.Now().Format("2006-01-02_15:04:05") + "_" + handlerFile.Filename

		cld, err1 := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUD_KEY"), os.Getenv("CLOUD_SEC"))
		if err1 != nil {
			helpers.New(err1.Error(), 400, true).Send(w)
			return
		}

		upload, err2 := cld.Upload.Upload(cntx, file, uploader.UploadParams{Folder: "image", PublicID: name})
		if err2 != nil {
			helpers.New(err2.Error(), 400, true).Send(w)
			return
		}

		helpers.New("success", 200, false).Send(w)

		ctx := context.WithValue(r.Context(), "imageName", upload.SecureURL)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
