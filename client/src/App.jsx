import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/Home.jsx";
import Admin from "./pages/LibraryAdminDashboard.jsx";
// import UserDashboard from "./pages/UserDashboard";
// import NotFound from "./pages/NotFound";
import LoginForm from "./components/auth/LoginModal.jsx";
// import RegisterForm from "./components/auth/RegisterForm";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/admindashboard" element={<Admin/>} />
        {/* <Route path="/user" element={<UserDashboard />} /> */}
       
      </Routes>
    </Router>
  );
}

export default App;
