package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second*3)
		fmt.Println("go")
	}()

	fmt.Println("gooovinada")
	time.Sleep(time.Second*5)
}
