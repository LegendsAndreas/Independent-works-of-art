package main

import "fmt"

type element struct {
	num  int
	num2 int
}

func main() {
	sli := make([]element, 0)
	sli = append(sli, element{num: 1, num2: 2})
	sli = append(sli, element{num: 10, num2: 11})
	sli = append(sli, element{num: 6, num2: 0})
	sli = append(sli, element{num: 2, num2: 5})
	sli = append(sli, element{num: 8, num2: 8})
	sli = append(sli, element{num: 1, num2: 2})
	sli = append(sli, element{num: 0, num2: 4})
	sli = append(sli, element{num: 4, num2: 8})
	sli = append(sli, element{num: 8, num2: 5})
	sli = append(sli, element{num: 7, num2: 12})

	fmt.Println(sli)
}
