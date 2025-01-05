package models

import (
	"database/sql"
	"fmt"
	"go-web/db"
	"time"
)

// Book model: A representation for books in the database
type Book struct {
	ID                int64      `json:"id"`
	Name              string     `json:"name" binding:"required"`
	Author            string     `json:"author" binding:"required"`
	Description       string     `json:"description" binding:"required"`
	ImageName         string     `json:"image_name" binding:"required"`
	ISBN              string     `json:"isbn" binding:"required"`
	Borrowed          bool       `json:"borrowed"`                      // If the book is currently borrowed
	LastBorrowedAt    *time.Time `json:"last_borrowed_at,omitempty"`    // Timestamp of the most recent borrow
	LastReturnedAt    *time.Time `json:"last_returned_at,omitempty"`    // Timestamp of the most recent return
	CurrentBorrowerID *int64     `json:"current_borrower_id,omitempty"` // ID of the current borrower, if any
}

// BorrowRecord model: Tracks books borrowed by users
type BorrowRecord struct {
	ID         int64      `json:"id"`
	BookID     int64      `json:"book_id"`
	UserID     int64      `json:"user_id"`
	BorrowedAt time.Time  `json:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at"` // NULL if not yet returned
}

// GetAllBooks : Retrieve all books from the database, including borrow information
func GetAllBooks() ([]Book, error) {
	var books []Book

	query := `
		SELECT 
    b.id, 
    b.name, 
    b.author, 
    b.description, 
    b.isbn,
    b.image_name,
    CASE 
        WHEN EXISTS (
            SELECT 1
            FROM borrow_records br2
            WHERE br2.book_id = b.id AND br2.returned_at IS NULL
            ORDER BY br2.borrowed_at DESC
            LIMIT 1
        ) THEN TRUE -- Currently borrowed
        ELSE FALSE -- Not borrowed
    END AS borrowed,
    DATETIME(MAX(br.borrowed_at)) AS last_borrowed_at,
    DATETIME(MAX(br.returned_at)) AS last_returned_at,
    (SELECT br2.user_id 
        FROM borrow_records br2 
        WHERE br2.book_id = b.id AND br2.returned_at IS NULL
        ORDER BY br2.borrowed_at DESC
        LIMIT 1) AS current_borrower_id
	FROM books b
	LEFT JOIN borrow_records br ON b.id = br.book_id
	GROUP BY b.id, b.name, b.author, b.description, b.isbn, b.image_name
	ORDER BY b.id ASC;`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("database query error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		var lastBorrowedAtString, lastReturnedAtString sql.NullString
		var currentBorrowerID *int64 // Nullable for currently not borrowed

		err := rows.Scan(
			&book.ID,
			&book.Name,
			&book.Author,
			&book.Description,
			&book.ISBN,
			&book.ImageName,
			&book.Borrowed,
			&lastBorrowedAtString,
			&lastReturnedAtString,
			&currentBorrowerID,
		)
		if err != nil {
			return nil, fmt.Errorf("row scanning error: %v", err)
		}

		// Convert strings to *time.Time
		if lastBorrowedAtString.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", lastBorrowedAtString.String)
			if err != nil {
				return nil, fmt.Errorf("error parsing last_borrowed_at: %v", err)
			}
			book.LastBorrowedAt = &parsedTime
		} else {
			book.LastBorrowedAt = nil
		}

		if lastReturnedAtString.Valid {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", lastReturnedAtString.String)
			if err != nil {
				return nil, fmt.Errorf("error parsing last_returned_at: %v", err)
			}
			book.LastReturnedAt = &parsedTime
		} else {
			book.LastReturnedAt = nil
		}

		book.CurrentBorrowerID = currentBorrowerID

		books = append(books, book)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", rows.Err())
	}

	return books, nil
}

// GetBookByID : Retrieve a book by its ID, including borrow information
func GetBookByID(id int64) (*Book, error) {
	query := `
		SELECT 
    b.id, 
    b.name, 
    b.author, 
    b.description,
    b.isbn,
    b.image_name,
    CASE 
        WHEN EXISTS (
            SELECT 1
            FROM borrow_records br2
            WHERE br2.book_id = b.id AND br2.returned_at IS NULL
            ORDER BY br2.borrowed_at DESC
            LIMIT 1
        ) THEN TRUE -- Currently borrowed
        ELSE FALSE -- Not borrowed
    END AS borrowed,
    DATETIME(MAX(br.borrowed_at)) AS last_borrowed_at,
    DATETIME(MAX(br.returned_at)) AS last_returned_at,
    (SELECT br2.user_id 
        FROM borrow_records br2
        WHERE br2.book_id = b.id AND br2.returned_at IS NULL
        ORDER BY br2.borrowed_at DESC
        LIMIT 1) AS current_borrower_id
	FROM books b
	LEFT JOIN borrow_records br ON b.id = br.book_id
	WHERE b.id = ?
	GROUP BY b.id, b.name, b.author, b.description, b.isbn`

	row := db.DB.QueryRow(query, id)

	var book Book
	var lastBorrowedAtString, lastReturnedAtString sql.NullString
	var currentBorrowerID *int64

	// Scan the result row
	err := row.Scan(
		&book.ID,
		&book.Name,
		&book.Author,
		&book.Description,
		&book.ISBN,
		&book.ImageName,
		&book.Borrowed,
		&lastBorrowedAtString,
		&lastReturnedAtString,
		&currentBorrowerID,
	)
	if err != nil {
		return nil, fmt.Errorf("row scanning error: %v", err)
	}

	// Convert strings to *time.Time
	if lastBorrowedAtString.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", lastBorrowedAtString.String)
		if err != nil {
			return nil, fmt.Errorf("error parsing last_borrowed_at: %v", err)
		}
		book.LastBorrowedAt = &parsedTime
	} else {
		book.LastBorrowedAt = nil
	}

	if lastReturnedAtString.Valid {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", lastReturnedAtString.String)
		if err != nil {
			return nil, fmt.Errorf("error parsing last_returned_at: %v", err)
		}
		book.LastReturnedAt = &parsedTime
	} else {
		book.LastReturnedAt = nil
	}

	book.CurrentBorrowerID = currentBorrowerID

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
