import "./Sidebar.css"
import { FaTachometerAlt, FaBook, FaUsers, FaClipboardList, FaBookOpen, FaUserShield } from "react-icons/fa"

const Sidebar = ({ activeTab, setActiveTab }) => {
  return (
    <div className="sidebar">
      <button
        className={`sidebar-btn ${activeTab === "dashboard" ? "active" : ""}`}
        onClick={() => setActiveTab("dashboard")}
      >
        <FaTachometerAlt className="sidebar-icon" /> Overview
      </button>
      <button className={`sidebar-btn ${activeTab === "books" ? "active" : ""}`} onClick={() => setActiveTab("books")}>
        <FaBook className="sidebar-icon" /> Books
      </button>
      <button
        className={`sidebar-btn ${activeTab === "requests" ? "active" : ""}`}
        onClick={() => setActiveTab("requests")}
      >
        <FaClipboardList className="sidebar-icon" /> Requests
      </button>
      <button className={`sidebar-btn ${activeTab === "users" ? "active" : ""}`} onClick={() => setActiveTab("users")}>
        <FaUsers className="sidebar-icon" /> Users
      </button>
      <button
        className={`sidebar-btn ${activeTab === "issued" ? "active" : ""}`}
        onClick={() => setActiveTab("issued")}
      >
        <FaBookOpen className="sidebar-icon" /> Issued Books
      </button>
      <button className={`sidebar-btn ${activeTab === "adminmanagement" ? "active" : ""}`} onClick={() => setActiveTab("adminmanagement")}>
        <FaUserShield className="sidebar-icon" /> Admin Management
      </button>
    </div>
  )
}

export default Sidebar
