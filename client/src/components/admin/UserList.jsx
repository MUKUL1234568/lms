 "use client";

import { useState } from "react";
import { FaSearch, FaUser, FaEnvelope, FaPhone, FaBook } from "react-icons/fa";
import { format } from "date-fns";
import "./UserList.css";

const UserList = ({ users }) => {
  const [selectedUser, setSelectedUser] = useState(null);
  const [searchTerm, setSearchTerm] = useState("");

  const filteredUsers = users.filter(
    (user) =>
      (user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        user.email.toLowerCase().includes(searchTerm.toLowerCase())) &&
      user.role !== "LibraryAdmin" &&
      user.role !== "Owner"
  );

  const formatDate = (date) => {
    return format(new Date(date), "dd MMM yyyy");
  };

  return (
    <div className="user-list-container">
      <div className="search-bar"> 
     
        <input
         
          type="text"
          placeholder="Search by name or email"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="search-input"
          
        />  
      </div>

      <div className="user-list">
        {filteredUsers.map((user) => (
          <div
            key={user.id}
            className={`user-item ${selectedUser?.id === user.id ? "selected" : ""}`}
            onClick={() => setSelectedUser(user)}
          >
            <h3><FaUser className="icon" /> {user.name}</h3>
            <p><FaEnvelope className="icon" /> <strong>Email:</strong> {user.email}</p>
          </div>
        ))}
      </div>

      {selectedUser && (
        <div className="user-details">
          <h2><FaUser className="icon" /> {selectedUser.name}'s Details</h2>
          <p><FaEnvelope className="icon" /> <strong>Email:</strong> {selectedUser.email}</p>
          <p><FaPhone className="icon" /> <strong>Phone No:</strong> {selectedUser.contact_number}</p>

          <h3><FaBook className="icon" /> Issued Books</h3>
          <div className="issued-books-container">
            {selectedUser.issue_records.length > 0 ? (
              selectedUser.issue_records.map((issuedBook) => (
                <div key={issuedBook.issue_id} className="issued-book-card">
                  <p><strong>ISBN:</strong> {issuedBook.isbn}</p>
                  <p><strong>Issued Date:</strong> {formatDate(issuedBook.issue_date)}</p>
                  <p><strong>Expected Return:</strong> {formatDate(issuedBook.expected_return_date)}</p>
                  <p className={`status ${issuedBook.issue_status.toLowerCase()}`}>
                    <strong>Status:</strong> {issuedBook.issue_status}
                  </p>
                </div>
              ))
            ) : (
              <p className="no-books">No books issued to this user.</p>
            )}
          </div>
        </div>
      )}
    </div>
  );
};

export default UserList;