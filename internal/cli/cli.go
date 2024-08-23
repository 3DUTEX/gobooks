package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"time"
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
		return
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
		
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simualate <book_id> <book_id> <book_id> ...")
			return
		}

		booksIDs := os.Args[2:]

		cli.simulateReading(booksIDs)
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
		fmt.Printf("ID: %d, Título: %s, Autor: %s, Gênero: %s\n",
		book.ID, book.Title, book.Author, book.Genre)
	}
}

func (cli *BookCLI) simulateReading(bookIDsStr []string) {
	var bookIDs []int
	for _, idStr := range bookIDsStr {
			id, err := strconv.Atoi(idStr)

			if err != nil {
				continue;
			}

			bookIDs = append(bookIDs, id)
	}

	responses := cli.service.SimulateMultipleReadings(bookIDs, 2*time.Second)

	for _, response := range responses {
		fmt.Println(response)
	}
}