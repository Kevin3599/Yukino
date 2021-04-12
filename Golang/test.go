package main

import (
	"fmt"

	"github.com/goinbox/ipquery"
)

func main() {
	fmt.Println("欢迎使用")
	varl1 := "testdata/test_10000.data"
	err := ipquery.Load
	if err != nil {
		fmt.Println(err)
	}
	ip := "varl2"
	varl1, err = ipquery.find(ip)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ip, string(varl1))
	}
}
