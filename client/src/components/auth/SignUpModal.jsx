import { useState, useEffect } from "react";
import "./SignUpModal.css";
import axios from "axios";

function SignUpModal({ onClose, onRegisterSuccess }) {
  const [libraries, setLibraries] = useState([]); // ✅ Store fetched libraries
  const [selectedLibID, setSelectedLibID] = useState(""); // ✅ Store selected library ID

  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [contactNumber, setContactNumber] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  // ✅ Fetch all libraries on component mount
  useEffect(() => {
    const fetchLibraries = async () => {
      try {
        const response = await axios.get("http://localhost:8080/libraries/");
        setLibraries(response.data.libraries);
      } catch (error) {
        console.error("Error fetching libraries:", error);
      }
    };

    fetchLibraries();
  }, []);

  const handleSignUp = async (e) => {
    e.preventDefault();
    setError("");

    if (!selectedLibID) {
      setError("Please select a library.");
      return;
    }

    try {
      const response = await fetch("http://localhost:8080/user/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ 
          name, 
          email, 
          password, 
          contact_number: contactNumber, 
          lib_id: parseInt(selectedLibID) // ✅ Send selected library ID
        }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.error || "Registration failed");
      }

      setSuccess(true);
      setTimeout(() => {
        onClose(); // Close sign-up modal
        onRegisterSuccess(); // ✅ Open login modal
      }, 1000);
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h2>📝 Sign Up</h2>
        {error && <p className="error">{error}</p>}
        {success && <p className="success">✅ Sign-up successful! Redirecting to login...</p>}
        <form onSubmit={handleSignUp}>
          <input type="text" placeholder="Full Name" value={name} onChange={(e) => setName(e.target.value)} required />
          <input type="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} required />
          <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} required />
          <input type="text" placeholder="Contact Number" value={contactNumber} onChange={(e) => setContactNumber(e.target.value)} />

          {/* ✅ Library Selection Dropdown */}
          <select value={selectedLibID} onChange={(e) => setSelectedLibID(e.target.value)} required>
            <option value="">Select Library</option>
            {libraries.map((lib) => (
              <option key={lib.id} value={lib.id}>
                {lib.name}
              </option>
            ))}
          </select>

          <button type="submit">Sign Up</button>
        </form>
        <button className="close-btn" onClick={onClose}>Close</button>
      </div>
    </div>
  );
}

export default SignUpModal;
