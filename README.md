# GOST2001
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/gost2001/blob/master/LICENSE.md) 
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/gost2001/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/gost2001/releases)
[![GoDoc](https://godoc.org/github.com/pedroalbanese/gost2001?status.png)](http://godoc.org/github.com/pedroalbanese/gost2001)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/gost2001)](https://goreportcard.com/report/github.com/pedroalbanese/gost2001)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/gost2001)](https://github.com/pedroalbanese/gost2001/releases)

### GOST R 34.10-2001 CryptoPro ParamSets (January 2006)
Package implements the elliptic curves originally described in RFC 4357

## Usage
```
Usage of gost2001:
  -derive
        Derive shared secret.
  -key string
        Private/Public key.
  -keygen
        Generate keypair.
  -pub string
        Remote's side Public key.
  -sign
        Sign with Private key.
  -signature string
        Signature.
  -verify
        Verify with Public key.
```
## Examples
#### Asymmetric keypair generation:
```sh
./gost2001 -keygen 
```
#### Digital signature (ECDSA):
```sh
./gost2001 -sign -key $prvkey < file.ext > sign.txt
sign=$(cat sign.txt)
./gost2001 -verify -key $pubkey -signature $sign < file.ext
```
#### Shared key agreement (ECDH a.k.a. VKO):
```sh
./gost2001 -derive -key $prvkey -pub $pubkey
```

## License

This project is licensed under the ISC License.

##### Industrial-Grade Reliability. Copyright (c) 2020-2022 ALBANESE Research Lab.
