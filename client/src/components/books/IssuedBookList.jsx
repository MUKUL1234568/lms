import { useState } from "react";
import { format } from "date-fns"; // Import date-fns for better date formatting
import { FaBook, FaUser, FaCalendarAlt, FaRegClock, FaCheckCircle, FaExclamationCircle } from "react-icons/fa"; // Icons for better visualization
import "./IssuedBookList.css";

const IssuedBookList = ({ issueregistry }) => {
  const [searchTerm, setSearchTerm] = useState("");

  // Filter issued books based on search term (email, status, or title)
  const filteredIssuedBooks = (issueregistry || []).filter(
    (issuedBook) =>
      issuedBook.user.email.toLowerCase().includes(searchTerm.toLowerCase()) ||
      issuedBook.issue_status.toLowerCase().includes(searchTerm.toLowerCase()) ||
      issuedBook.book.title.toLowerCase().includes(searchTerm.toLowerCase())
  );

  // Function to handle invalid date values
  const formatDate = (date) => {
    const parsedDate = new Date(date);
    return parsedDate instanceof Date && !isNaN(parsedDate) ? format(parsedDate, "dd MMM yyyy") : "Invalid Date";
  };

  return (
    <div className="issued-book-list-container">
      <h2 className="heading">Issued Books</h2>
      <div className="search-bar">
        <input
          type="text"
          placeholder="Search by Email, Title, or Status"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="search-input"
        />
      </div>
      <div className="issued-book-list">
        {filteredIssuedBooks.length > 0 ? (
          filteredIssuedBooks.map((issuedBook) => (
            <div key={issuedBook.issue_id} className="issued-book-item">
              <h3><FaBook /> {issuedBook.book.title}</h3>
              <p><FaUser /> <strong>Issued to:</strong> {issuedBook.user.name}</p>
              <p><FaUser /> <strong>User email:</strong> {issuedBook.user.email}</p>
              <p>
                <FaCalendarAlt /> <strong>Issue Date:</strong> {formatDate(issuedBook.issue_date)}
              </p>
              <p>
                <FaRegClock /> <strong>Return Date:</strong> {formatDate(issuedBook.return_date)}
              </p>
              <p>
                <strong>Status:</strong>
                <span className={`status ${issuedBook.issue_status.toLowerCase()}`}>
                  {issuedBook.issue_status === "returned" && <FaCheckCircle />}
                  {issuedBook.issue_status === "overdue" && <FaExclamationCircle />}
                  {issuedBook.issue_status}
                </span>
              </p>
            </div>
          ))
        ) : (
          <p className="no-issued-books">No issued books found.</p>
        )}
      </div>
    </div>
  );
};

export default IssuedBookList;
