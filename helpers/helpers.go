package helpers

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadImage(file multipart.File, id time.Time) (string, error) {
	cldCloudName, cldApiKey, cldApiSecret := os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET")

	cld, _ := cloudinary.NewFromParams(
		cldCloudName,
		cldApiKey,
		cldApiSecret,
	)
	ctx := context.Background()

	resp, err := cld.Upload.Upload(
		ctx,
		file,
		uploader.UploadParams{
			PublicID: fmt.Sprintf("docs/sdk/go/community_forum_id_%v", id),
		})

	if err != nil {
		return "", fmt.Errorf("CLOUDINARY_ERROR:- %s", err.Error())
	}

	return resp.SecureURL, nil
}
