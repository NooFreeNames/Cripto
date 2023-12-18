package cipherfactory

import (
	"crypto/rc4"

	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher"
	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher/meta"
	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher/stream"
	"github.com/NooFreeNames/Cripto/pkg/crypto/key"
)

// RC4Factory represents a factory for creating RC4 ciphers.
type RC4Factory struct {
	meta.MetaProvider
	keySize byte
	keyGen  key.IKeyGen
}

// Cipher creates an RC4 cipher using the provided password and salt.
func (rc4f RC4Factory) Cipher(password string, salt []byte) (cipher.ICipher, error) {
	rc4Stream, err := rc4.NewCipher(rc4f.keyGen.Generate(password, int(rc4f.keySize)))
	if err != nil {
		return nil, err
	}
	return stream.NewStreamCipher(rc4Stream, salt, rc4f.Meta()), nil
}

// NewRC4Factory creates a new instance of RC4Factory with the given
// keySize, keyGen and metadata.
func NewRC4Factory(keySize byte, keyGen key.IKeyGen, m meta.IMeta) RC4Factory {
	return RC4Factory{
		MetaProvider: meta.NewMetaProvider(m),
		keySize:      keySize,
		keyGen:       keyGen,
	}
}
