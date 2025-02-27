import "./BookList.css"
import { useState, version } from "react"
const BookList = ({ books, onEdit, onRemove }) => {
  const [searchTerm, setSearchTerm] = useState("")

  const filteredBooks = books.filter(
    (book) =>
      book.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
      book.isbn?.toLowerCase().includes(searchTerm.toLowerCase()) ||
      book.publisher?.toLowerCase().includes(searchTerm.toLowerCase()),
  )
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
        <div key={book.id} className="book-item">
          <h3>{book.title}</h3>
          <p>
            <strong>ISBN:</strong> {book.isbn}
          </p>
          <p>
            <strong>Author:</strong> {book.authors}
          </p>
          <p>
          <p>
            <strong>Publisher:</strong> {book.publisher}
          </p>
            <strong>Version:</strong> {book.version}
          </p>
          <p>
            <strong>Total Copies:</strong> {book.total_copies}
          </p>
          <p>
            <strong>Available  Copies:</strong> {book.available_copies}
          </p>
          
          <div className="book-actions">
            <button className="edit-btn" onClick={() => onEdit(book)}>
              Edit
            </button>
            <button className="remove-btn" onClick={() => onRemove(book.id)}>
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

