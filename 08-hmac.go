package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

/*
// 使用分析

// 接收端和验证端都要执行
// New函数返回一个采用hash.Hash作为底层hash接口、key作为密钥的HMAC算法的hash接口。
func New(h func() hash.Hash, key []byte) hash.Hash
- 参数1：自己指定哈希算法，是一个函数
	-md5.New
	-sha1.New
	-sha256.New
- 参数2：秘钥
- 参数3：哈希函数对象


// 仅在验证端执行
// New函数返回一个采用hash.Hash作为底层hash接口、key作为密钥的HMAC算法的hash接口。
func New(h func() hash.Hash, key []byte) hash.Hash
- 参数1：自己计算的哈希值
- 参数2：接收到的哈希值
- 返回值：对比结果
 */
func generateHMAC(src []byte, key []byte) []byte {
	// 1.创建哈希器
	hasher := hmac.New(sha256.New, key)

	// 2.生成mac值
	hasher.Write(src)
	mac := hasher.Sum(nil)
	return mac
}

// 认证mac
func verifyHMAC(src, key, mac1 []byte) bool {
	// 1. 对端接受到的源数据

	// 2。对端接受到的mac1

	// 3.对端计算本地的mac2
	mac2 := generateHMAC(src, key)
	// 4. 对比mac1与mac2
	return hmac.Equal(mac1, mac2)

}
func main() {
	src := []byte("hello world")
	key := []byte("1234567890")
	mac1 := generateHMAC(src, key)

	fmt.Printf("mac1: %x\n", mac1)

	isEqual := verifyHMAC(src, key, mac1)
	fmt.Printf("isEqual : %v\n", isEqual)

	srcChanged := []byte("hello world!!!")

	isEqual = verifyHMAC(srcChanged, key, mac1)
	fmt.Printf("After changed, isEqual : %v\n", isEqual)
}
