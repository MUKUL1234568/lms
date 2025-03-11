 "use client"

import { useState } from "react"
import { FaBook, FaUser, FaBuilding, FaBarcode, FaSearch } from "react-icons/fa"
import "./BookList.css"

const BookList = ({ books = [], onRequestBook }) => {
  const [searchTerm, setSearchTerm] = useState("")

  const filteredBooks = books.filter(
    (book) =>
      book.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
      book.isbn.toLowerCase().includes(searchTerm.toLowerCase()) ||
      book.publisher.toLowerCase().includes(searchTerm.toLowerCase())
  )

  return (
    <div className="book-list-container">
      <h2 className="heading">📚 Available Books</h2>

      {/* Search Bar */}
      <div className="search-bar">
        <FaSearch className="search-icon" />
        <input
          type="text"
          placeholder="Search by Title, ISBN, or Publisher"
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
                <FaUser className="icon" /> <strong>Author(s):</strong> {book.authors}
              </p>
              <p>
                <FaBuilding className="icon" /> <strong>Publisher:</strong> {book.publisher}
              </p>
              <p>
                <FaBarcode className="icon" /> <strong>ISBN:</strong> {book.isbn}
              </p>
              <p>
                <strong>Available Copies:</strong> {book.available_copies}
              </p>
              <button
                onClick={() => onRequestBook(book.isbn)}
                disabled={book.available_copies === 0}
                className={book.available_copies > 0 ? "request-btn" : "disabled-btn"}
              >
                {book.available_copies > 0 ? "Request Book" : "Not Available"}
              </button>
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
