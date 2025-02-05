package lib

import (
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func GetCloudinaryConfig() *cloudinary.Cloudinary {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cloudKey := os.Getenv("CLOUDINARY_API_KEY")
	cloudSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cld, err := cloudinary.NewFromParams(cloudName, cloudKey, cloudSecret)
	if err != nil {
		panic(err)
	}
	return cld
}
