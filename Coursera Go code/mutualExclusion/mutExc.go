package main

import (
	"fmt"
	"sync"
)

var j int = 0
var ug sync.WaitGroup
var mut sync.Mutex

func main() {
	for {
		ug.Add(2)
		go ink()
		go ink()
		ug.Wait()
		fmt.Println(j)
		if j != 2 {
			// After waiting between 1 sec to 1 min, this usually executes.
			fmt.Println("It fucking happened!")
			break
		}
		j = 0
	}
}

func ink() {
	mut.Lock()
	j = j + 1
	ug.Done()
	mut.Unlock()
}
