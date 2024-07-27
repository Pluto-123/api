package utils

import (
	"fmt"
	"testing"
)

func TestBcryptHash(t *testing.T) {
	hash, _ := BcryptHash("123456")
	fmt.Println(hash)
	return
}

func TestBcryptCheck(t *testing.T) {
	b := BcryptCheck("123456", "$2a$10$ze.d0x2I8BuqrxMqoTPy8enhspVsHaMwivu/.kyoCZPSYp9/BKEvK")
	fmt.Println(b)
}
