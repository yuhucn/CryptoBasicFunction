package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

/*
需求：生成并保存私钥、公钥

生成私钥分析：
1. GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥。
func GenerateKey(random io.Reader, bits int) (priv *PrivateKey, err error)
- 参数1： 随机数
- 参数2： 秘钥长度
- 返回值： 私钥

2. 要对生成的私钥进行编码处理，X509，按照规则，进行序列化处理，生成der编码的数据
MarshalPKCS1PublicKey将公钥序列化为PKCS格式DER编码。
func MarshalPKCS1PublicKey(pub interface{}) ([]byte, error)

3. 创建Block代表PEM编码的结构，并填入der编码的数据
type Block struct {
	Type string					// 得自前言的类型（如“RSA PRIVATE KEY"）
	Headers map[string]string	// 可选的头项
	Bytes []byte				// 内容解码后的数据，一般是DER编码额ASN.1结构
}

4. 将Pem Block数据写入到磁盘文件
func Encode(out io.Writer, b *Block) error

 */
func generateKeyPair(bits int) error {
	// 1. 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	// 2. 生成der编码的数据
	priDerText := x509.MarshalPKCS1PrivateKey(privateKey)

	// 3. 创建Block
	block := pem.Block{
		Type:"ZZ RSA PRIVATE KEY",
		Headers: nil,
		Bytes: priDerText,
	}

	// 4. 将Pem.Block数据写入到磁盘文件
	filehander1, err := os.Create(PrivateKeyFile)
	if err != nil {
		return err
	}

	// 关闭句柄
	defer filehander1.Close()

	err = pem.Encode(filehander1, &block)
	if err != nil {
		return err
	}

	fmt.Println("+++++++++++++生成公钥+++++++++++")
	/*
	1. 获取公钥，通过私钥函数
	2. x509
	3. 创建Block结构
	4. 将PEM写入磁盘
	 */
	pubkey := privateKey.PublicKey // 注意是对象，不是地址

	pubKeyDerText := x509.MarshalPKCS1PublicKey(&pubkey)

	block1 := pem.Block{
		Type:"ZZ RSA Public Key",
		Headers: nil,
		Bytes:   pubKeyDerText,
	}
	filehander2, err := os.Create(PublicKeyFile)
	if err != nil {
		return err
	}
	// 关闭句柄
	defer filehander2.Close()
	pem.Encode(filehander2,&block1)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Printf("generate rsa private key ...\n")
	err := generateKeyPair(1024)
	if err != nil {
		fmt.Printf("generate rsa private failed, err : %v", err)
	}
	fmt.Printf("generate rsa private successfully!\n")

}