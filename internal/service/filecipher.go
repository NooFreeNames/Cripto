package service

import (
	"crypto/md5"

	"github.com/NooFreeNames/Cripto/pkg/crypto/cipher/block"
	"github.com/NooFreeNames/Cripto/pkg/crypto/filecipher"
)

// FileCipherService implements encryption and decryption operations for files.
type FileCipherService struct {
}

// Encrypt encrypts the file using the provided input parameters.
func (cs *FileCipherService) Encrypt(input FileCipherInput) error {
	return initFileCipher(input.Path, input.Password).Encrypt()
}

// Decrypt decrypts the file using the provided input parameters.
func (cs *FileCipherService) Decrypt(input FileCipherInput) error {
	return initFileCipher(input.Path, input.Password).Decrypt()
}

// initFileCipher initializes the file cipher.
func initFileCipher(path, password string) filecipher.FileCipherI {
	hash := md5.Sum([]byte(password))
	aes, _ := block.NewAESCipher(hash[:])
	return filecipher.NewFileCipher(path, aes)
}

// NewFileCipherService creates a new instance of FileCipherService.
func NewFileCipherService() *FileCipherService {
	return &FileCipherService{}
}