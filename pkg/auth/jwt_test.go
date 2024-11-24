package auth_test

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yogamandayu/ohmytp/pkg/auth"
	"log"
	"strings"
	"testing"
	"time"
)

func TestJWT_Generate(t *testing.T) {
	jwt := auth.NewJWT("example_secret_key")
	token, err := jwt.Generate(map[string]interface{}{
		"sub": "example",
		"exp": time.Now().Add(2 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestJWT_ValidateToken(t *testing.T) {

	t.Run("validate token with valid secret", func(t *testing.T) {
		jwt := auth.NewJWT("example_secret_key")
		token, err := jwt.Generate(map[string]interface{}{
			"sub": "example",
			"exp": time.Now().Add(1 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})
		require.NoError(t, err)
		require.NotEmpty(t, token)

		claims, err := jwt.ValidateToken(token)
		require.NoError(t, err)
		assert.NotEmpty(t, claims)
	})

	t.Run("validate token with valid secret and check claims", func(t *testing.T) {
		jwt := auth.NewJWT("example_secret_key")
		token, err := jwt.Generate(map[string]interface{}{
			"sub": "example",
			"exp": time.Now().Add(1 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})
		require.NoError(t, err)
		require.NotEmpty(t, token)

		claims, err := jwt.ValidateToken(token)
		require.NoError(t, err)
		require.NotEmpty(t, claims)
		assert.Equal(t, "example", claims["sub"])
	})

	t.Run("validate token with invalid secret", func(t *testing.T) {
		jwt := auth.NewJWT("example_secret_key")
		token, err := jwt.Generate(map[string]interface{}{
			"sub": "example",
			"exp": time.Now().Add(1 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})
		require.NoError(t, err)
		require.NotEmpty(t, token)

		jwt.Secret = "changed_secret"
		_, err = jwt.ValidateToken(token)
		require.Error(t, err)
	})

	t.Run("validate token with tampered payload", func(t *testing.T) {
		jwt := auth.NewJWT("example_secret_key")
		token, err := jwt.Generate(map[string]interface{}{
			"sub": "example",
		})
		require.NoError(t, err)
		require.NotEmpty(t, token)

		newPayload := `{"sub":"tampered"}`
		b64 := base64.RawStdEncoding.EncodeToString([]byte(newPayload))
		log.Println(token)

		splitToken := strings.Split(token, ".")
		require.Len(t, splitToken, 3)

		splitToken[1] = b64
		token = splitToken[0] + "." + splitToken[1] + "." + splitToken[2]
		log.Println(token)

		_, err = jwt.ValidateToken(token)
		require.Error(t, err)
		log.Println(err)
	})
}
