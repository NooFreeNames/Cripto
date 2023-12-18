// Package cipher provides the ability to encrypt and decrypt data.
package cipher

type ICipher interface {
	// Encrypt encrypts a number of blocks. The length of
	// src must be a multiple of the block size. Dst and src must overlap
	// entirely or not at all.
	//
	// If len(dst) < len(src), Encrypt return error. It is acceptable
	// to pass a dst bigger than src, and in that case, Encrypt will
	// only update dst[:len(src)] and will not touch the rest of dst.
	//
	// Multiple calls to Encrypt behave as if the concatenation of
	// the src buffers was passed in a single run. That is, ICipher
	// maintains state and does not reset at each Encrypt call.
	Encrypt(dst, src []byte) error
	// Decrypt decrypts a number of blocks. The length of
	// src must be a multiple of the block size. Dst and src must overlap
	// entirely or not at all.
	//
	// If len(dst) < len(src), Decrypt return error. It is acceptable
	// to pass a dst bigger than src, and in that case, Decrypt will
	// only update dst[:len(src)] and will not touch the rest of dst.
	//
	// Multiple calls to Decrypt behave as if the concatenation of
	// the src buffers was passed in a single run. That is, ICipher
	// maintains state and does not reset at each Decrypt call.
	Decrypt(dst, src []byte) error
	// BlockSize returns the block size of the cipher.
	// If the cipher is stream then 1 will be returned.
	BlockSize() int
	// Salt returns the salt associated with the cipher.
	Salt() []byte
}
