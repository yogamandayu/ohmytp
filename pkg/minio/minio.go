package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Host            string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func NewClient(config *Config) (*minio.Client, error) {
	client, err := minio.New(config.Host, &minio.Options{
		Creds: credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
