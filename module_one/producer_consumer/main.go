package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	cha := make(chan string, 5)
	defer close(cha)
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(cha, &wg)
	go consumer(cha, &wg)
	wg.Wait()
}

func producer(cha chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		fmt.Printf("produce: %s\n", s)
		cha <- s
	}
	wg.Done()
}

func consumer(cha <-chan string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		s := <-cha
		fmt.Printf("consume: %s\n", s)
	}
	wg.Done()
}
