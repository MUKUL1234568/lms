import "./IssuedBookList.css"

const IssuedBookList = ({ issueregistry}) => {
  return (
    <div className="issued-book-list">
      { issueregistry.map((issuedBook) => {
         
        return (
          <div key={issuedBook.issue_id} className="issued-book-item">
            <h3>{issuedBook.book.title}</h3>
            <p>
              <strong>Issued to:</strong> {issuedBook.user.name}
            </p>
            <p>
              <strong>User email:</strong> {issuedBook.user.email}
            </p>
            <p>
              <strong>Issue Date:</strong> {issuedBook.issue_date}
            </p>
            <p>
              <strong>Return Date:</strong> {issuedBook.return_date}
            </p>
            <p>
              <strong>Status:</strong> {issuedBook.issue_status}
            </p>
          </div>
        )
      })}
    </div>
  )
}

export default IssuedBookList

