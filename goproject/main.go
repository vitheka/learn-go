package main

import (
	//"errors"
	"fmt"
)

//funcao anonima

func main() {

	test()

}

func test() {

	func(valorString string, valorInt int) {
		fmt.Println(valorString, valorInt)
	}("otavio", 20)
}

//funcao

// func main() {

// 	value, err := test()

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(value)

// }

// func test() (string, error) {
// 	return "", errors.New("test")
// }
