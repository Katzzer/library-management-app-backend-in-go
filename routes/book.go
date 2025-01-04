package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/models"
	"net/http"
	"strconv"
	"time"
)

// List all books
func getAllBooks(context *gin.Context) {
	books, err := models.GetAllBooks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve books", "error": err.Error()})
		return
	}

	// Transform books to add `lastBorrowedAtString` and `lastReturnedAtString`
	var response []gin.H
	for _, book := range books {
		response = append(response, gin.H{
			"id":                  book.ID,
			"name":                book.Name,
			"author":              book.Author,
			"description":         book.Description,
			"isbn":                book.ISBN,
			"image_name":          book.ImageName,
			"borrowed":            book.Borrowed,
			"last_borrowed_at":    formatTimeToString(book.LastBorrowedAt),
			"last_returned_at":    formatTimeToString(book.LastReturnedAt),
			"current_borrower_id": book.CurrentBorrowerID,
		})
	}

	context.JSON(http.StatusOK, response)
}

// Helper function to format *time.Time to a string or make it empty if nil
func formatTimeToString(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05") // Adjust format as needed
}

// Get details of a specific book
func getBook(context *gin.Context) {
	bookID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse book id"})
		return
	}

	book, err := models.GetBookByID(bookID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch book"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"book": book})
}

// Borrow a book
func borrowBook(context *gin.Context) {
	bookID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse book id"})
		return
	}

	userID := context.GetInt64("userId")

	// Attempt to borrow the book
	err = models.BorrowBook(bookID, userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Book borrowed successfully"})
}

// Return a book
func returnBook(context *gin.Context) {
	bookID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse book id"})
		return
	}

	userID := context.GetInt64("userId")

	// Attempt to return the book
	err = models.ReturnBook(bookID, userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}
