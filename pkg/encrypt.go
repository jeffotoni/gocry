/***********
*
*
* project crypt files, strings
*
* @package     main
* @author      jeffotoni
* @copyright   Copyright (c) 2017
* @license     --
* @link        --
* @since       Version 0.1
*
 */

package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"io"
)

//
//
// See alternate IV creation from ciphertext below
// var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
func Encrypt(key, text []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return ciphertext, nil
}

//
// Only multiples of 16, ex: 16, 32, 64, 128 .. max 1024
//
func GetToken(length int) string {

	//
	// 16 / 32 / 64 / 128 / 256 / 512 / 1024
	//
	if length < 16 || length > 1024 {

		return "Only multiples of 16, ex: 16, 32, 64, 128 .. max 1024"
	}

	lengthF := float32(length)

	//
	//
	//
	lengthPerm := lengthF / 16

	//
	//
	//
	if lengthPerm == 1 || lengthPerm == 2 || lengthPerm == 4 || lengthPerm == 8 || lengthPerm == 16 || lengthPerm == 32 || lengthPerm == 64 {

		randomBytes := make([]byte, 256)
		_, err := rand.Read(randomBytes)
		if err != nil {
			panic(err)
		}

		return base32.StdEncoding.EncodeToString(randomBytes)[:length]

	} else {

		return "Only multiples of 16, ex: 16, 32, 64, 128 .. max 1024"
	}
}
