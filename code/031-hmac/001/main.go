package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("test")
	fmt.Println(c)
	c = getCode("test")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("secret"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
