package main

import (
	. "asdf"
	"encoding/hex"
	"fmt"
)

var plaintext = []byte("0123456789abcdef")
var key = "0123456789abcdef0123456789ABCDEF"

func main() {
	crypt := CryptCreate(0, 0, []byte(key))

	ciphertext := make([]byte, len(plaintext))

	crypt.Cipher.Encrypt(ciphertext, plaintext)
	fmt.Printf("plaintext(%s)==>ciphertext(%v)\n",
		hex.EncodeToString(plaintext),
		hex.EncodeToString(ciphertext))

	crypt.Cipher.Decrypt(plaintext, ciphertext)
	fmt.Printf("ciphertext(%s)==>plaintext(%s)\n",
		hex.EncodeToString(ciphertext),
		hex.EncodeToString(plaintext))

	plaintext = []byte("00112233445566778899aabbccddeeff")
	crypt.Hash.Reset()
	crypt.Hash.Write(plaintext)
	hmac := crypt.Hash.Sum(nil)

	fmt.Printf("plaintext(%s)==>hmac(%s)\n",
		hex.EncodeToString(plaintext),
		hex.EncodeToString(hmac))
}
