"use client"

import { useState, useEffect } from "react"
import axios from "axios"
import UserSidebar from "../components/user/UserSidebar"
import Header from "../components/layout/Header"
import BookList from "../components/user/BookList"
import RequestList from "../components/user/RequestList"
import IssuedBookList from "../components/user/IssuedBooklist"
 import UserProfile from "../components/user/UserProfile"
import "./UserDashboard.css"

const UserDashboard = () => {
  const [activeTab, setActiveTab] = useState("books")
   
  const [userProfile, setUserProfile] = useState(null)

  useEffect(() => {
     
    fetchUserProfile()
  }, [])

  
  const fetchUserProfile = async () => {
    try {
      const token = localStorage.getItem("token")
      const response = await axios.get("http://localhost:8080/user/profile", {
        headers: { Authorization: `Bearer ${token}` },
      })
      setUserProfile(response.data.user)
    } catch (error) {
      console.error("Error fetching user profile:", error)
    }
  }

  const handleBookRequest = async (isbn) => {
    try {
      const token = localStorage.getItem("token");
      const response = await axios.post(
        "http://localhost:8080/request/",
        { isbn, request_type: "Issue" },
        { headers: { Authorization: `Bearer ${token}` } }
      );
  
      if (response.status === 201) {
        alert("Book request submitted successfully");
        fetchUserProfile(); // Refresh user profile to update requests
      } else {
        alert(response.data.error || "Failed to submit book request.");
      }
    } catch (error) {
      console.error("Error requesting book:", error);
      alert(error.response?.data?.message || "Failed to submit book request. Please try again.");
    }
  };
  
  return (
    <><Header/>
    <div className="user-dashboard">
        
      <UserSidebar activeTab={activeTab} setActiveTab={setActiveTab} />
      <div className="dashboard-content">
        {activeTab === "books" && <BookList   onRequestBook={handleBookRequest} />}
        {activeTab === "requests" && userProfile && <RequestList requests={userProfile.requests}  />}
        {activeTab === "issued" && userProfile && <IssuedBookList issuedBooks={userProfile.issue_records} fetchUserProfile={fetchUserProfile} />}
        {activeTab === "profile" && userProfile && <UserProfile user={userProfile} />}
      </div>
    </div>
    </>
  )
}

export default UserDashboard

