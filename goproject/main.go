package main

import "fmt"

func main() {

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// test := 8

	// for test <= 10 {
	// 	fmt.Println("Valor ", test)
	// 	test++
	// }

	test := []string{"test1", "teste2", "test3"}

	for i, value := range test {
		fmt.Println(value, i)
	}

	for _, value := range test {
		fmt.Println("Ignorando " + value)
	}

}
