package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
)

type BookCLI struct {
	service *service.BookService
}

func NewBookCLI(service *service.BookService) *BookCLI {
	return &BookCLI{service: service}
}

// package for create CLIs complex (Viper)
func (cli *BookCLI) Run() {
	if len(os.Args) < 2 {
		fmt.Print("Usage: books <command> [arguments]")
	}

	command := os.Args[1]

	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <book title>")
			return
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)
	}
}

func (cli *BookCLI) searchBooks (name string) {
	books, err := cli.service.SearchBooksByName(name)

	if err != nil {
		fmt.Println("Error searching books: ", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found.")
	}

	fmt.Printf("%d books found\n", len(books))

	for _, book := range books {
		fmt.Printf("ID: %d, TITLE: %d, AUTHOR: %d, GENRE: %d\n", book.ID, book.Title, book.Author, book.Genre)
	}
}