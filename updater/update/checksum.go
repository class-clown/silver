/*
 * Copyright © 2021 PaperCut Software International Pty. Ltd.
 */

package update

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
)

func checksum(hashType string, file string) string {
	var hasher hash.Hash

	switch {
	case hashType == "sha256":
		hasher = sha256.New()
	case hashType == "sha1":
		hasher = sha1.New()
	default:
		hasher = sha1.New()
	}
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.Copy(hasher, f)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
