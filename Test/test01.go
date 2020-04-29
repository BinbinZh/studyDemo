package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 3
	go func() {
		for x := range ch {
			fmt.Println(x)
		}
	}()
	//go func() {
	//	x,ok:= <-ch
	//	fmt.Println(x,ok)
	//}()
	time.Sleep(1 * time.Second)
	fmt.Println("456")

}
