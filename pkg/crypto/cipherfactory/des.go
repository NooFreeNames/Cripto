package cipherfactory

import (
	"crypto/des"

	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher"
	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher/block"
	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher/meta"
	"github.com/NooFreeNames/Cripto/pkg/crypto/key"
)

// DESFactory represents a factory for creating DES ciphers.
type DESFactory struct {
	meta.MetaProvider
	keyGen key.IKeyGen
}

// Cipher creates an DES cipher using the provided password and salt.
func (desf DESFactory) Cipher(password string, salt []byte) (cipher.ICipher, error) {
	desBlock, err := des.NewCipher(desf.keyGen.Generate(password, des.BlockSize))
	if err != nil {
		return block.BlockCipher{}, err
	}
	return block.NewBlockCipher(desBlock, salt, desf.Meta())
}

// NewDESFactory creates a new instance of DESFactory with the given
// keyGen and metadata.
func NewDESFactory(keyGen key.IKeyGen, m meta.IMeta) DESFactory {
	return DESFactory{meta.NewMetaProvider(m), keyGen}
}
