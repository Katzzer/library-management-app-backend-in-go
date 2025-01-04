package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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
	    password TEXT NOT NULL 
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

// insertTestData: Inserts initial test data into the books table
func insertTestData() {
	insertBooks := `
	INSERT INTO books (name, author, description, isbn, image_name) VALUES
	('The Catcher in the Rye', 'J.D. Salinger', 'A story about adolescent alienation and loss of innocence.', '9780316769488', 'The_Catcher_in_the_Rye.jpg'),
	('To Kill a Mockingbird', 'Harper Lee', 'A novel of race and injustice in the Deep South.', '9780061120084', 'To_Kill_a_Mockingbird.jpg'),
	('1984', 'George Orwell', 'A dystopian novel set in a totalitarian regime.', '9780451524935', '1984.jpg'),
	('The Great Gatsby', 'F. Scott Fitzgerald', 'A critique of the American Dream in the Jazz Age.', '9780743273565', 'The_Great_Gatsby.jpg'),

	-- Programming Books
	('Effective Java', 'Joshua Bloch', 'Comprehensive coding practices for writing robust and modern Java programs.', '9780134685991', 'Effective_Java.jpg'),
	('The Go Programming Language', 'Alan A. A. Donovan and Brian W. Kernighan', 'A thorough introduction to programming with GoLang.', '9780134190440', 'The_Go_Programming_Language.jpg'),
	('Learning React', 'Alex Banks and Eve Porcello', 'A hands-on guide to building web interfaces using React.', '9781491954621', 'Learning_React.jpg'),
	('Next.js in Action', 'Liang Yi', 'Learn to build scalable and fast web applications with Next.js.', '9781617297343', 'Next.js_in_Action.jpg')
	`

	_, err := DB.Exec(insertBooks)
	if err != nil {
		// Log the error but don't panic, as the data might already exist
		println("Warning: Could not insert test data into the books table. It might already exist.")
	}
}
