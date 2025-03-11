import { useState } from "react";
import axios from "axios";
import { FaBook, FaCheckCircle } from "react-icons/fa";
import "./IssuedBookList.css";

const IssuedBookList = ({ issuedBooks, fetchUserProfile }) => {
  const [isSubmitting, setIsSubmitting] = useState({});

  const handleReturnRequest = async (isbn) => {
    setIsSubmitting((prev) => ({ ...prev, [isbn]: true }));
    try {
      const token = localStorage.getItem("token");
      const response = await axios.post(
        "http://localhost:8080/request/",
        { isbn, request_type: "Return" },
        { headers: { Authorization: `Bearer ${token}` } }
      );

      if (response.status === 201) {
        alert("Return request submitted successfully");
        if (fetchUserProfile) fetchUserProfile();
      } else {
        alert(response.data.error || "Failed to submit return request.");
      }
    } catch (error) {
      console.error("Error requesting return:", error);
      alert(error.response?.data?.message || "Failed to submit return request. Please try again.");
    } finally {
      setIsSubmitting((prev) => ({ ...prev, [isbn]: false }));
    }
  };

  return (
    <div className="issued-books-wrapper">
      <h2>Issued Books</h2>
      <div className="book-list">
        {issuedBooks.length > 0 ? (
          issuedBooks.map((book) => (
            <div key={book.issue_id} className={`book-card ${book.issue_status.toLowerCase()}`}>
              <FaBook className="book-icon" />
              <h3>{book.isbn}</h3>
              <p><strong>Issue Date:</strong> {new Date(book.issue_date).toLocaleDateString()}</p>
              {book.issue_status === "Issued" ? (
                <>
                  <p><strong>Expected Return:</strong> {new Date(book.expected_return_date).toLocaleDateString()}</p>
                  <p><strong>Status:</strong> Issued</p>
                  <button 
                    className="return-btn" 
                    onClick={() => handleReturnRequest(book.isbn)}
                    disabled={isSubmitting[book.isbn]}
                  >
                    {isSubmitting[book.isbn] ? "Processing..." : "Request Return"}
                  </button>
                </>
              ) : (
                <>
                  <p><strong>Return Date:</strong> {new Date(book.return_date).toLocaleDateString()}</p>
                  <div className="returned-status">
                    <FaCheckCircle className="check-icon" /> Returned
                  </div>
                </>
              )}
            </div>
          ))
        ) : (
          <p className="no-books">No books found.</p>
        )}
      </div>
    </div>
  );
};

export default IssuedBookList;
