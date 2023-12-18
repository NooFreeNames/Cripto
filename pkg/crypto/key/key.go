// Package key provides a simple key generation implementation using a
// hash function.
package key

import (
	"hash"
)

// IKeyGen represents an interface for key generation.
type IKeyGen interface {
	// Generate generates a key of the specified size from the given password.
	Generate(pasword string, keySize int) []byte
}

// KeyGen represents a key generator that uses a hash function to generate keys.
type KeyGen struct {
	hasher hash.Hash
}

// hash calculates the hash of the given data using the underlying
// hash function.
func (kg KeyGen) hash(data []byte) []byte {
	kg.hasher.Reset()
	kg.hasher.Write(data)
	return kg.hasher.Sum(nil)
}

// Generate generates a key of the specified size from the given password using
// the hash function.
func (kg KeyGen) Generate(pasword string, keySize int) []byte {
	hash := kg.hash([]byte(pasword))
	hashSize := len(hash)

	var keyRepeat int
	if keySize%hashSize == 0 {
		keyRepeat = keySize / hashSize
	} else {
		keyRepeat = keySize/hashSize + 1
	}

	buffer := make([]byte, 0, keyRepeat*keySize)
	for i := 0; i < keyRepeat; i++ {
		buffer = append(buffer, hash...)
	}
	return buffer[:keySize]
}

// NewKeyGen creates a new instance of KeyGen with the given hash function.
func NewKeyGen(hash hash.Hash) KeyGen {
	return KeyGen{hasher: hash}
}
