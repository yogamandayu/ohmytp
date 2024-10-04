package util_test

import (
	"github.com/yogamandayu/ohmytp/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootDir(t *testing.T) {
	rootDir := util.RootDir()
	assert.NotEmpty(t, rootDir)
}
