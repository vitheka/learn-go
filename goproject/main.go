package main

/*
antes da v1.18

o tipo interface{} é o no "anny", ele é o que chamam de interface vazia

func userTest(user interface{}) {

	switch v:= user.(type) {

    case UserTest1:

	case UserTest2:

	default:
		panic("Invalid type!")

	}

}

---------------comparar tipos any----------------

func anyTest[T comparable](arg1 T, arg2 T) {
	fmt.Println(arg1 == arg2)
}

func main() {
	anyTest(1, 1)
}

------------------------------------------------
Token ~ em go ou "Elemento aproximação"

type AnyString interface {
	~string
}

type ApproximateMyString interface{
	~MyString
}

*/

func main() {

}
