 "use client"

import "./BookList.css"
import { useState ,useEffect} from "react"
import { FaBook, FaUser, FaBarcode, FaBuilding, FaTrash, FaEdit, FaCopy } from "react-icons/fa"
 

const BookList = ({books,onEdit, onRemove }) => {
  const [searchTerm, setSearchTerm] = useState("")
  

  
  const filteredBooks = books.filter(
    (book) =>
      book.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
      book.isbn?.toLowerCase().includes(searchTerm.toLowerCase()) ||
      book.publisher?.toLowerCase().includes(searchTerm.toLowerCase())
  )

  const handleRemove = (book) => {
    const hasPendingRequests =
      book.requests &&
      book.requests.some(
        (request) => request.status.toLowerCase() === "pending" && request.isbn === book.isbn
      )

    const hasIssuedRecords =
      book.issue_records &&
      book.issue_records.some(
        (record) => record.issue_status.toLowerCase() === "issued" && record.isbn === book.isbn
      )

    if (hasPendingRequests || hasIssuedRecords) {
      alert("This book has pending requests or is currently issued. It cannot be removed.")
      return
    }

    onRemove(book.isbn)
  }
  

  return (
    <div className="book-list-container">
      {/* Search Bar */}
      <div className="search-bar">
        <input
          type="text"
          placeholder="🔍 Search by title, ISBN, or publisher"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="search-input"
        />
      </div>

      {/* Book Cards Grid */}
      <div className="book-list">
        {filteredBooks.length > 0 ? (
          filteredBooks.map((book) => (
            <div key={book.isbn} className="book-item">
              <h3>
                <FaBook className="book-icon" /> {book.title}
              </h3>
              <p>
                <FaBarcode className="icon" /> <strong>ISBN:</strong> {book.isbn}
              </p>
              <p>
                <FaUser className="icon" /> <strong>Author:</strong> {book.authors}
              </p>
              <p>
                <FaBuilding className="icon" /> <strong>Publisher:</strong> {book.publisher}
              </p>
              <p>
                <FaCopy className="icon" /> <strong>Version:</strong> {book.version}
              </p>
              <p>
                <strong>Total Copies:</strong> {book.total_copies}
              </p>
              <p>
                <strong>Available Copies:</strong> {book.available_copies}
              </p>

              <div className="book-actions">
                <button className="edit-btn" onClick={() => onEdit(book)}>
                  <FaEdit className="action-icon" /> Edit
                </button>
                <button className="remove-btn" onClick={() => handleRemove(book)}>
                  <FaTrash className="action-icon" /> Remove
                </button>
              </div>
            </div>
          ))
        ) : (
          <p className="no-books">No books found.</p>
        )}
      </div>
    </div>
  )
}

export default BookList
