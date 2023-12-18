package cipherfactory

import (
	"crypto/rc4"

	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher"
	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher/stream"
	"github.com/NooFreeNames/Cripto/pkg/crypto/key"
)

// RC4Factory represents a factory for creating RC4 ciphers.
type RC4Factory struct {
	keySize byte
	keyGen  key.IKeyGen
}

// Cipher creates an RC4 cipher using the provided password and salt.
func (c RC4Factory) Cipher(password string, salt []byte) (cipher.ICipher, error) {
	rc4Stream, err := rc4.NewCipher(c.keyGen.Generate(password, int(c.keySize)))
	if err != nil {
		return nil, err
	}
	return stream.NewStreamCipher(rc4Stream, salt), nil
}

// NewRC4Factory creates a new instance of RC4Factory with the given keyGen.
func NewRC4Factory(keySize byte, keyGen key.IKeyGen) RC4Factory {
	return RC4Factory{
		keySize: keySize,
		keyGen:  keyGen,
	}
}
