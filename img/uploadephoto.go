package img

import (
	"io"
	"mime/multipart"
	"os"
	"time"
)

// Uploadimg use this function to upload photos
func Uploadimg(file *multipart.FileHeader, destination string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	// switch {
	// case destination == project:
	dst, err := os.Create("img/" + destination + "/" + time.Now().Format("2006-01-02-15-04-05") + file.Filename)
	if err != nil {
		return "", err
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return time.Now().Format("2006-01-02-15-04-05") + file.Filename, err
}
