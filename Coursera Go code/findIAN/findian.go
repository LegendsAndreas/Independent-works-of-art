package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Assigns a user entered string to x, and then converts it to lower case, to make sure that we only work with lower case letters.
	fmt.Println("Enter a string> ")
	x := bufio.NewScanner(os.Stdin)
	var line string
	if x.Scan() {
		line = x.Text()
	}
	/*n, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println(n, err)
	}*/
	str := strings.ToLower(line)

	//Checks if "i", "n" and "a" are at their desired places.
	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}

}
