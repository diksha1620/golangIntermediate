package main

import (
	"fmt"
	"time"
)

func main() {
	greeter("Diksha")
	go greeter("Hello")

}

func greeter(s string) {

	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}

}
