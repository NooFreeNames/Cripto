// Package block provides a block cipher implementation.
package block

import (
	"crypto/cipher"
	"fmt"
)

// BlockCipher represents a block cipher that can be used for encryption
// and decryption.
type BlockCipher struct {
	salt        []byte
	encryptMode cipher.BlockMode
	decryptMode cipher.BlockMode
}

func (bc BlockCipher) Salt() []byte {
	return bc.salt
}

func (bc BlockCipher) Encrypt(dst, src []byte) (err error) {
	defer func() {
		if msg := recover(); msg != nil {
			err = fmt.Errorf("%v", msg)
		}
	}()
	bc.encryptMode.CryptBlocks(dst, src)
	return nil
}

func (bc BlockCipher) Decrypt(dst, src []byte) (err error) {
	defer func() {
		if msg := recover(); msg != nil {
			err = fmt.Errorf("%v", msg)
		}
	}()
	bc.decryptMode.CryptBlocks(dst, src) //panic
	return nil
}

func (bc BlockCipher) BlockSize() int {
	return bc.encryptMode.BlockSize()
}

// saltToIV creates an IV from the given salt and IV size.
func saltToIV(salt []byte, IVsize int) []byte {
	iv := make([]byte, IVsize)
	copy(iv, salt)
	return iv
}

// NewBlockCipher creates a new BlockCipher using the provided block and salt.
func NewBlockCipher(block cipher.Block, salt []byte) (BlockCipher, error) {
	iv := saltToIV(salt, block.BlockSize())
	return BlockCipher{
		salt,
		cipher.NewCBCEncrypter(block, iv),
		cipher.NewCBCDecrypter(block, iv),
	}, nil
}
