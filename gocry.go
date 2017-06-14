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
	flag.String("crypt", "32 bytes", "Ex: 16|32|64|128|256 bytes")

	//
	//
	//
	flag.String("descr", "key", "Ex: DKYPENJXW43SMOJCU6F5TMFVOUANMJNL")

	//
	//
	//
	flag.String("file", "", "Exs: file.pdf")

	//
	//
	//
	sizeArgs := len(os.Args)

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
	var crypt int

	//
	//
	//
	crypt = 32

	//
	//
	//
	var descr string

	//
	//
	//
	var file string

	//
	// Validate hidden flags
	//
	for x := range os.Args {

		stringCmd = strings.Trim(os.Args[x], "-")
		stringCmd = strings.TrimSpace(stringCmd)
		stringCmd = strings.ToLower(stringCmd)

		//fmt.Println("args: ", sizeArgs, " ", x)

		switch stringCmd {

		case "crypt":

			stringCmd2 = strings.Trim(os.Args[x+1], "-")
			stringCmd2 = strings.TrimSpace(stringCmd2)

			if xx, err := strconv.Atoi(stringCmd2); err != nil {

				boldRed.Println("Error, Only multiple integers of 16!")
				os.Exit(0)

			} else {

				crypt = xx
			}

			//crypt = fmt.Sprintf("%d", stringCmd2)

		case "descr":

			stringCmd2 = strings.Trim(os.Args[x+1], "-")
			stringCmd2 = strings.TrimSpace(stringCmd2)
			descr = fmt.Sprintf("%s", stringCmd2)

		case "file":

			stringCmd2 = strings.Trim(os.Args[x+1], "-")
			stringCmd2 = strings.TrimSpace(stringCmd2)
			file = fmt.Sprintf("%s", stringCmd2)

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

	tokenKey := cry.GetToken(crypt)
	fmt.Println(tokenKey)

	key := []byte(tokenKey) // 32 bytes

	boldWhite.Println("Key used to encrypt: ", tokenKey)

	fmt.Println("Key used to encrypt: ", key)
	fmt.Println("Key used to encrypt: ", descr)
	fmt.Println("Key used to encrypt: ", file)
	fmt.Println("Key used to encrypt: ", crypt)

}
