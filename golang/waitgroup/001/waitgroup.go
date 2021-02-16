/*
  author='du'
  date='2021/1/12 2:10'
  github='https://github.com/anmutu'
*/
package main

import (
	"fmt"
	"sync"
)

const N = 10

var wg = &sync.WaitGroup{}

func main() {
	Test01()
	fmt.Println("+++++++++++")
	Test02()
}

func Test01() {
	for i := 0; i < N; i++ {
		go func(i int) {
			wg.Add(1)
			println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

func Test02() {
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
