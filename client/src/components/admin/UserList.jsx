 "use client"

import { useState } from "react"
import "./UserList.css"

const UserList = ({ users }) => {
  const [selectedUser, setSelectedUser] = useState(null)
  const [searchTerm, setSearchTerm] = useState("")

  // Filter users based on search term and exclude "LibraryAdmin" & "Owner"
  const filteredUsers = users.filter(
    (user) =>
      (user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        user.email.toLowerCase().includes(searchTerm.toLowerCase())) &&
      user.role !== "LibraryAdmin" &&
      user.role !== "Owner"
  )

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
          <div key={user.id} className="user-item" onClick={() => setSelectedUser(user)}>
            <h3>{user.name}</h3>
            <p>
              <strong>Email:</strong> {user.email}
            </p>
          </div>
        ))}
      </div>

      {selectedUser && (
        <div className="user-details">
          <h2>{selectedUser.name}'s Details</h2>
          <p>
            <strong>Email:</strong> {selectedUser.email}
          </p>
          <p>
            <strong>Phone No:</strong> {selectedUser.contact_number}
          </p>

          <h3>Issued Books</h3>
          <div className="issued-books-container">
            {selectedUser.issue_records.map((issuedBook) => (
              <div key={issuedBook.issue_id} className="issued-book-card">
                <p><strong>ISBN:</strong> {issuedBook.isbn}</p>
                <p><strong>Issued Date:</strong> {issuedBook.issue_date}</p>
                <p><strong>Expected Return:</strong> {issuedBook.expected_return_date}</p>
                <p className={`status ${issuedBook.issue_status.toLowerCase()}`}>
                  <strong>Status:</strong> {issuedBook.issue_status}
                </p>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  )
}

export default UserList
