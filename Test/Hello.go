package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 11
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 22
	}()
	time.Sleep(3 * time.Second)

	x := <-ch
	fmt.Println(x)




	close(ch)
	k := <-ch
	fmt.Println(k)
	x = <-ch
	fmt.Println(x)
}