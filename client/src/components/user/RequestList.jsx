import { useState } from "react";
import { format } from "date-fns"; // Importing date formatting library
import {
  FaBook,
  FaCalendarAlt,
  FaCheckCircle,
  FaExclamationCircle,
  FaSearch,
  FaTimesCircle,
} from "react-icons/fa"; // Icons for better UI
import "./RequestList.css";

const RequestList = ({ requests = [] }) => {
  const [searchTerm, setSearchTerm] = useState("");

  // Function to safely format dates
  const formatDate = (date) => {
    const parsedDate = new Date(date);
    return parsedDate instanceof Date && !isNaN(parsedDate)
      ? format(parsedDate, "dd MMM yyyy")
      : "N/A";
  };

  // Filter requests based on search input (ISBN, status)
  const filteredRequests = requests.filter(
    (request) =>
      request.isbn.toLowerCase().includes(searchTerm.toLowerCase()) ||
      request.status.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="request-list-container">
      <h2 className="heading">📚 My Requests</h2>

      {/* Search Bar */}
      <div className="search-bar">
        <FaSearch className="search-icon" />
        <input
          type="text"
          placeholder="Search by ISBN or Status"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="search-input"
        />
      </div>

      {/* Request Cards Grid */}
      <div className="request-grid">
        {filteredRequests.length > 0 ? (
          filteredRequests.map((request) => (
            <div key={request.req_id} className="request-card">
              <h3>
                <FaBook /> Request for ISBN: {request.isbn}
              </h3>
              <p>
                <FaCalendarAlt /> <strong>Request Date:</strong> {formatDate(request.request_date)}
              </p>
              <p>
                {request.status.toLowerCase() === "approved" ? (
                  <FaCheckCircle className="status-icon approved" />
                ) : request.status.toLowerCase() === "rejected" ? (
                  <FaTimesCircle className="status-icon rejected" />
                ) : (
                  <FaExclamationCircle className="status-icon pending" />
                )}
                <strong>Status:</strong> {request.status}
              </p>
              <p>
                <strong>Request Type:</strong> {request.request_type}
              </p>
              {request.approval_date && (
                <p>
                  <FaCalendarAlt /> <strong>Approval Date:</strong> {formatDate(request.approval_date)}
                </p>
              )}
            </div>
          ))
        ) : (
          <p className="no-requests">No matching requests found.</p>
        )}
      </div>
    </div>
  );
};

export default RequestList;
