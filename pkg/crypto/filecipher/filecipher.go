// Package filecipher provides functionality for encrypting and decrypting
// files using a specified cipher algorithm.
package filecipher

import (
	"os"

	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher"
)

// processFunc type represents a function that takes a byte slice as input,
// processes it, and returns a modified byte slice along with an error.
type processFunc func([]byte) ([]byte, error)

// FileCipherI defining methods for file encryption and decryption.
type FileCipherI interface {
	Encrypt() error
	Decrypt() error
}

// FileCipher represents a file cipher with a specific path and cipher
// algorithm.
type FileCipher struct {
	// path to the file to be encrypted or decrypted.
	path string
	// cipher algorithm to use for encryption or decryption.
	cipher cipher.Cipher
}

// processFileData reads the file data, applies the provided process function
// to it, and then writes the modified data back to the file.
// It returns an error if any IO or processing error occurs.
func (fc FileCipher) processFileData(processFunc processFunc) error {
	data, err := os.ReadFile(fc.path)
	if err != nil {
		return err
	}
	new_data, err := processFunc(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(fc.path, new_data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Encrypt encrypts the file data using the cipher algorithm specified in
// FileCipher. It returns an error if the encryption process fails.
func (fc FileCipher) Encrypt() error {
	return fc.processFileData(fc.cipher.Encrypt)
}

// Decrypt decrypts the file data using the cipher algorithm specified
// in FileCipher. It returns an error if the decryption process fails.
func (fc FileCipher) Decrypt() error {
	return fc.processFileData(fc.cipher.Decrypt)
}

// NewFileCipher creates a new instance of FileCipher with the given file
// path and cipher algorithm.
func NewFileCipher(path string, cipher cipher.Cipher) *FileCipher {
	return &FileCipher{path: path, cipher: cipher}
}
