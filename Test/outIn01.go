package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan int, 100)
	ch1 := make(chan int, 10)
	go read01(ch, ch1)
	go write01(ch)
	<-ch1
	fmt.Println("第三步执行")
}

func read01(ch chan int, ch1 chan int) {
	for x := range ch {
		if len(ch) == 0 {
			ch1 <- 2
		}
		fmt.Println("----第二步执行,  取出x的值为：" + strconv.Itoa(x))
	}
}

func write01(ch chan int) {
	for i := 1; i <= 100; i++ {
		fmt.Println("第一步执行,写入的值为：" + strconv.Itoa(i))
		ch <- i
	}
}
