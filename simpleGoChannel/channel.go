package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	array1 := [10]int{1, 9, 4, 7, 2, 5, 8, 0, 1, 4}
	array2 := [10]int{2, 6, 6, 5, 1, 0, 7, 3, 7, 9}
	array3 := [10]int{3, 8, 2, 8, 1, 2, 2, 0, 9, 8}
	array4 := [10]int{4, 1, 0, 6, 2, 7, 4, 2, 8, 2}

	ch := make(chan int)

	wg.Add(4)
	go goFunc(ch, array1)
	go goFunc(ch, array2)
	go goFunc(ch, array3)
	go goFunc(ch, array4)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

}

func goFunc(ch chan int, arr [10]int) {
	for _, num := range arr {
		ch <- num
	}
	wg.Done()
}

/*
Your program runs into a deadlock because it's trying to write to the channel from multiple goroutines but no other goroutines reading
from the channel at the same time. Sending to a channel will block until another goroutine receives from the channel. If there's no other
goroutine to read from the channel when sending, goroutines will be blocked, causing a deadlock.
*/

/*
By adding a separate goroutine that waits for all send operations to complete and then closes the channel, we can then safely iterate over
the channel contents using range in the main goroutine. Additionally, removed unnecessary mutex lock and unlock, as they were not being used
correctly and are unnecessary with the usage of channels.
*/
