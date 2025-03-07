import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./LoginModal.css";
import { jwtDecode } from 'jwt-decode'

function LoginModal({ onClose, onLogin }) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    setError("");

    try {
      const response = await fetch("http://localhost:8080/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || "Login failed");
      }

      localStorage.setItem("token", data.token);
      const decodetoken =jwtDecode(data.token)
      console.log(decodetoken)
      const userRole=decodetoken.role
      if (userRole === "LibraryAdmin") {
        navigate("/admindashboard");
      } else if (userRole === "Owner") {
        navigate("/ownerdashboard");
      } else if (userRole === "Reader") {
        navigate("/readerdashboard");
      } else {
        navigate("/"); // Default route if role is undefined
      }

      onLogin(); // ✅ Notify Header to update login state
     
      onClose(); // ✅ Close modal
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h2>🔐 Login</h2>
        {error && <p className="error">{error}</p>}
        <form onSubmit={handleLogin}>
          <input
            type="email"
            placeholder="Enter your email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
          <input
            type="password"
            placeholder="Enter your password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <button type="submit">Login</button>
        </form>
        <button className="close-btn" onClick={onClose}>Close</button>
      </div>
    </div>
  );
}

export default LoginModal;
