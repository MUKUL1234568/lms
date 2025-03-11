import { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import LoginModal from "../auth/LoginModal";
import SignUpModal from "../auth/SignUpModal"; // Import SignUpModal
import "./Header.css";
import { jwtDecode } from "jwt-decode";

function Header() {
  const navigate = useNavigate();
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [showLogin, setShowLogin] = useState(false);
  const [showSignUp, setShowSignUp] = useState(false); // Toggle sign-up modal

  useEffect(() => {
    const checkLoginStatus = () => {
      const token = localStorage.getItem("token");
      setIsLoggedIn(!!token);
    };

    checkLoginStatus();
    window.addEventListener("storage", checkLoginStatus);

    return () => window.removeEventListener("storage", checkLoginStatus);
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token");
    setIsLoggedIn(false);
    navigate("/");
  };

  const handleDashboard=()=>{
      const tokenn =localStorage.getItem("token")
      const decodetoken=jwtDecode(tokenn)

       const role=decodetoken.role
      if(role=="LibraryAdmin")
        navigate("/admindashboard")
      else if(role=="Owner")
        navigate("/ownerdashboard")
      else
      navigate("/readerdashboard")
  }
  const openLoginModal = () => {
    setShowSignUp(false); // Close sign-up if open
    setShowLogin(true); // Open login modal
  };

  const openSignUpModal = () => {
    setShowLogin(false); // Close login if open
    setShowSignUp(true); // Open sign-up modal
  };

  return (
    <header className="header">
      <Link to="/" className="logo">📚 Library System</Link>
      <nav>
        <ul>
          {isLoggedIn ? (
            <>
              <li><button onClick={handleDashboard}>Dashboard</button></li>
              <li><button className="logout-btn" onClick={handleLogout}>Logout</button></li>
            </>
          ) : (
            <>
              <li><button className="login-btn" onClick={openLoginModal}>Login</button></li>
              <li><button className="signup-btn" onClick={openSignUpModal}>SignUp</button></li>
            </>
          )}
        </ul>
      </nav>

      {/* Show Login or Sign-Up Modals */}
      {showLogin && <LoginModal onClose={() => setShowLogin(false)} onLogin={() => setIsLoggedIn(true)} />}
      {showSignUp && <SignUpModal onClose={() => setShowSignUp(false)} onRegisterSuccess={openLoginModal} />}
    </header>
  );
}

export default Header;
