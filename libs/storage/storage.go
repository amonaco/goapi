package storage

import (
	"log"
	"path/filepath"
	"time"

	"github.com/amonaco/goapi/libs/config"
	"github.com/minio/minio-go"
)

var client *minio.Client

func Setup() {
	conf := config.Get()

	// Initialize minio client object
	minioClient, err := minio.New(conf.Storage.Endpoint, conf.Storage.AccessKey, conf.Storage.SecretKey, false)
	if err != nil {
		log.Fatalln(err)
	}

	client = minioClient
}

// Returns a client instance
func Get() *minio.Client {
	return client
}

// Creates a Pre-signed URL
func PreSignURL(object string, path string) (string, error) {

	conf := config.Get()
	signatureExpiry := time.Duration(conf.Storage.SignatureExpiry) * time.Second

	url, err := client.PresignedPutObject(conf.Storage.Bucket, filepath.Join(object, path), signatureExpiry)
	if err != nil {
		return "", err
	}

	return url.String(), nil
}
