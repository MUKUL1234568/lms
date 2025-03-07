 "use client"

import { useState } from "react"
import "./BookList.css"

const BookList = ({ books, onRequestBook }) => {
  const [searchTerm, setSearchTerm] = useState("")

  const filteredBooks = books.filter(
    (book) =>
      book.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
      book.isbn.toLowerCase().includes(searchTerm.toLowerCase()) ||
      book.publisher.toLowerCase().includes(searchTerm.toLowerCase()),
  )

  return (
    <div className="book-list-container">
      <h2>Available Books</h2>
      {/* Search Bar */}
      <div className="search-bar">
        <input
          type="text"
          placeholder="Search by title, ISBN, or publisher"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="search-input"
        />
      </div>
      {/* Book Cards */}
      <div className="book-list">
        {filteredBooks.map((book) => (
          <div key={book.isbn} className="book-item">
            <h3>{book.title}</h3>
            <p><strong>Author(s):</strong> {book.authors}</p>
            <p><strong>Publisher:</strong> {book.publisher}</p>
            <p><strong>ISBN:</strong> {book.isbn}</p>
            <p><strong>Available Copies:</strong> {book.available_copies}</p>
            <button
              onClick={() => onRequestBook(book.isbn)}
              disabled={book.available_copies === 0}
            >
              {book.available_copies > 0 ? "Request Book" : "Not Available"}
            </button>
          </div>
        ))}
      </div>
    </div>
  )
}

export default BookList
