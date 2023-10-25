package block

import (
	"crypto/aes"
)

// NewAESCipher creates a new AES cipher with the specified key.
// It returns an instance of BlockCipher that can be used for encryption
// and decryption.
func NewAESCipher(key []byte) (*BlockCipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return &BlockCipher{block: block}, nil
}
