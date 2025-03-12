import { useState, useEffect, useRef } from "react";
import { Link, useNavigate, useLocation } from "react-router-dom";
import LoginModal from "../auth/LoginModal";
import SignUpModal from "../auth/SignUpModal";
import "./Header.css";
import { jwtDecode } from "jwt-decode";

function Header() {
  const navigate = useNavigate();
  const location = useLocation();
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [showLogin, setShowLogin] = useState(false);
  const [showSignUp, setShowSignUp] = useState(false);
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
  const [userRole, setUserRole] = useState(null);
  const [showDropdown, setShowDropdown] = useState(false);
  const dropdownRef = useRef(null);

  useEffect(() => {
    const checkLoginStatus = () => {
      const token = localStorage.getItem("token");
      setIsLoggedIn(!!token);
      
      if (token) {
        try {
          const decoded = jwtDecode(token);
          setUserRole(decoded.role);
        } catch (error) {
          console.error("Invalid token", error);
        }
      }
    };

    checkLoginStatus();
    window.addEventListener("storage", checkLoginStatus);

    // Close mobile menu when route changes
    setMobileMenuOpen(false);

    return () => window.removeEventListener("storage", checkLoginStatus);
  }, [location.pathname]);

  // Close dropdown when clicking outside
  useEffect(() => {
    function handleClickOutside(event) {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
        setShowDropdown(false);
      }
    }

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token");
    setIsLoggedIn(false);
    setUserRole(null);
    navigate("/");
  };

  const handleDashboard = () => {
    if (userRole === "LibraryAdmin") {
      navigate("/admindashboard");
    } else if (userRole === "Owner") {
      navigate("/ownerdashboard");
    } else {
      navigate("/readerdashboard");
    }
  };

  const openLoginModal = () => {
    setShowSignUp(false);
    setShowLogin(true);
    setMobileMenuOpen(false);
  };

  const openSignUpModal = () => {
    setShowLogin(false);
    setShowSignUp(true);
    setMobileMenuOpen(false);
  };

  const toggleMobileMenu = () => {
    setMobileMenuOpen(!mobileMenuOpen);
  };

  const toggleDropdown = () => {
    setShowDropdown(!showDropdown);
  };

  return (
    <header className="header">
      <div className="header-container">
        <Link to="/" className="logo">
          <span className="logo-icon">📚</span>
          <span className="logo-text"> LibraEase</span>
        </Link>

        {/* Desktop Navigation */}
        <nav className="desktop-nav">
          <ul className="nav-links">
            <li>
              <Link to="/" className={location.pathname === "/" ? "active" : ""}>
                Home
              </Link>
            </li>
            <li>
              <Link to="/about" className={location.pathname === "/about" ? "active" : ""}>
                About
              </Link>
            </li>
            <li>
              <Link to="/services" className={location.pathname === "/services" ? "active" : ""}>
                Services
              </Link>
            </li>
            <li>
              <Link to="/contact" className={location.pathname === "/contact" ? "active" : ""}>
                Contact
              </Link>
            </li>
          </ul>

          <div className="auth-buttons">
            {isLoggedIn ? (
              <div className="user-menu" ref={dropdownRef}>
                <button className="user-button" onClick={toggleDropdown}>
                  <span className="user-icon">👤</span>
                  <span className="user-role">{userRole || "User"}</span>
                  <span className={`dropdown-arrow ${showDropdown ? "up" : "down"}`}>▼</span>
                </button>
                {showDropdown && (
                  <div className="dropdown-menu">
                    <button onClick={handleDashboard} className="dropdown-item">
                      <span className="dropdown-icon">🔍</span> Dashboard
                    </button>
                    {/* <button onClick={() => navigate("/profile")} className="dropdown-item">
                      <span className="dropdown-icon">👤</span> Profile
                    </button> */}
                    {/* <button onClick={() => navigate("/settings")} className="dropdown-item">
                      <span className="dropdown-icon">⚙️</span> Settings
                    </button> */}
                    <div className="dropdown-divider"></div>
                    <button onClick={handleLogout} className="dropdown-item logout">
                      <span className="dropdown-icon">🚪</span> Logout
                    </button>
                  </div>
                )}
              </div>
            ) : (
              <>
                <button className="login-btn" onClick={openLoginModal}>
                  Login
                </button>
                <button className="signup-btn" onClick={openSignUpModal}>
                  Sign Up
                </button>
              </>
            )}
          </div>
        </nav>

        {/* Mobile Menu Button */}
        <button className="mobile-menu-button" onClick={toggleMobileMenu}>
          <div className={`hamburger ${mobileMenuOpen ? "active" : ""}`}>
            <span></span>
            <span></span>
            <span></span>
          </div>
        </button>
      </div>

      {/* Mobile Navigation */}
      <div className={`mobile-nav ${mobileMenuOpen ? "open" : ""}`}>
        <ul className="mobile-nav-links">
          <li>
            <Link to="/" onClick={() => setMobileMenuOpen(false)}>
              Home
            </Link>
          </li>
          <li>
            <Link to="/about" onClick={() => setMobileMenuOpen(false)}>
              About
            </Link>
          </li>
          <li>
            <Link to="/services" onClick={() => setMobileMenuOpen(false)}>
              Services
            </Link>
          </li>
          <li>
            <Link to="/contact" onClick={() => setMobileMenuOpen(false)}>
              Contact
            </Link>
          </li>
          {isLoggedIn ? (
            <>
              <li>
                <button onClick={() => { handleDashboard(); setMobileMenuOpen(false); }}>
                  Dashboard
                </button>
              </li>
              <li>
                <button onClick={() => { handleLogout(); setMobileMenuOpen(false); }} className="mobile-logout-btn">
                  Logout
                </button>
              </li>
            </>
          ) : (
            <>
              <li>
                <button onClick={openLoginModal} className="mobile-login-btn">
                  Login
                </button>
              </li>
              <li>
                <button onClick={openSignUpModal} className="mobile-signup-btn">
                  Sign Up
                </button>
              </li>
            </>
          )}
        </ul>
      </div>

      {/* Modals */}
      {showLogin && <LoginModal onClose={() => setShowLogin(false)} onLogin={() => setIsLoggedIn(true)} />}
      {showSignUp && <SignUpModal onClose={() => setShowSignUp(false)} onRegisterSuccess={openLoginModal} />}
    </header>
  );
}

export default Header;