package models

import (
	"go-web/db"
	"time"
)

// Book model: A representation for books in the database
type Book struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
}

// BorrowRecord model: Tracks books borrowed by users
type BorrowRecord struct {
	ID         int64      `json:"id"`
	BookID     int64      `json:"book_id"`
	UserID     int64      `json:"user_id"`
	BorrowedAt time.Time  `json:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at"` // NULL if not yet returned
}

// GetAllBooks : Retrieve all books from the database
func GetAllBooks() ([]Book, error) {
	var books []Book

	query := `SELECT * FROM books`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Description, &book.ISBN)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// GetBookByID : Retrieve a book by its ID
func GetBookByID(id int64) (*Book, error) {
	query := `SELECT * FROM books WHERE id=?`
	row := db.DB.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Name, &book.Author, &book.Description, &book.ISBN)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// BorrowBook : Borrow a book by creating a borrow record
func BorrowBook(bookID, userID int64) error {
	// Check if the book is already borrowed
	query := `SELECT COUNT(*) FROM borrow_records WHERE book_id = ? AND returned_at IS NULL`
	var count int
	err := db.DB.QueryRow(query, bookID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return &BookError{"Book is currently borrowed by another user"}
	}

	query = `INSERT INTO borrow_records (book_id, user_id, borrowed_at) VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(bookID, userID, time.Now())
	return err
}

// ReturnBook : Return a previously borrowed book
func ReturnBook(bookID, userID int64) error {
	// Check if the book is borrowed by the user and not yet returned
	query := `SELECT COUNT(*) FROM borrow_records WHERE book_id = ? AND user_id = ? AND returned_at IS NULL`
	var count int
	err := db.DB.QueryRow(query, bookID, userID).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return &BookError{"No active borrow record found for this book and user"}
	}

	// Mark the book as returned
	query = `UPDATE borrow_records SET returned_at = ? WHERE book_id = ? AND user_id = ? AND returned_at IS NULL`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), bookID, userID)
	return err
}

// BookError Custom error type for book operations
type BookError struct {
	Message string
}

func (e *BookError) Error() string {
	return e.Message
}
