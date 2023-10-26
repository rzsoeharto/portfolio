package database

import (
	"fmt"
	"io"
	"os"
	"portfolio/server/responses"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func UploadImage(c *gin.Context) {

	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_KEY"))

	app, appErr := firebase.NewApp(c, nil, opt)
	if appErr != nil {
		fmt.Println("Firebase error: ", appErr)
		// logger.Logger.Println("Firebase error: ", appErr)
		return
	}

	client, storageErr := app.Storage(c)

	if storageErr != nil {
		fmt.Println("Firebase error: ", storageErr)
		// logger.Logger.Println("Firebase error: ", storageErr)
		return
	}

	// File path
	file, header, reqErr := c.Request.FormFile("file")

	if reqErr != nil {
		fmt.Println(reqErr)
		return
	}

	fmt.Println(header.Filename)

	remoteImagePath := "images/" + header.Filename

	bucket, bucketErr := client.Bucket("portfolio-project-6ac0e.appspot.com")
	if bucketErr != nil {
		fmt.Println("BUCKET error: ", bucketErr)
		// logger.Logger.Println("Firebase error: ", bucketErr)
		return
	}

	obj := bucket.Object(remoteImagePath)
	wc := obj.NewWriter(c)

	if _, copyErr := io.Copy(wc, file); copyErr != nil {
		fmt.Println("Copy error: ", copyErr)
		return
	}

	if writerErr := wc.Close(); writerErr != nil {
		fmt.Println("Close error: ", writerErr)
		return
	}

	responses.Code200(c, "OK")

}
