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
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

//
//
//
func genDecrypt(key, text []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}

//
//
//
func Decrypt(keyUser string, file string) {

	//
	//
	//
	if Exists(file) != true {

		fmt.Println("Error, File does not exist!")
		os.Exit(0)
	}

	//
	//
	//
	ValidateKey(keyUser)

	keyByte := []byte(keyUser)

	file_cry, _ := os.Open(file) // For read access.

	///pegando o tamanho em bytes do file..
	ficry, _ := file_cry.Stat()

	data_cry := make([]byte, ficry.Size())

	count_cry, _ := file_cry.Read(data_cry)

	file_copy_cry, _ := os.Create(file + ".descr")

	defer file_copy_cry.Close()

	data_descry, _ := genDecrypt(keyByte, data_cry[:count_cry])

	file_copy_cry.Write(data_descry)
}
