package main

// "go mod tidy" import modules in project
import (
	"database/sql"
	"fmt"

	"gobooks/internal/service"
	"gobooks/internal/web"

	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	db := openCon()
	defer db.Close()

	bookService := service.NewBookService(db)

	bookHandlers := web.NewBookHandlers(bookService)

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("GET /books/{id}", bookHandlers.GetBookByID)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)

	PORT := "8080"
	fmt.Println("SERVER IS RUNNING ON: http://localhost:", PORT)
	http.ListenAndServe(":" + PORT, router)
}

func openCon() *sql.DB {
	db, err := sql.Open("sqlite", "./books.db")

	if err != nil {
		panic(err)
	}

	return db;
}