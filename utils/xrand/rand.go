package xrand

import (
	"github.com/gofrs/uuid"
	"strings"
)

func GenNonce() string {
	nonce, _ := uuid.NewV4()
	return strings.ReplaceAll(nonce.String(), "-", "")
}
