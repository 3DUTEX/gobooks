package main

import (
	"fmt"
	"gobooks/internal/service"
)

// não tem try catch
// retorno duplo
// func sum(x int, y int) (int, error) {
// 	if x + y == 10 {
// 		return x + y, nil
// 	}

// 	return 0, errors.New("Resultado deve ser igual 10!");
// }

// func main() {
// 	// ":=" cria e atribui valor na variável
// 	x, err := sum(5, 5)

// 	if err != nil { panic(err) }

// 	println(x)
// }

func main() {
	book := service.Book{
		ID: 1,
		Title: "Harry Potter",
		Author: "J.K Rowling",
		Genre: "Suspense",
	}

	fmt.Println(book.GetFullBook())
}