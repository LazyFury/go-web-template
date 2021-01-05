package sha

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

// Sha1 Sha1
type Sha1 struct {
	IV  []byte `json:"iv"`
	Key []byte `json:"key"`
}

// EnCode Encode
func (s *Sha1) EnCode(str string) string {
	c, _ := aes.NewCipher(s.Key)
	strNew := []byte(str)

	cfb := cipher.NewCFBEncrypter(c, s.IV)
	ciphertext := make([]byte, len(strNew))
	cfb.XORKeyStream(ciphertext, strNew)
	// fmt.Printf("%s=>%x\n", strNew, ciphertext)
	return hex.EncodeToString(ciphertext)
}

// AesDecryptCFB DeCode
func (s *Sha1) AesDecryptCFB(str string) (decrypted string) {
	block, _ := aes.NewCipher(s.Key)
	encrypted, _ := hex.DecodeString(str)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}

	stream := cipher.NewCFBDecrypter(block, s.IV)
	stream.XORKeyStream(encrypted, encrypted)
	return fmt.Sprintf("%s", encrypted)
}
