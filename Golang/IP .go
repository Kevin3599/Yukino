package main

import (
	"fmt"
)

func main() {
	df := "testdata/test_10000.data"
	err := ipquery.load
	if err != nil {
		fmt.Println(err)
	}
	ip := ""
	dt, err = ipquery.find(ip)
	if err != nil {
		fmt.Println(err，check your ip)
	} else {
		fmt.Println(ip, string(dt))
	}
}
