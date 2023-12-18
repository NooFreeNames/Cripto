package cipherfactory

import (
	"crypto/aes"

	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher"
	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher/block"
	"github.com/NooFreeNames/Cripto/pkg/crypto/key"
)

// AESFactory represents a factory for creating AES ciphers.
type AESFactory struct {
	keyGen key.IKeyGen
}

// Cipher creates an AES cipher using the provided password and salt.
func (bcs AESFactory) Cipher(password string, salt []byte) (cipher.ICipher, error) {
	aesBlock, err := aes.NewCipher(bcs.keyGen.Generate(password, aes.BlockSize))
	if err != nil {
		return block.BlockCipher{}, err
	}
	return block.NewBlockCipher(aesBlock, salt)
}

// NewAESFactory creates a new instance of AESFactory with the given keyGen.
func NewAESFactory(keyGen key.IKeyGen) AESFactory {
	return AESFactory{keyGen}
}
