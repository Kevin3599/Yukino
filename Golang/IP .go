package main

import (
	"errors"
	"fmt"

	"github.com/tabalt/ipquery"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)
func main(){
	def:"testdata/test_10000.data"
     err=ipquery.load
    if err!=nil{
		fmt.Println(err)
	}
ip:=""
dt,err=ipquery.find(ip)
if err !=nil{
	
}
