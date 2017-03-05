package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

// Hasher calculates hash values for files.
type Hasher interface {
	Hash(path string) string
}

// SHA256 calculates the digest of file contents.
type SHA256 struct{}

// Hash returns SHA256 hash of a file.
func (sha SHA256) Hash(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hash := sha256.New()

	_, err = io.Copy(hash, f)
	if err != nil {
		log.Fatal(err)
	}

	result := hex.EncodeToString(hash.Sum(nil))

	return result
}
