// Parameters for the GOST R 34.10-2001 CryptoPro Elliptic curves
package gost2001

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var initonce sync.Once
var gost2001a *elliptic.CurveParams
var gost2001b *elliptic.CurveParams
var gost2001c *elliptic.CurveParams

func initGOST2001A() {
	gost2001a = new(elliptic.CurveParams)
	gost2001a.P, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd97", 16)
	gost2001a.N, _ = new(big.Int).SetString("ffffffffffffffffffffffffffffffff6c611070995ad10045841b09b761b893", 16)
	gost2001a.B, _ = new(big.Int).SetString("a6", 16)
	gost2001a.Gx, _ = new(big.Int).SetString("01", 16)
	gost2001a.Gy, _ = new(big.Int).SetString("8d91e471e0989cda27df505a453f2b7635294f2ddf23e3b122acc99c9e9f1e14", 16)
	gost2001a.BitSize = 256
}

func GOST2001A() elliptic.Curve {
	initonce.Do(initGOST2001A)
	return gost2001a
}

func initGOST2001B() {
	gost2001b = new(elliptic.CurveParams)
	gost2001b.P, _ = new(big.Int).SetString("8000000000000000000000000000000000000000000000000000000000000c99", 16)
	gost2001b.N, _ = new(big.Int).SetString("800000000000000000000000000000015f700cfff1a624e5e497161bcc8a198f", 16)
	gost2001b.B, _ = new(big.Int).SetString("3e1af419a269a5f866a7d3c25c3df80ae979259373ff2b182f49d4ce7e1bbc8b", 16)
	gost2001b.Gx, _ = new(big.Int).SetString("01", 16)
	gost2001b.Gy, _ = new(big.Int).SetString("3fa8124359f96680b83d1c3eb2c070e5c545c9858d03ecfb744bf8d717717efc", 16)
	gost2001b.BitSize = 256
}

func GOST2001B() elliptic.Curve {
	initonce.Do(initGOST2001B)
	return gost2001b
}

func initGOST2001C() {
	gost2001c = new(elliptic.CurveParams)
	gost2001c.P, _ = new(big.Int).SetString("9b9f605f5a858107ab1ec85e6b41c8aacf846e86789051d37998f7b9022d759b", 16)
	gost2001c.N, _ = new(big.Int).SetString("9b9f605f5a858107ab1ec85e6b41c8aa582ca3511eddfb74f02f3a6598980bb9", 16)
	gost2001c.B, _ = new(big.Int).SetString("805a", 16)
	gost2001c.Gx, _ = new(big.Int).SetString("00", 16)
	gost2001c.Gy, _ = new(big.Int).SetString("41ece55743711a8c3cbf3783cd08c0ee4d4dc440d4641a8f366e550dfdb3bb67", 16)
	gost2001c.BitSize = 256
}

func GOST2001C() elliptic.Curve {
	initonce.Do(initGOST2001C)
	return gost2001c
}
