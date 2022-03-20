package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)
const privateKeyFile = "./RsaPrivateKey.pem"
const publicKeyFile  = "./RsaPublicKey.pem"

func rsaPubEncrypt(filename string, plainText []byte) (error, []byte) {
	// 1. 通过公钥文件，读取公钥信息 -> pem encode的数据
	info, err := ioutil.ReadFile(filename)
	if err != nil {
		return err, nil
	}
	// 2. pem decode，得到block中的der编码数据
	block, _ := pem.Decode(info)
	// 返回值1：pem.block
	// 返回值2：rest参加是未解码完的数据，存储在这里
	// 3. 解码der，得到公钥
	derText := block.Bytes
	publickey, err := x509.ParsePKCS1PublicKey(derText)
	if err != nil {
		return err, nil
	}
	// 4.共钥加密
	//EncryptPKCS1v15使用PKCS#1 v1.5规定的填充方案和RSA算法加密msg。
	// func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) (out []byte, err error)
	cipherData, err := rsa.EncryptPKCS1v15(rand.Reader, publickey, plainText)
	if err != nil {
		return err, nil
	}
	return nil, cipherData
}

func rsaPriKeyDecrypt(filename string, cipherData []byte) (error, []byte) {
	// 1. 通过私钥文件，读取私钥信息 -> pem encode的数据
	info, err := ioutil.ReadFile(filename)
	if err != nil {
		return err, nil
	}
	// 2. pem decode，得到block中的der编码数据
	block, _ := pem.Decode(info)
	// 返回值1：pem.block
	// 返回值2：rest参加是未解码完的数据，存储在这里
	// 3. 解码der，得到私钥
	derText := block.Bytes
	privatekey, err := x509.ParsePKCS1PrivateKey(derText)
	if err != nil {
		return err, nil
	}
	// 4.私钥解密
	//DecryptPKCS1v15使用PKCS#1 v1.5规定的填充方案和RSA算法解密密文。
	//func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) (out []byte, err error)
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privatekey, cipherData)
	if err != nil {
		return err, nil
	}
	return nil, plainText
}

func main() {
	src := []byte("祝党的生日快乐！")
	err, cipherData := rsaPubEncrypt(publicKeyFile, src)
	if err != nil {
		fmt.Println("公钥加密失败！")
	}
	fmt.Printf("cipherData : %x\n",cipherData)
	fmt.Println("++++++++++++++++++++++++")

	err, plainText := rsaPriKeyDecrypt(privateKeyFile, cipherData)
	if err != nil {
		fmt.Println("私钥解密失败！ err: ", err)
	}
	fmt.Printf("plainText : %s\n", plainText)
}