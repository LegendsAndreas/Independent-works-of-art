package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	var input string
	fmt.Printf("Enter text file> ")
	fmt.Scan(&input)

	fileData, err := os.Open(input + ".txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	structSli := make([]person, 0, 3)
	structSli = addStructs(structSli, fileData)

	for _, n := range structSli {
		fmt.Printf("First Name: %s, Last Name: %s\n", n.fname, n.lname)
	}
}

func addStructs(arr []person, files *os.File) []person {
	scans := bufio.NewScanner(files)
	for scans.Scan() {
		line := scans.Text()

		sections := strings.Fields(line)
		if len(sections) >= 2 {
			arr = append(arr, person{fname: sections[0], lname: sections[1]})
		}
	}
	files.Close()

	return arr
}
