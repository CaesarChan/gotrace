package main

import (
	"fmt"
	"time"
)

func main() {
	ch, ch1 := make(chan int), make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(10 * time.Millisecond)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			time.Sleep(20 * time.Millisecond)
		}
	}()
	/*
		for i := 0; i < 10; i++ {
			v := <-ch
			fmt.Println("Recv from ch", v)
			v1 := <-ch1
			fmt.Println("Recv from ch1", v1)
		}
	*/
	for i := 0; i < 20; i++ {
		select {
		case v := <-ch:
			fmt.Println("Recv from ch", v)
		case v := <-ch1:
			fmt.Println("Recv from ch1", v)
		}
	}
}