package utils

import (
	"time"
	"crypto/md5"
	"crypto/sha256"
	"log"
)

// TODO functions

func EncryptResponse(content *string) *string {
	return content
}

func DecryptRequest(*content *string) *string {
	return content
}

func EncryptToken(token *string) *string {
	encryptor := sha256.New()
	encryptor.Write([]byte(token))

	encryptedData := encryptor.Sum(data)
	return &hex.EncodeToString(encryptedData[:])
}

func CreateToken(basString *string, encryptionTime *time.Time) *string {
	data := []byte(basString + encryptionTime.String())
	hashedData := md5.Sum(data)
	return &hex.EncodeToString(hashedData[:])
}