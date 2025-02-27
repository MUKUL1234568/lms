import "./IssuedBookList.css"

const IssuedBookList = ({ issuedBooks, books, users }) => {
  return (
    <div className="issued-book-list">
      {issuedBooks.map((issuedBook) => {
        const book = books.find((b) => b.isbn === issuedBook.bookId)
        const user = users.find((u) => u.id === issuedBook.userId)
        return (
          <div key={issuedBook.id} className="issued-book-item">
            <h3>{book.title}</h3>
            <p>
              <strong>Issued to:</strong> {user.name}
            </p>
            <p>
              <strong>User email:</strong> {user.email}
            </p>
            <p>
              <strong>Issue Date:</strong> {issuedBook.issueDate}
            </p>
            <p>
              <strong>Return Date:</strong> {issuedBook.returnDate}
            </p>
            <p>
              <strong>Status:</strong> {issuedBook.status || "Issued"}
            </p>
          </div>
        )
      })}
    </div>
  )
}

export default IssuedBookList

