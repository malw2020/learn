package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Printf("only once\n")
	}

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		go func() {
			fmt.Printf("call once count: %v\n", i)
			once.Do(onceBody) // 多次调用只执行一次
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("wait to exit, count:%v\n", i)
		<-done
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("end\n")
}
