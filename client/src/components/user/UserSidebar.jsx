"use client"
import "./UserSidebar.css"

const UserSidebar = ({ activeTab, setActiveTab }) => {
  return (
    <div className="sidebar">
      <button className={`sidebar-btn ${activeTab === "books" ? "active" : ""}`} onClick={() => setActiveTab("books")}>
        All Books
      </button>
      <button
        className={`sidebar-btn ${activeTab === "requests" ? "active" : ""}`}
        onClick={() => setActiveTab("requests")}
      >
        My Requests
      </button>
      <button
        className={`sidebar-btn ${activeTab === "issued" ? "active" : ""}`}
        onClick={() => setActiveTab("issued")}
      >
        Issued Books
      </button>
      <button
        className={`sidebar-btn ${activeTab === "profile" ? "active" : ""}`}
        onClick={() => setActiveTab("profile")}
      >
        My Profile
      </button>
    </div>
  )
}

export default UserSidebar

