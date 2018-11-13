package libraries

import (
	"bytes"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jacky-htg/api-news/config"
)

func AwsUploadS3(img string, path string) error {
	i := strings.Index(img, ",")
	if i < 0 {
		return errors.New("Please suplay valid base64 image")
	}

	i2 := strings.Index(img, ";")
	if i2 < 0 {
		return errors.New("Please suplay valid base64 image")
	}

	image, err := base64.StdEncoding.DecodeString(img[i+1:])
	if err != nil {
		return err
	}

	creds := credentials.NewStaticCredentials(config.GetString("aws.s3.key"), config.GetString("aws.s3.secret"), "")
	_, err = creds.Get()

	if err != nil {
		return err
	}

	cfg := aws.NewConfig().WithRegion(config.GetString("aws.s3.region")).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(config.GetString("aws.s3.bucket")),
		Key:           aws.String(path),
		Body:          aws.ReadSeekCloser(bytes.NewReader(image)),
		ContentLength: aws.Int64(int64(len(image))),
		ContentType:   aws.String(img[5:i2]),
	}
	_, err = svc.PutObject(params)

	return err
}
