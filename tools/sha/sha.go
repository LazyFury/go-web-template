package sha

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

// Sha1 Sha1
type Sha1 struct {
	IV  string `json:"iv"`
	Key string `json:"key"`
}

// EnCode Encode
func (s *Sha1) EnCode(str string) string {
	c, err := aes.NewCipher([]byte(s.Key))
	if err != nil {
		panic(err)
	}
	strNew := []byte(str)

	cfb := cipher.NewCFBEncrypter(c, []byte(s.IV))
	ciphertext := make([]byte, len(strNew))
	cfb.XORKeyStream(ciphertext, strNew)
	// fmt.Printf("%s=>%x\n", strNew, ciphertext)
	return hex.EncodeToString(ciphertext)
}

// AesDecryptCFB DeCode
func (s *Sha1) AesDecryptCFB(str string) (decrypted string) {
	block, _ := aes.NewCipher([]byte(s.Key))
	encrypted, _ := hex.DecodeString(str)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}

	stream := cipher.NewCFBDecrypter(block, []byte(s.IV))
	stream.XORKeyStream(encrypted, encrypted)
	return fmt.Sprintf("%s", encrypted)
}
