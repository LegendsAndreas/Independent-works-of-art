package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var destination = "C:/Users/andre/Desktop/"
	var folderName string
	var fileName string
	var fileContent string

	/* Asks for folder name, file name and content to add to the file. We are using
	bufio.NewScanner(os.Stdin), to make names with spaces possible. */
	fmt.Println("Insert folder name, file name and file content(one word each)>")

	// Folder name.
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		folderName = input.Text()
	}

	// File name.
	input = bufio.NewScanner(os.Stdin)
	if input.Scan() {
		fileName = input.Text()
	}

	// File content.
	input = bufio.NewScanner(os.Stdin)
	if input.Scan() {
		fileContent = input.Text()
	}

	// Makes the new folder.
	err1 := os.MkdirAll(destination+folderName, 0777)
	if err1 != nil {
		log.Fatal(err1)
	}

	// Makes the new text file.
	_, err2 := os.Create(destination + folderName + "/" + fileName + ".txt")
	if err2 != nil {
		log.Fatal(err2)
	}

	// Adds the content into the text file.
	err3 := os.WriteFile(destination+folderName+"/"+fileName+".txt", []byte(fileContent), 0777)
	if err3 != nil {
		log.Fatal(err3)
	}
}
