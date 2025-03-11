"use client"

import { useState } from "react"
import "./RequestList.css"
import { FaUser, FaEnvelope, FaBarcode, FaBook, FaCheck, FaTimes } from "react-icons/fa"

const RequestList = ({ requests = [], onApprove, onReject }) => {
  const [selectedActions, setSelectedActions] = useState({})
  const [searchTerm, setSearchTerm] = useState("")

  const handleActionChange = (requestId, action) => {
    setSelectedActions((prevActions) => ({
      ...prevActions,
      [requestId]: action,
    }))
  }

  const handleSubmit = (requestId) => {
    if (selectedActions[requestId] === "approve") {
      onApprove(requestId)
    } else if (selectedActions[requestId] === "reject") {
      onReject(requestId)
    }

    setSelectedActions((prevActions) => ({
      ...prevActions,
      [requestId]: "",
    }))
  }

  const filteredRequests = requests.filter(
    (request) =>
      request.status === "pending" &&
      (request.user.email.toLowerCase().includes(searchTerm.toLowerCase()) ||
        request.isbn.toLowerCase().includes(searchTerm.toLowerCase()))
  )

  return (
    <div className="request-list-container">
      <h2>📜 Pending Requests</h2>

      {/* Search Bar */}
      <div className="search-bar">
        <input
          type="text"
          placeholder="🔍 Search by Email or ISBN"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="search-input"
        />
      </div>

      {/* Request Grid */}
      <div className="request-list">
        {filteredRequests.length > 0 ? (
          filteredRequests.map((request) => (
            <div key={request.req_id} className="request-item">
              <p className="request-info">
                <FaUser className="icon" /> <strong>{request.user.name}</strong>  
              </p>
              <p className="request-info">
                <FaEnvelope className="icon" /> <em>{request.user.email}</em>
              </p>
              <p className="request-info">
                <FaBook className="icon" /> <strong>Book:</strong> <em>{request.book.title}</em>
              </p>
              <p className="request-info">
                <FaBarcode className="icon" /> <strong>ISBN:</strong> <em>{request.isbn}</em>
              </p>
              <p className="request-info">
                <strong>Request Type:</strong> <em>{request.request_type}</em>
              </p>

              <div className="request-actions">
                <select
                  value={selectedActions[request.req_id] || ""}
                  onChange={(e) => handleActionChange(request.req_id, e.target.value)}
                  className="action-select"
                >
                  <option value="">Select Action</option>
                  <option value="approve">Approve</option>
                  <option value="reject">Reject</option>
                </select>
                <button
                  className={`action-btn ${
                    selectedActions[request.req_id] === "approve" ? "approve" : "reject"
                  }`}
                  onClick={() => handleSubmit(request.req_id)}
                  disabled={!selectedActions[request.req_id]}
                >
                  {selectedActions[request.req_id] === "approve" ? (
                    <>
                      <FaCheck className="action-icon" /> Approve
                    </>
                  ) : (
                    <>
                      <FaTimes className="action-icon" /> Reject
                    </>
                  )}
                </button>
              </div>
            </div>
          ))
        ) : (
          <p className="no-requests">No pending requests found.</p>
        )}
      </div>
    </div>
  )
}

export default RequestList
