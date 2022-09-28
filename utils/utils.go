package utils

import (
	"bytes"
	"crypto/sha1"
	"math/rand"
)

func EncryptPassword(password string, salt [8]byte) [20]byte {
	p1 := sha1.Sum([]byte(password))
	buf := new(bytes.Buffer)
	buf.Write(salt[:])
	buf.Write(p1[:])
	return sha1.Sum(buf.Bytes())
}

func GenSalt() [8]byte {
	salt := make([]byte, 8, 8)
	rand.Read(salt)
	return *(*[8]byte)(salt)
}
