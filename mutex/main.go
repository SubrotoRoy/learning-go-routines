package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	counter := 0
	gs := 100
	var m sync.Mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
			m.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Final Counter :", counter)
}
