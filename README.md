# multibase
Utility for multi base encoding and decoding

Bored of using multiples tools for base encoding?
Multibase is a tool written in Go that allows you to encode/decode strings and files in differents base


Supported encoding and decoding methods:

- [x] base32
- [x] base58
- [x] base64
- [x] URL-compatible base64
- [x] PostScript/PDF base85

## Usage

```
Usage:
  mb  [OPTIONS]

Application Options:
  -f, --file=   Specify input file
      --b32     Generate base32 of given string/file
      --b58     Generate base58 of given string/file
      --b64     Generate base64 of given string/file
      --b64u    Generate URL-compatible base64
      --b85     Generate Abobe's PostScript/PDF base85 of given string/file
  -d, --decode  Decode data

Help Options:
  -h, --help    Show this help message
```
## Installation

- Solving dependencies
```
$ go get -v "github.com/jbenet/go-base58"
$ go get -v "github.com/jessevdk/go-flags"
```
- Cloning and building

```
$ git clone https://github.com/Nhoya/multibase && cd multibase
$ go build
```

The package is also available on [AUR](https://aur.archlinux.org/packages/multibase/)

- Installing

`# mv multibase /usr/local/bin/mb`


