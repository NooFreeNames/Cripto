// Package block provides a block cipher implementation.
package block

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

// BlockCipher represents a block cipher that can be used for encryption
// and decryption.
type BlockCipher struct {
	block cipher.Block
}

// pkcs7Pad pads the given data using PKCS#7 padding.
func (a BlockCipher) pkcs7Pad(data []byte) []byte {
	blockSize := a.block.BlockSize()
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// pkcs7Unpad removes PKCS7 padding from the given byte slice.
// It returns an error if the padding is invalid.
func (a BlockCipher) pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	unpadding := int(data[length-1])
	if unpadding > a.block.BlockSize() || unpadding == 0 {
		return nil, errors.New("invalid padding")
	}
	return data[:(length - unpadding)], nil
}

// generateIV generates a random initialization vector (IV) for encryption.
// It returns the generated IV, along with an error if generation fails.
func (a BlockCipher) generateIV() ([]byte, error) {
	iv := make([]byte, a.block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	return iv, nil
}

// isCipherDataValid checks if the given ciphertext is valid for decryption.
// It returns true if the length of the ciphertext is a multiple of
// the block size and at least two blocks long.
func (bc BlockCipher) isCipherDataValid(cipherData []byte) bool {
	l := len(cipherData)
	blockSize := bc.block.BlockSize()
	return l%blockSize == 0 && l >= blockSize*2
}

// Encrypt encrypts the given plainData using CBC mode with PKCS#7 padding
// and returns the encrypted data along with the IV used.
func (bc BlockCipher) Encrypt(plainData []byte) ([]byte, error) {
	iv, err := bc.generateIV()
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(bc.block, iv)
	plainData = bc.pkcs7Pad(plainData)

	cipherData := make([]byte, len(plainData))
	mode.CryptBlocks(cipherData, plainData)

	return append(cipherData, iv...), nil
}

// Decrypt decrypts the given cipherData using CBC mode with PKCS#7 unpadding
// and returns the decrypted plain data.
func (bc BlockCipher) Decrypt(cipherData []byte) ([]byte, error) {
	if !bc.isCipherDataValid(cipherData) {
		return cipherData, errors.New("invalid cipher data format")
	}

	ivStartIndex := len(cipherData) - bc.block.BlockSize()
	iv := cipherData[ivStartIndex:]
	cipherData = cipherData[:ivStartIndex]

	mode := cipher.NewCBCDecrypter(bc.block, iv)

	plainData := make([]byte, len(cipherData))
	mode.CryptBlocks(plainData, cipherData)

	plainData, err := bc.pkcs7Unpad(plainData)
	if err != nil {
		return cipherData, err
	}

	return plainData, nil
}
