package main

import (
	"fmt"
	"sync"
)

var i int = 0
var wg sync.WaitGroup

func main() {
	for {
		wg.Add(2)
		go inc()
		go inc()
		wg.Wait()
		fmt.Println(i)
		if i != 2 {
			// After waiting between 1 sec to 1 min, this usually executes.
			fmt.Println("It fucking happened!")
			break
		}
		i = 0
	}
}

func inc() {
	i = i + 1
	wg.Done()
}
