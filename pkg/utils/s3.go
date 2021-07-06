package utils

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

type S3Client struct {
	sess *session.Session
	svc  *s3.S3
}

// NewS3Client returns a new S3Client.
func NewS3Client() (*S3Client, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	return &S3Client{sess: sess, svc: svc}, nil
}

// DownloadFile from a bucket and key to a destination file.
func (client S3Client) DownloadFile(bucket, key string, dest *os.File) error {
	downloader := s3manager.NewDownloader(client.sess)

	_, err := downloader.Download(dest, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	return err
}

// GetContent from a bucket and key as string.
func (client S3Client) GetContent(bucket, key string) (string, error) {
	raw, err := client.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	_, err = buf.ReadFrom(raw.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
