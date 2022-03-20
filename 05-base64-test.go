package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Printf("标准base64编码...\n")
	info := []byte("国足宇宙第一！")
	encodeInfo := base64.StdEncoding.EncodeToString(info)
	fmt.Printf("encode info 1: %s\n", encodeInfo)
	fmt.Printf("URL base64编码...\n")

	urlEncodeInfo := base64.URLEncoding.EncodeToString(info)
	fmt.Printf("url encode info 2 : %s\n", urlEncodeInfo)

}
