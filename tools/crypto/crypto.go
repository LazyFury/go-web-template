package crypto

import (
	"crypto/sha256"
	"fmt"
)

func SHA256String(data string) string {
	_data_b := []byte(data)
	hash := sha256.Sum256(_data_b)
	return fmt.Sprintf("%x", hash[:])
}
