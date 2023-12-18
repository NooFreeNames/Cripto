// Package stream provides a simple implementation of a stream cipher.
package stream

import (
	"crypto/cipher"
	"fmt"
)

// StreamCipher allows you to encrypt and decrypt data using cipher.Stream
type StreamCipher struct {
	salt   []byte
	stream cipher.Stream
}

func (sc StreamCipher) Encrypt(dst, src []byte) (err error) {
	defer func() {
		if msg := recover(); msg != nil {
			err = fmt.Errorf("%v", msg)
		}
	}()
	sc.stream.XORKeyStream(dst, src)
	return nil
}

func (sc StreamCipher) Decrypt(dst, src []byte) (err error) {
	defer func() {
		if msg := recover(); msg != nil {
			err = fmt.Errorf("%v", msg)
		}
	}()
	sc.stream.XORKeyStream(dst, src)
	return nil
}

func (sc StreamCipher) BlockSize() int {
	return 1
}

func (sc StreamCipher) Salt() []byte {
	return sc.salt
}

// NewStreamCipher creates a new instance of StreamCipher with the given
// stream and salt.
func NewStreamCipher(stream cipher.Stream, salt []byte) StreamCipher {
	return StreamCipher{salt, stream}
}
