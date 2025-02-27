"use client"

import { useState } from "react"
import "./UserList.css"

const UserList = ({ users, issuedBooks, books }) => {
    console.log(books)
  const [selectedUser, setSelectedUser] = useState(null)
  const [searchTerm, setSearchTerm] = useState("")
   
  const getUserIssuedBooks = (userId) => {
    
    if (!issuedBooks || !Array.isArray(issuedBooks)) return [] // Handle undefined case
    return issuedBooks
      .filter((book) => book.userId === userId)
      .map((issuedBook) => {
        const book = books.find((b) => b.isbn === issuedBook.bookId) || {} // Avoid undefined errors
        return { ...issuedBook, title: book.title || "Unknown Title" }
      })
  }
  

  const filteredUsers = users
  .filter((user) =>
    (user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
     user.email.toLowerCase().includes(searchTerm.toLowerCase())) &&
    user.role !== "LibraryAdmin" && user.role !== "Owner" // Exclude admin and owner
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
            // if(user.role=="LibraryAdmin"||"owner")
            //     continue
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
          <ul className="issued-books-list">
            {getUserIssuedBooks(selectedUser.id).map((book) => (
              <li key={book.id}>
                {book.title} - Issued: {book.issueDate}, Return: {book.returnDate}
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  )
}

export default UserList

