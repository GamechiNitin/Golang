package main

import (
	"fmt"
	"time"
)

func main() {
	go greater("Hello")
	greater("World")
	// Concurrency - Switch task/work.
	// Parallelism - Do task/work at same time.
	// Go routines - Managed by Go runtime. Flexible stack - 2kb
	// Thread - Managed by OS. Fixed stack 1mb

}

func greater(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(s, i)
	}
}
