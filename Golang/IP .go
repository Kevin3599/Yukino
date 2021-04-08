package main

import (
	"fmt"
    "github.com/tabalt/ipquery"
)
    func main(){
	  df :="testdata/test_10000.data"
     err :=ipquery.load
    if err!=nil {
		fmt.Println(err)
	}
    ip:=""
    dt,err=ipquery.find(ip)
    if err !=nil{
	             fmt.println(error)
	}else {
	             fmt.println（ip，string(dt)）
	} 
}

