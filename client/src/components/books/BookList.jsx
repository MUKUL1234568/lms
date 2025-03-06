import "./BookList.css"
import { useState } from "react"

const BookList = ({ books, onEdit, onRemove }) => {
  const [searchTerm, setSearchTerm] = useState("")

  const filteredBooks = books.filter(
    (book) =>
      book.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
      book.isbn?.toLowerCase().includes(searchTerm.toLowerCase()) ||
      book.publisher?.toLowerCase().includes(searchTerm.toLowerCase())
  )

  const handleRemove = (book) => {
    // Check if there are any pending requests for this book
    const hasPendingRequests =
      book.requests && book.requests.some(
        (request) => request.status.toLowerCase() === "pending" && request.isbn === book.isbn
      )

    // Check if this book is issued to any user
    const hasIssuedRecords =
      book.issue_records && book.issue_records.some(
        (record) => record.status.toLowerCase() === "issued" && record.isbn === book.isbn
      )

    if (hasPendingRequests || hasIssuedRecords) {
      alert("This book has pending requests or is currently issued to a user. It cannot be removed.")
      return
    }

    // If no issues, proceed with removal
    onRemove(book.isbn)
  }

  return (
    <div className="book-list-container">
      <div className="search-bar">
        <input
          type="text"
          placeholder="Search by title, ISBN, or publisher"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="search-input"
        />
      </div>
      <div className="book-list">
        {filteredBooks.map((book) => (
          <div key={book.isbn} className="book-item">
            <h3>{book.title}</h3>
            <p><strong>ISBN:</strong> {book.isbn}</p>
            <p><strong>Author:</strong> {book.authors}</p>
            <p><strong>Publisher:</strong> {book.publisher}</p>
            <p><strong>Version:</strong> {book.version}</p>
            <p><strong>Total Copies:</strong> {book.total_copies}</p>
            <p><strong>Available Copies:</strong> {book.available_copies}</p>

            <div className="book-actions">
              <button className="edit-btn" onClick={() => onEdit(book)}>
                Edit
              </button>
              <button className="remove-btn" onClick={() => handleRemove(book)}>
                Remove
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default BookList
