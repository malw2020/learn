package main

import (
	"fmt"
	"time"
)

func main() {
	dataChan := make(chan int, 5)

	fmt.Printf("timestamp:%v, 启动生产、消费协程...\n", time.Now().UnixNano())

	go producers(dataChan)
	go consumer(dataChan)

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
	}

}

func producers(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Printf("timestamp:%v, send data:%v\n", time.Now().UnixNano(), i)
	}
}

func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Printf("timestamp:%v, reviced data:%v\n", time.Now().UnixNano(), v)
	}
}
