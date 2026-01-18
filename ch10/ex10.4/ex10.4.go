package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 23.1, 11.1}

	for i, v := range t {
		fmt.Println(i, v)
	}
}
