package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// 生成私钥公钥
func genetateEccKeypair() {
	// - 选择一个椭圆曲线（在elliptic包）
		// type Curve
		// func P224() Curve
		// func P256() Curve
		// func P384() Curve
		// func P521() Curve
	curve := elliptic.P256()

	// - 使用ecdsa包，创建私钥 // ecdsa椭圆曲线数字签名
	// GenerateKey函数生成密钥对
	// func GenerateKey(curve Curve, rand io.Reader) (priv []byte, x, y *big.Int, err error)
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	checkErr("generate keypair failed!", err)

	// - 使用x509进行解码
	// MarshalECPrivateKey将ecdsa私钥序列化为ASN.1 DER编码
	// func MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error)
	derText, err := x509.MarshalECPrivateKey(privateKey)
	checkErr("MarshalECPrivateKey failed!", err)

	// 写入pem.Block中
	block1 := pem.Block{
		Type:	"ECC PRIVATE KEY",
		Headers:	nil,
		Bytes:	derText,
	}

	// -pem.Encode
	fileHander, err := os.Create(EccPrivateKeyFile)
	checkErr("os.Create Failed!", err)
	defer fileHander.Close()
	err = pem.Encode(fileHander, &block1)
	checkErr("pem Encode failed!", err)

	fmt.Printf("++++++++++++++++++\n")

	// 获取公钥
	publicKey := privateKey.PublicKey

	// 使用x509进行编码
	// 通用的序列化方法
	derText2, err := x509.MarshalPKIXPublicKey(&publicKey)
	checkErr("MarshalPKIXPublicKey", err)

	// 写入pem.Block中
	block2 := pem.Block{
		Type:	"ECC PUBLIC KEY",
		Headers:	nil,
		Bytes:	derText2,
	}

	// -pem.Encode
	fileHander2, err := os.Create(EccPublicKeyFile)
	checkErr("public key os.Create Failed!", err)
	defer fileHander2.Close()
	err = pem.Encode(fileHander2, &block2)
	checkErr("public key pem Encode failed!", err)
}

func main(){
	genetateEccKeypair()
}
