package main

import (
	"fmt"
	"time"
)

func main() {
	s := New(1)

	go func() {
		for i := 0; i < 10; i++ {
			r, e := s.AcquireTimeout(1000)
			fmt.Printf("acqure timeout 01, index:%v, state: %v, %v\n", i, r, e)

			if r == false {
				continue
			}

			s.Release()
			fmt.Printf("relase timeout 01, index:%v, %v\n", i, e)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			r, e := s.AcquireTimeout(1000)
			fmt.Printf("acqure timeout 02, index:%v, state: %v, %v\n", i, r, e)

			if r == false {
				continue
			}

			s.Release()
			fmt.Printf("relase timeout 02, index:%v, %v\n", i, e)
		}
	}()

	time.Sleep(time.Second * time.Duration(10))

}
