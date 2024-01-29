package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	"github.com/a-g-sanchez/Snipper_Snippets_API/config"
)

func Decrypt(ciphertext string) ([]byte, error) {
	data, err := base64.URLEncoding.DecodeString(ciphertext)

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(config.LoadKey())
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)

	return data, nil
}
