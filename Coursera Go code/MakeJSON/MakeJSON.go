package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, addr string
	var name2, addr2 string
	var idMap map[string]string
	idMap = make(map[string]string)

	//Asks for name and address
	fmt.Printf("Enter name> ")
	fmt.Scan(&name)

	fmt.Printf("Enter address> ")
	fmt.Scan(&addr)

	fmt.Printf("Enter name> ")
	fmt.Scan(&name2)

	fmt.Printf("Enter address> ")
	fmt.Scan(&addr2)

	//Assigns the value name to "Name" and value addr to "Address"
	idMap["Name"] = name
	idMap["Address"] = addr
	idMap["Two"] = name2
	idMap["Three"] = addr2

	//Marshals the map
	barr, err := json.Marshal(idMap)
	if err != nil {
		fmt.Println(err)
	}

	//Prints the byte array as a string, so it is readable to humans
	fmt.Print(string(barr))
}
