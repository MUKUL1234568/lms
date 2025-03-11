 "use client"
import "./UserSidebar.css"
import { FaBook, FaClipboardList, FaBookOpen, FaUser } from "react-icons/fa"

const UserSidebar = ({ activeTab, setActiveTab }) => {
  return (
    <div className="sidebar">
      <button className={`sidebar-btn ${activeTab === "books" ? "active" : ""}`} onClick={() => setActiveTab("books")}>
        <FaBook className="sidebar-icon" /> All Books
      </button>
      <button
        className={`sidebar-btn ${activeTab === "requests" ? "active" : ""}`}
        onClick={() => setActiveTab("requests")}
      >
        <FaClipboardList className="sidebar-icon" /> My Requests
      </button>
      <button
        className={`sidebar-btn ${activeTab === "issued" ? "active" : ""}`}
        onClick={() => setActiveTab("issued")}
      >
        <FaBookOpen className="sidebar-icon" /> Issued Books
      </button>
      <button
        className={`sidebar-btn ${activeTab === "profile" ? "active" : ""}`}
        onClick={() => setActiveTab("profile")}
      >
        <FaUser className="sidebar-icon" /> My Profile
      </button>
    </div>
  )
}

export default UserSidebar
