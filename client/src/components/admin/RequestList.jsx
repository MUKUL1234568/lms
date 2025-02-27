 

import { useState } from "react"
import "./RequestList.css"

const RequestList = ({ requests, books, users, onApprove, onReject }) => {
  const [selectedActions, setSelectedActions] = useState({})
   
  
  const handleActionChange = (requestId, action) => {
    setSelectedActions((prevActions) => ({
      ...prevActions,
      [requestId]: action
    }));
  };
  
  const handleSubmit = (requestId) => {
    if (selectedActions[requestId] === "approve") {
      onApprove(requestId);
    } else if (selectedActions[requestId] === "reject") {
      onReject(requestId);
    }
  
    // Reset only the specific request's action
    setSelectedActions((prevActions) => ({
      ...prevActions,
      [requestId]: ""
    }));
  };
  

  return (
    <div className="request-list">
      {requests
        .map((request) => {
          const book = books.find((b) => b.isbn === request.book_id)
          const user = users.find((u) => u.id === request.reader_id)
          return (
            <div key={request.req_id} className="request-item">
              <p>
                <strong>{user.name}</strong> requested <strong><em>{book.title}</em></strong>
              </p>
              <p>
                <strong>Request Type: </strong>  <em>{request.request_type}</em> 
              </p>
              <p>
                <strong>Status: </strong>  <em>{request.status}</em> 
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
                  className="action-btn"
                  onClick={() => handleSubmit(request.req_id)}
                  disabled={!selectedActions[request.req_id]}
                >
                  Submit
                </button>
              </div>
            </div>
          )
        })}
    </div>
  )
}

export default RequestList

