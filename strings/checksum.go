package strings

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

func FileMD5(path string) (string, error) {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
