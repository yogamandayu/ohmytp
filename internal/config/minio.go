package config

import (
	"github.com/yogamandayu/ohmytp/pkg/minio"
)

// MinioConfig is minio config.
type MinioConfig struct {
	*minio.Config
}
