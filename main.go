package main

import "fmt"

func main() {
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "test")

	fmt.Println(names)
}
