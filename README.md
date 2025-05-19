# gocry

Encrypting files, strings, with "crypto / aes" and "crypto / cipher" algorithms.
It's cool that you will be able to use Crypt for your files.


# Example of operation

![gocry demo](img/gocry.gif)

# Packages

go get github.com/fatih/color

# Install

$ go build gocry.go

$ sudo mv gocry /usr/bin

# Example 1

```go

$ gocry --version

$ v.1.0

$ gocry --token 32

$ Your token [32] is:  TYJDUBTQWOLBEF4YONHD4ZDEL74F3IUW

$ gocry --crypt --file ~/Downloads/crypt/ex3.png --key TYJDUBTQWOLBEF4YONHD4ZDEL74F3IUW

$ gocry --descr --file ~/Downloads/crypt/ex3.png.crypt --key TYJDUBTQWOLBEF4YONHD4ZDEL74F3IUW

```

# Example 2

```go
	
$ gocry --help

  --crypt string
    	empty

  --descr string
    	empty

  --file string
    	Exs: file.pdf

  --key string
    	default 32 bytes: [DKYPENJXW43SMOJCU6F5TMFVOUANMJNL]

  --token string
    	Generate token: 16|32|64|128|512|1024 bytes (default "32")

```
