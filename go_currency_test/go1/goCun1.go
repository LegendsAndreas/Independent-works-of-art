package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// We measure how long it took to complete the loop with the 'start' variable, in minutes.
	start := time.Now()
	for i := 0; i < 1_000_000; i++ {
		fmt.Println(i)
	}

	// We save the time passed in the 'timePassed' variable as seconds.
	timePassed := strconv.FormatInt(int64(time.Since(start).Seconds()), 10)
	fmt.Println(time.Since(start))

	// We open the text file 'goCun1Records.txt', in the append mode, so that we can add other values to get, rather than just create a new value.
	file, err := os.OpenFile("goCun1Records.txt", os.O_APPEND|os.O_WRONLY, 0644)
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
