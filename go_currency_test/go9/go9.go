package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	wg.Add(9)
	go func() { // Goroutine 1
		for i := 0; i < 111_111; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() { // Goroutine 2
		for i := 0; i < 111_111; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() { // Goroutine 3
		for i := 0; i < 111_111; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() { // Goroutine 4
		for i := 0; i < 111_111; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() { // Goroutine 5
		for i := 0; i < 111_111; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() { // Goroutine 6
		for i := 0; i < 111_111; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() { // Goroutine 7
		for i := 0; i < 111_111; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() { // Goroutine 8
		for i := 0; i < 111_111; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() { // Goroutine 9
		for i := 0; i < 111_111; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()
	wg.Wait()

	// We save the time passed in the 'timePassed' variable as seconds.
	timePassed := strconv.FormatInt(int64(time.Since(start).Seconds()), 10)
	fmt.Println(time.Since(start))

	// We open the text file 'goCun1Records.txt', in the append mode, so that we can add other values to get, rather than just create a new value.
	file, err := os.OpenFile("goCun9Records.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// We defer the file to close, so that it closes properly when the program ends.
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// We then write the actual string into the file. We also add a newline, so that it's more readable.
	_, err2 := file.WriteString(timePassed + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}
}
