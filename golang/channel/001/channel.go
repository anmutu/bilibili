/*
  author='du'
  date='2021/1/12 2:04'
  github='https://github.com/anmutu'
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	//写
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	//读
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
