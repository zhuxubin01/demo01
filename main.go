package main

import (
	"os"
	"fmt"
)

func cal()  {



}

func main(){

	var i int

	var j int

	for _,val := range os.Args {
		fmt.Println(val)
	}
	fmt.Println("hello world")

	for i =1;i<10;i++ {
		fmt.Println(i)
	}
	for j =1;j<10;j++ {
		fmt.Println(i)

	}
}
