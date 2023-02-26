package main

import "fmt"

func main() {

	array := [5]int{4, 2, 5, 6, 7}

	// realizar la funcionalidad

	fmt.Println("Los valores del array son: ", array)

	// For que gusrda el idice y el valot
	for index, value := range array {
		array[index] = value + 20
	}

	fmt.Println("Los valores del array son: ", array)
}
