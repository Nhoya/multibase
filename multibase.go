package main

import (
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	b58 "github.com/jbenet/go-base58"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	File      string `short:"f" long:"file" description:"Specify input file"`
	Base32    bool   `long:"b32" description:"Generate base32 of given string/file"`
	Base58    bool   `long:"b58" description:"Generate base58 of given string/file"`
	Base64    bool   `long:"b64" description:"Generate base64 of given string/file"`
	Base64URL bool   `long:"b64u" description:"Generate URL-compatible base64"`
	Base85    bool   `long:"b85" description:"Generate Abobe's PostScript/PDF base85 of given string/file"`
	Decode    bool   `short:"d" long:"decode" description:"Decode data"`
}

type Base struct {
	Target string
	Result string
	Decode bool
}

func main() {

	out := Base{}

	//parsing flags
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	parser := flags.NewParser(&opts, flags.Default)

	//if the file flag is set use it as target, otherwhise get the string
	if opts.File != "" {
		buf, err := ioutil.ReadFile(opts.File)
		if err != nil {
			fmt.Println(err)
		}
		out.Target = string(buf)
		fmt.Println(out.Target)
	} else {
		buf, _ := ioutil.ReadAll(os.Stdin)
		out.Target = string(buf)
	}
	out.Decode = opts.Decode

	//this should be prettier
	if opts.Base32 {
		out.b32()
	} else if opts.Base58 {
		out.b58()
	} else if opts.Base64 {
		out.b64()
	} else if opts.Base64URL {
		out.b64u()
	} else if opts.Base85 {
		out.b85()
	} else {
		parser.WriteHelp(os.Stderr)
	}

	if out.Result != "" {
		fmt.Println(out.Result)
	}
}

func (b *Base) b32() {
	if b.Decode {
		byteResult, _ := base32.StdEncoding.DecodeString(b.Target)
		b.Result = string(byteResult[:])
	} else {
		b.Result = base32.StdEncoding.EncodeToString([]byte(b.Target))
	}
}

func (b *Base) b58() {
	if b.Decode {
		b.Result = string(b58.DecodeAlphabet(b.Target, b58.BTCAlphabet)[:])
	} else {
		b.Result = b58.EncodeAlphabet([]byte(b.Target), b58.BTCAlphabet)
	}
}

func (b *Base) b64() {
	if b.Decode {
		byteResult, _ := base64.StdEncoding.DecodeString(b.Target)
		b.Result = string(byteResult[:])
	} else {
		b.Result = base64.StdEncoding.EncodeToString([]byte(b.Target))
	}
}

func (b *Base) b64u() {
	if b.Decode {
		byteResult, _ := base64.URLEncoding.DecodeString(b.Target)
		b.Result = string(byteResult[:])
	} else {
		b.Result = base64.URLEncoding.EncodeToString([]byte(b.Target))
	}
}

func (b *Base) b85() {
	if b.Decode {
		buffer := make([]byte, len(b.Target))
		ascii85.Decode(buffer, []byte(b.Target), true)
		b.Result = string(buffer)
	} else {
		buffer := make([]byte, ascii85.MaxEncodedLen(len(b.Target)))
		ascii85.Encode(buffer, []byte(b.Target))
		b.Result = string(buffer)
	}
}
