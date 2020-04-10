package main

import (
	"fmt"
	"time"
)

func say1() {
	for i := 0; i < 1000001; i++ {
		fmt.Println(i)
	}
}

func main() {
	go say1()
	saya()
}

func saya() {
	time.Sleep(1 * time.Second)
	fmt.Println("aaa")
}