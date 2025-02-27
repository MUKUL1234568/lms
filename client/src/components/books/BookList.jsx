import "./BookList.css"

const BookList = ({ books, onEdit, onRemove }) => {
  return (
    <div className="book-list">
      {books.map((book) => (
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
  )
}

export default BookList

