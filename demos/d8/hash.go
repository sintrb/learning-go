package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	str := "this is a test string~~~~~"
	md := md5.New()
	md.Write([]byte(str))
	rst := md.Sum([]byte(""))
	fmt.Printf("%x\n\n", rst)

	sh := sha1.New()
	sh.Write([]byte(str))
	rst = sh.Sum([]byte(""))
	fmt.Printf("%x\n\n", rst)
}
