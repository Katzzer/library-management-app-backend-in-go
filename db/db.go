package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

	insertTestData()
}

func createTables() {
	// Users Table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY, -- SQLite provides autoincrement functionality when using INTEGER PRIMARY KEY
	    email TEXT NOT NULL UNIQUE,
	    password TEXT NOT NULL,
	    latest_jwt_token TEXT
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table.")
	}

	// Books Table
	createBooksTable := `
	CREATE TABLE IF NOT EXISTS books (
	    id INTEGER PRIMARY KEY, -- PRIMARY KEY with INTEGER ensures auto-increment behavior
	    name TEXT NOT NULL,
	    author TEXT NOT NULL,
	    description TEXT NOT NULL,
	    isbn TEXT NOT NULL UNIQUE,
	    image_name TEXT NOT NULL
	)
	`

	_, err = DB.Exec(createBooksTable)
	if err != nil {
		panic("Could not create books table.")
	}

	// Borrow Records Table
	createBorrowRecordsTable := `
	CREATE TABLE IF NOT EXISTS borrow_records (
	    id INTEGER PRIMARY KEY, -- PRIMARY KEY with INTEGER ensures auto-increment behavior
	    book_id INTEGER NOT NULL,
	    user_id INTEGER NOT NULL,
	    borrowed_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	    returned_at DATETIME,
	    FOREIGN KEY (book_id) REFERENCES books(id),
	    FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createBorrowRecordsTable)
	if err != nil {
		panic("Could not create borrow_records table.")
	}
}

// insertTestData: Inserts initial data into the books table
func insertTestData() {

	err := insertTestDataFromFile("sql/books.sql", DB)
	if err != nil {
		// Log the error but don't panic, as the data might already exist
		println("Warning: Could not insert test data into the books table. It might already exist.")
		return
	}

}

func insertTestDataFromFile(filepath string, db *sql.DB) error {
	// Open the SQL file
	sqlFile, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open SQL file: %v", err)
	}
	defer sqlFile.Close()

	// Read all content of the file
	sqlData, err := io.ReadAll(sqlFile)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %v", err)
	}

	// Execute the SQL file content
	_, err = db.Exec(string(sqlData))
	if err != nil {
		return fmt.Errorf("failed to execute SQL script: %v", err)
	}

	log.Println("SQL script executed successfully.")
	return nil
}
