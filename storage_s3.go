package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/url"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Storage struct {
	bucket string
	prefix string
}

func s3StorageFromURI(uri string) s3Storage {
	u, _ := url.Parse(uri)

	return s3Storage{
		bucket: u.Host,
		prefix: u.Path,
	}
}

// WriteFile takes the content of a file and writes it into the underlying
// storage system.
func (s s3Storage) WriteFile(filepath string, content io.Reader, contentType string) error {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	body, err := ioutil.ReadAll(content)
	if err != nil {
		return err
	}

	_, err = svc.PutObject(&s3.PutObjectInput{
		Body:        bytes.NewReader(body),
		Bucket:      aws.String(s.bucket),
		ContentType: aws.String(contentType),
		Key:         aws.String(path.Join(s.prefix, filepath)),
	})

	return err
}

// ReadFile retrieves a file from the underlying storage, needs to return
// errFileNotFound when file is not present.
func (s s3Storage) ReadFile(filepath string) (io.ReadCloser, error) {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	res, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path.Join(s.prefix, filepath)),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == s3.ErrCodeNoSuchKey {
			return nil, errFileNotFound
		}
		return nil, err
	}

	return res.Body, nil
}
