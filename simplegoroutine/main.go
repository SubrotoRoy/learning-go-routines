package main

import (
	"fmt"
	"runtime"
)

func main() { 
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GoRoutines before foo", runtime.NumGoroutine())
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("GOROOT", runtime.GOROOT())
	go foo()
	fmt.Println("GoRoutines after foo", runtime.NumGoroutine())
	runtime.Gosched()
}

func foo() {
	fmt.Println("hello")
}
