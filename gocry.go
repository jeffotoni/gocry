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

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	cry "github.com/jeffotoni/gocry/pkg"
)

const (
	keyDefault = "DKYPENJXW43SMOJCU6F5TMFVOUANMJNL"
)

func main() {

	//
	//
	//
	white := color.New(color.FgWhite)
	boldWhite := white.Add(color.Bold)

	//
	//
	//
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	//
	//
	//
	yellow := color.New(color.FgYellow)
	boldYellow := yellow.Add(color.Bold)

	//
	//
	//
	flag.String("crypt", "", "empty")

	//
	//
	//
	flag.String("descr", "", "empty")

	//
	//
	//
	flag.String("key", "", "default 32 bytes: ["+keyDefault+"]")

	//
	//
	//
	flag.String("token", "32", "Generate token: 16|32|64|128|512|1024 bytes")

	//
	//
	//
	flag.String("file", "", "Exs: file.pdf")

	//
	//
	//
	sizeArgs := len(os.Args)

	//fmt.Println(sizeArgs)
	//
	// Validate flags
	// and Validate hidden flags
	// --v
	// -v
	// --version
	// -version
	//
	// -help
	// --help
	// -h
	// --h
	//
	if sizeArgs <= 1 {

		flag.PrintDefaults()
		os.Exit(0)
	}

	//
	//
	//
	var stringCmd string

	//
	//
	//
	var stringCmd2 string

	//
	//
	//
	var file string

	file = ""
	//
	//
	//
	var keyUser string

	//
	//
	//
	keyUser = keyDefault

	//
	//
	//
	var cmdIn int

	cmdIn = 0
	//
	//
	//

	//
	// Validate hidden flags
	//
	for x := range os.Args {

		stringCmd = strings.Trim(os.Args[x], "-")
		stringCmd = strings.Trim(stringCmd, "/")
		stringCmd = strings.Trim(stringCmd, ".")
		stringCmd = strings.Trim(stringCmd, "-")

		stringCmd = strings.TrimSpace(stringCmd)
		stringCmd = strings.ToLower(stringCmd)

		//fmt.Println("args: ", sizeArgs, " ", x)

		switch stringCmd {

		case "crypt":

			cmdIn += 1

			//crypt = fmt.Sprintf("%d", stringCmd2)

		case "descr":

			stringCmd2 = strings.Trim(os.Args[x+1], "-")
			stringCmd2 = strings.TrimSpace(stringCmd2)

			if stringCmd2 != "" {

				boldRed.Println("Error, command has no value!")
				os.Exit(0)
			}

			//descr = fmt.Sprintf("%s", stringCmd2)

			cmdIn += 3

		case "file":

			stringCmd2 = strings.Trim(os.Args[x+1], "-")
			stringCmd2 = strings.TrimSpace(stringCmd2)
			file = fmt.Sprintf("%s", stringCmd2)

			//
			// exist file
			//
			if cry.Exists(file) != true {

				boldRed.Println("Error, File does not exist!")
				os.Exit(0)
			}

			cmdIn += 1

		case "key":

			stringCmd2 = strings.Trim(os.Args[x+1], "-")
			stringCmd2 = strings.TrimSpace(stringCmd2)
			keyUser = fmt.Sprintf("%s", stringCmd2)

			//
			// validate
			//
			cry.ValidateKey(keyUser)

		case "token":

			stringCmd2 = strings.Trim(os.Args[x+1], "-")
			stringCmd2 = strings.TrimSpace(stringCmd2)

			if tok, err := strconv.Atoi(stringCmd2); err != nil {

				boldRed.Println("Error, Only multiple integers of 16!")
				os.Exit(0)

			} else {

				tokenKey := cry.GetToken(tok)
				boldRed.Println("Your token ["+stringCmd2+"] is: ", tokenKey)
				os.Exit(0)
			}

		case "version":

			boldYellow.Println("v.1.0")
			os.Exit(0)

		case "v":

			boldYellow.Println("v.1.0")
			os.Exit(0)

		case "help":

			flag.PrintDefaults()
			os.Exit(0)

		case "h":

			flag.PrintDefaults()
			os.Exit(0)

		default:
			//flag.PrintDefaults()
			//os.Exit(0)

		}
	}

	var keyByte = []byte("")

	//
	// crypt == 2
	//
	// decry == 4
	//
	// key default
	//
	if cmdIn == 2 || cmdIn == 4 {

		if keyUser == "" {

			keyUser = keyDefault
		}

		keyByte = []byte(keyUser) // 32 bytes | 64 bytes | 128 bytes ... max 1024 bytes

		// boldWhite.Println("Key used to encrypt: ", keyUser, " :: byte: ", keyByte)

		FileL, _ := os.Open(file)

		fi, _ := FileL.Stat()

		data := make([]byte, 16*fi.Size())

		count, _ := FileL.Read(data)

		file_copy, _ := os.Create(file + ".crypt")

		defer file_copy.Close()

		ciphertext, _ := cry.Encrypt(keyByte, data[:count])

		///gravando arquivo cryptografado
		file_copy.Write(ciphertext)

		boldWhite.Println("New file created and encrypted: [" + file + ".crypt" + "]")
		boldWhite.Println("Used Key: [" + keyUser + "]")

	} else {

		flag.PrintDefaults()
		os.Exit(0)
	}

}
