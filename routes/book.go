package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/models"
	"net/http"
	"strconv"
)

// List all books
func getAllBooks(context *gin.Context) {
	books, err := models.GetAllBooks()
	if err != nil {
		fmt.Printf("Error fetching books: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch books"})
		return
	}
	context.JSON(http.StatusOK, books)
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
