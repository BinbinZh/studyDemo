package main

import "fmt"

func main() {
	ch := make(chan int)
	ch1 := make(chan int)
	go read(ch, ch1)
	go write(ch)
	<-ch1
	fmt.Println("第三步执行")
}

func read(ch chan int, ch1 chan int) {
	<-ch
	fmt.Println("第二步执行")
	ch1 <- 2
}

func write(ch chan int) {
	fmt.Println("第一步执行")
	ch <- 1
}
