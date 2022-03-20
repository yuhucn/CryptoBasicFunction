package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// 使用打开文件方式获取哈希
const filename = "C:/Users/w3sch/downloads/Ethereum-Wallet-installer-0-11-1.exe"
func main(){
	// 1. open file
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	// 2. 创建hash
	hasher := sha256.New()
	/*
	type Hash interface {
		io.Writer
		Sum(b []byte) []byte
		Reset()
		Size() int
		BlockSize() int
	}
	 */
	// 3. copy句柄
	// func Copy(dst Writer, src Reader) (written int64, err error)
	length, err:= io.Copy(hasher, file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("length: %d\n", length)
	// 4. hash sum操作
	hash := hasher.Sum(nil)

	fmt.Printf("hash: %x\n", hash)
}
