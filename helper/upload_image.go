package helper

import (
	"context"
	"ikan-cupang/lib"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(fileHeader *multipart.FileHeader) (string, error) {
	// params : cloud name, api key, api secret
	cld := lib.GetCloudinaryConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // timeout after 10s
	defer cancel()

	fileHandle, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()

	res, err := cld.Upload.Upload(ctx, fileHandle, uploader.UploadParams{Folder: "ikan-cupang"})
	if err != nil {
		return "", err
	}
	return res.SecureURL, nil
}
