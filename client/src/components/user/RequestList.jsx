import { useState } from "react"
import "./RequestList.css"

const RequestList = ({ requests }) => {
  const [searchTerm, setSearchTerm] = useState("")

  const filteredRequests = requests.filter(
    (request) =>
      request.isbn.toLowerCase().includes(searchTerm.toLowerCase())
  )

  return (
    <div className="request-list-container">
      <h2>My Requests</h2>
      {/* Search Bar */}
      <div className="search-bar">
        <input
          type="text"
          placeholder="Search by ISBN"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="search-input"
        />
      </div>
      {/* Request Cards */}
      <div className="request-grid">
        {filteredRequests.map((request) => (
          <div key={request.req_id} className="request-card">
            <h3>Request for ISBN: {request.isbn}</h3>
            <p>
              <strong>Request Date:</strong> {new Date(request.request_date).toLocaleDateString()}
            </p>
            <p>
              <strong>Status:</strong> {request.status}
            </p>
            {request.approval_date && (
              <p>
                <strong>Approval Date:</strong> {new Date(request.approval_date).toLocaleDateString()}
              </p>
            )}
          </div>
        ))}
      </div>
    </div>
  )
}

export default RequestList
