"use client"

import { useState } from "react"
import "./AdminManagement.css"

const AdminManagement = ({ users, onUpdateUserRole }) => {
  const [searchTerm, setSearchTerm] = useState("")

  const filteredUsers = users.filter(
    (user) =>
      user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.email.toLowerCase().includes(searchTerm.toLowerCase()),
  )

  const admins = filteredUsers.filter((user) => user.role === "LibraryAdmin")
  const nonAdmins = filteredUsers.filter((user) => user.role !== "LibraryAdmin"&&user.role !== "Owner")

  return (
    <div className="admin-management">
      <h2>Admin Management</h2>
      <div className="search-bar">
        <input
          type="text"
          placeholder="Search by name or email"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="search-input"
        />
      </div>
      <div className="admin-section">
        <h3>Current Admins</h3>
        <div className="user-list">
          {admins.map((user) => (
            <div key={user.id} className="user-item">
              <div>
                <h4>{user.name}</h4>
                <p>{user.email}</p>
              </div>
              <button onClick={() => onUpdateUserRole(user.id, "Reader")} className="revoke-admin-btn">
                Revoke Admin
              </button>
            </div>
          ))}
        </div>
      </div>
      <div className="non-admin-section">
        <h3>All Users</h3>
        <div className="user-list">
          {nonAdmins.map((user) => (
            <div key={user.id} className="user-item">
              <div>
                <h4>{user.name}</h4>
                <p>{user.email}</p>
              </div>
              <button onClick={() => onUpdateUserRole(user.id, "LibraryAdmin")} className="make-admin-btn">
                Make Admin
              </button>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

export default AdminManagement

