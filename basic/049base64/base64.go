package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	url := "https://www.google.com"

	sEncode := base64.StdEncoding.EncodeToString([]byte(url))
	fmt.Println(sEncode)
	sDecode, _ := base64.StdEncoding.DecodeString(sEncode)
	fmt.Println(string(sDecode))

	uEncode := base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println(uEncode)
	uDecode, _ := base64.URLEncoding.DecodeString(uEncode)
	fmt.Println(string(uDecode))
}
