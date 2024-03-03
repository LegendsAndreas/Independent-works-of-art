package main

import (
	"fmt"
	"log"
)

type converted struct {
	name  string
	value float64
}

// Actually sorting a map sounded fucking retarded, so i made it into an array with structs and then sorted and printed it lmao.
func main() {
	var name string
	var num float64

	var valueMap map[string]float64
	valueMap = make(map[string]float64)

	for {
		fmt.Printf("Enter name and value please> ")
		_, err := fmt.Scan(&name, &num)

		if name == "x" {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		valueMap[name] = num

		for key, val := range valueMap {
			fmt.Println(key, val)
		}

	}

	valueSlice := make([]converted, 0, len(valueMap))
	valueSlice = mapToSlice(valueMap, valueSlice)

	println("Unsorted")
	for i := 0; i < len(valueSlice); i++ {
		fmt.Printf("Key: %s, Value: %f\n", valueSlice[i].name, valueSlice[i].value)
	}

	valueSlice = insertionSort(valueSlice)

	println("Sorted")
	for i := 0; i < len(valueSlice); i++ {
		fmt.Printf("Key: %s, Value: %f\n", valueSlice[i].name, valueSlice[i].value)
	}
}

func insertionSort(x []converted) []converted {
	for i := 1; i < len(x); i++ {
		var curr float64 = x[i].value
		var j int = i - 1

		for j > -1 && curr < x[j].value {
			x[j+1] = x[j]
			j--
		}
		x[j+1].value = curr
	}
	return x
}

func mapToSlice(m map[string]float64, s []converted) []converted {
	var tempt converted
	for key, value := range m {
		tempt.name = key
		tempt.value = value
		s = append(s, tempt)
	}
	return s
}
