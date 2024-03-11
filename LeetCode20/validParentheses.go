package main

import "fmt"

func main() {
	var input string = "()[]{}"
	boolShit := isValid(input)
	fmt.Println(boolShit)
}

func isValid(str string) bool {
	boolShit := false

	/* We know that, since you need a full pair of either "()", "{}" or "[]", the string cannot be true
	if the length of the string is not an even number. As such, we just return the boolShit variable right away.
	*/
	if (len(str) % 2) != 0 {
		return boolShit
	}

	return boolShit
}
