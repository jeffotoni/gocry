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
	"fmt"
	"io"
	"os"
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

func Crypt(keyUser string, file string) {

	if Exists(file) != true {

		fmt.Println("Error, File does not exist!")
		os.Exit(0)
	}

	//
	//
	//
	ValidateKey(keyUser)

	//
	// 32 bytes | 64 bytes | 128 bytes ... max 1024 bytes
	//
	keyByte := []byte(keyUser)

	//
	//
	//
	FileL, _ := os.Open(file)

	//
	//
	//
	fi, _ := FileL.Stat()

	//
	//
	//
	data := make([]byte, 16*fi.Size())

	//
	//
	//
	count, _ := FileL.Read(data)

	//
	//
	//
	file_copy, _ := os.Create(file + ".crypt")

	//
	//
	//
	defer file_copy.Close()

	//
	//
	//
	ciphertext, _ := Encrypt(keyByte, data[:count])

	///gravando arquivo cryptografado
	file_copy.Write(ciphertext)

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

		randomBytes := make([]byte, length)
		_, err := rand.Read(randomBytes)
		if err != nil {
			panic(err)
		}

		return base32.StdEncoding.EncodeToString(randomBytes)[:length]

	} else {

		return "Only multiples of 16, ex: 16, 32, 64, 128 .. max 1024"
	}
}

//
// Only multiples of 16, ex: 16, 32, 64, 128 .. max 1024
//
func ValidateKey(key string) {

	//
	// 16 / 32 / 64 / 128 / 256 / 512 / 1024
	//
	if len(key) < 16 || len(key) > 1024 {

		fmt.Println("Only multiples of 16, ex: 16, 32, 64, 128 .. max 1024 Ex: [DKYPENJXW43SMOJCU6F5TMFVOUANMJNL]")
		os.Exit(0)
	}

	lengthF := len(key)

	//
	//
	//
	if lengthF == 16 || lengthF == 32 || lengthF == 64 || lengthF == 128 || lengthF == 256 || lengthF == 512 || lengthF == 1024 {

		//
		// ok
		//
		return

	} else {

		fmt.Println("Only multiples of 16, ex: 16, 32, 64, 128 .. max 1024, Ex: [DKYPENJXW43SMOJCU6F5TMFVOUANMJNL]")
		os.Exit(0)
	}
}

//
//
//
func Exists(fileName string) bool {

	_, err := os.Stat(fileName)

	return !os.IsNotExist(err)
}
