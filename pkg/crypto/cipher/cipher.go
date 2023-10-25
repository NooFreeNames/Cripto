// Package cipher provides the ability to encrypt and decrypt data.
package cipher

// Cipher represents a cipher that can be used to encrypt and decrypt data
type Cipher interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}
