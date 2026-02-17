package main

import (
	"fmt"
	"sync"
)

func main() {

wg := &sync.WaitGroup{} //. waitgrup 
mut := &sync.Mutex{} // normal mutex
RWmut := &sync.RWMutex{} // Read and Write Mutex 
 
wg.Add(3) //. wait grp input -> how gorotines other wise wg.(add1) inside the function

var footballers []string

	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		mut.Lock() //. normla lock 
		fmt.Println("first")
		footballers = append(footballers, "messi")
		mut.Unlock()
		wg.Done() // passing fnished
	}(wg, mut)

	go func(wg *sync.WaitGroup,  mut *sync.RWMutex){
		fmt.Println("secound")
		mut.Lock()
		footballers = append(footballers, "ronaldo")
		mut.Unlock()
		wg.Done()// passing fnished
	}(wg, RWmut)

	go func(wg *sync.WaitGroup,  mut *sync.Mutex){
		RWmut.RLock() //. readlocking its like reading allow only all execution ends
				fmt.Println(footballers)
		RWmut.RUnlock()
		wg.Done()
	}(wg, mut)
wg.Wait() //. wait for all go 

fmt.Println(footballers)

}