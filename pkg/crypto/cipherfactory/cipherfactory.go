// Package cipherfactory provides functionality for storing cipher until a
// password is received from the user
package cipherfactory

import (
	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher"
)

// ICipherFactory represents an interface for creating ciphers.
// ICipherFactory can be used to store the cipher until the password is received
// from the user
type ICipherFactory interface {
	// Cipher creates a new instance of ICipher with the given password and salt.
	Cipher(password string, salt []byte) (cipher.ICipher, error)
}
