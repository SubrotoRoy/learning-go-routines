package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int
	count = 0
	gs := 100 //number of rouroutines we want

	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := count
			v++
			count = v
			wg.Done()
			fmt.Println(runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("count", count)
}

//use go run -race main.go command to detect race condition