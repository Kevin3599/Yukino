package main

import (
	"fmt"

	"github.com/goinbox/ipquery"
)

func main() {
	fmt.Println("欢迎使用")
	var ip  float32
	fmt.Println("输入你的IP")
	varl2=fmt.Scanln(&ip)
	})
	varl1 := "testdata/test_10000.data"
	err := ipquery.Load
	if err != nil {
		fmt.Println(err)
	}
	ip := "varl1"
	varl1, err = ipquery.find(ip)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ip, string(varl1))
	}
}
