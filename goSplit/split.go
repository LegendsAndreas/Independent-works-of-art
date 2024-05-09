package main

import "fmt"

func main() {
	var s = "This, right here, is my, swag!"
	var subS []string

	subS = stringSplit(subS, s, ' ')

	for i := 0; i < len(subS); i++ {
		fmt.Printf("%d = %s \n", i, subS[i])
	}
}

// stringSplit splits a string into multiple substrings using a specified delimiter.
// It takes in a slice of strings, a string to be split, and a delimiter rune.
// The function iterates through the characters of the string and creates substrings based on the delimiter.
// The substrings are stored in the input slice and returned.
// If the delimiter is found, the current substring is appended to the slice, and a new substring is started.
// If the input string starts with the delimiter, the first character is ignored.
// An empty string is not appended to the slice.
// The last substring is appended separately to ensure all substrings are included.
// If the input string is empty or contains only the delimiter, an empty slice is returned.
func stringSplit(subStrings []string, strings string, splitter rune) []string {
	var tempSubString string // The substrings that we are creating will be stored here, then later appended to subStrings.
	for i, value := range strings {
		if value == splitter && i != 0 { // We do not add the first character of the full string 'strings' if it is a space.
			subStrings = append(subStrings, tempSubString) // We append the temporary substring to the subStrings slice.
			tempSubString = ""                             // Then we reset.
		} else if value != splitter { // We also have to check that we do not add a space to the temporary substring.
			tempSubString += string(value)
		}
	}
	// We append tempSubString into subStrings one more time, to make sure we get the last substring with us.
	// We also check if 'tempSubString' is empty, since we don't want to append an empty string into the slice.
	if tempSubString != "" {
		subStrings = append(subStrings, tempSubString)
	}

	// Returns the entire slice full of substrings.
	return subStrings
}
