package service

import "database/sql"

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