import "./Sidebar.css"

const Sidebar = ({ activeTab, setActiveTab }) => {
  return (
    <div className="sidebar">
      <button
        className={`sidebar-btn ${activeTab === "dashboard" ? "active" : ""}`}
        onClick={() => setActiveTab("dashboard")}
      >
        Overview
      </button>
      <button className={`sidebar-btn ${activeTab === "books" ? "active" : ""}`} onClick={() => setActiveTab("books")}>
        Books
      </button>
      <button
        className={`sidebar-btn ${activeTab === "requests" ? "active" : ""}`}
        onClick={() => setActiveTab("requests")}
      >
        Requests
      </button>
      <button className={`sidebar-btn ${activeTab === "users" ? "active" : ""}`} onClick={() => setActiveTab("users")}>
        Users
      </button>
      <button
        className={`sidebar-btn ${activeTab === "issued" ? "active" : ""}`}
        onClick={() => setActiveTab("issued")}
      >
        Issued Books
      </button>
    </div>
  )
}

export default Sidebar

