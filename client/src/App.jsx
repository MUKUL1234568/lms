import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import { useState } from "react";
import { jwtDecode } from 'jwt-decode'; // Corrected default import
import Home from "./pages/Home.jsx";
import Admin from "./pages/LibraryAdminDashboard.jsx";
import Owner from "./pages/OwnerDashboard.jsx";
import UserDashboard from "./pages/UserDashboard.jsx";
import LoginModal from "./components/auth/LoginModal.jsx"; // Assuming LoginModal component exists

function App() {
  const [isModalOpen, setIsModalOpen] = useState(false); // State to control modal visibility
  const [redirectTo, setRedirectTo] = useState(null); // State to store the route to redirect to after login
  const token = localStorage.getItem("token"); // Get the token from localStorage
  let userRole = null;

  // Decode the token to extract the user role (if the token exists)
  if (token) {
    const decodedToken = jwtDecode(token); // Using jwt_decode function correctly
    userRole = decodedToken?.role; // Extract the role from the decoded token
  }

  // Function to handle login success and redirect
  const handleLoginSuccess = () => {
    setIsModalOpen(false); // Close the modal after login
    if (redirectTo) {
      window.location.href = redirectTo; // Redirect to the protected route the user was trying to access
    }
  };

  // Protected Route Wrapper that opens modal if not logged in
  const ProtectedRoute = ({ element, roleRequired, path }) => {
    if (!token) {
      setRedirectTo(path); // Store the protected route path to redirect after login
      setIsModalOpen(true); // Open the login modal
      return null; // Don't render the route, just show the modal
    }

    // Check if the role from the decoded token matches the required role for this route
    if (userRole !== roleRequired) {
      return <Navigate to="/" />; // Redirect to the home page if the user doesn't have access to this route
    }

    return element; // If the role matches, render the protected route
  };

  return (
    <Router>
      <Routes>
        {/* Public Route */}
        <Route path="/" element={<Home />} />

        {/* Protected Routes with Role-Based Access */}
        <Route
          path="/admindashboard"
          element={<ProtectedRoute element={<Admin />} roleRequired="LibraryAdmin" path="/admindashboard" />}
        />
        <Route
          path="/ownerdashboard"
          element={<ProtectedRoute element={<Owner />} roleRequired="Owner" path="/ownerdashboard" />}
        />
        <Route
          path="/readerdashboard"
          element={<ProtectedRoute element={<UserDashboard />} roleRequired="Reader" path="/readerdashboard" />}
        />
      </Routes>

      {/* Login Modal */}
      {isModalOpen && <LoginModal onClose={() => setIsModalOpen(false)} onLogin={handleLoginSuccess} />}
    </Router>
  );
}

export default App;
