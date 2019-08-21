package main

import "fmt"

func demo() (int,int) {
	a:=1
	b:=2
	return a,b
}

func main() {
	d,e := demo()
	fmt.Println(d,e)
}