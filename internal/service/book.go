package service

import (
	"database/sql"
	"fmt"
	"time"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db: db}
}

// método em GO
func (b Book) GetFullBook() string {
	return b.Title + " by " + b.Author
}

type BookService struct {
	db *sql.DB
}

func (s *BookService) CreateBook(book *Book) error {
	query := "INSERT INTO books (Title, Author, Genre) VALUES(?, ?, ?)"
	
	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)

	if err != nil { return err }

	lastInsertID, err := result.LastInsertId()

	if err != nil { return err }

	book.ID = int(lastInsertID);

	return nil
}

func (s *BookService) GetBooks() ([]Book, error) {
	query := "SELECT Id, Title, Author, Genre FROM books"

	rows, err := s.db.Query(query);

	if err != nil {
		return nil, err
	}

	var books []Book
	for rows.Next() {
		var book Book

		// "&" altera exatamente o objeto na memória
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) GetBookByID(id int) (*Book, error) {
	query := "SELECT Id, Title, Author, Genre FROM books WHERE Id = ?"
	row := s.db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)

	if err != nil {
		return nil, err
	}

		return &book, nil
 }

 func (s *BookService) UpdateBook(book *Book) error {
	query := "UPDATE books SET Title = ?, Author = ?, Genre = ? WHERE Id = ?"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)
	return err
 }

 func (s *BookService) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE Id = ?"
	_, err := s.db.Exec(query, id)
	return err
 }

 // SimulateReading simula a leitura de um livro com base em um tempo de leitura.
func (s *BookService) SimulateReading(bookID int, duration time.Duration, results chan<- string) {
	book, err := s.GetBookByID(bookID)
	if err != nil || book == nil {
		results <- fmt.Sprintf("Livro com ID %d não encontrado.", bookID)
		return
	}

	time.Sleep(duration) // Simula o tempo de leitura.
	results <- fmt.Sprintf("Leitura do livro '%s' concluída!", book.Title)
}

// SimulateMultipleReadings simula a leitura de múltiplos livros simultaneamente.
func (s *BookService) SimulateMultipleReadings(bookIDs []int, duration time.Duration) []string {
	results := make(chan string, len(bookIDs)) // Canal com buffer para evitar bloqueio

	// Lança as goroutines para simular a leitura.
	for _, id := range bookIDs {
		go func(bookID int) {
			s.SimulateReading(bookID, duration, results)
		}(id)
	}

	var responses []string
	for range bookIDs {
		responses = append(responses, <-results)
	}
	close(results) // Fechamento do canal após coleta de todos os resultados

	return responses
}