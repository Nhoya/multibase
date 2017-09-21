package main

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
	b58 "github.com/jbenet/go-base58"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"os"
)

var opts struct {
	Base32    bool `long:"b32" description:"Generate base32 of given string/file"`
	Base58    bool `long:"b58" description:"Generate base58 of given string/file"`
	Base64    bool `long:"b64" description:"Generate base64 of given string/file"`
	Base64URL bool `long:"b64u" description:"Generate URL- compatible base64"`
	Decode    bool `short:"d" long:"decode" description:"Decode data"`
}

func main() {
	result := ""
	var byteResult []byte
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	target, _ := ioutil.ReadAll(os.Stdin)
	targetString := string(target[:])

	if target == nil {
		fmt.Println("You must specify input")
		os.Exit(1)
	}

	if opts.Base32 {
		if opts.Decode {
			byteResult, _ = base32.StdEncoding.DecodeString(targetString)
			result = string(byteResult[:])
		} else {
			result = base32.StdEncoding.EncodeToString(target)
		}
	} else if opts.Base58 {
		if opts.Decode {
			result = string(b58.DecodeAlphabet(targetString, b58.BTCAlphabet)[:])
		} else {
			result = b58.EncodeAlphabet(target, b58.BTCAlphabet)
		}
	} else if opts.Base64 {
		if opts.Decode {
			byteResult, _ = base64.StdEncoding.DecodeString(targetString)
			result = string(byteResult[:])
		} else {
			result = base64.StdEncoding.EncodeToString(target)
		}
	} else if opts.Base64URL {
		if opts.Decode {
			byteResult, _ = base64.URLEncoding.DecodeString(targetString)
			result = string(byteResult[:])
		} else {
			result = base64.URLEncoding.EncodeToString(target)
		}
	}

	if result != "" {
		fmt.Println(result)
	}
}
