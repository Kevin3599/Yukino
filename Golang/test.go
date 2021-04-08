package main

import (
	"fmt"

	"github.com/tabalt/ipquery"
)

func main() {
	varl1 := "testdata/test_10000.data"
	err := ipquery.load
	if err != nil {
		fmt.Println(err)
	}
	ip := ""
	varl1, err = ipquery.find(ip)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ip, string(varl1))
	}
}

}
