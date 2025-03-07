import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/Home.jsx";
import Admin from "./pages/LibraryAdminDashboard.jsx";
import Owner from "./pages/OwnerDashboard.jsx"
import UserDashboard from "./pages/UserDashboard.jsx";
 

 

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/admindashboard" element={<Admin/>} />
        <Route path="ownerdashboard" element={<Owner/>}/>
        <Route path="/readerdashboard" element={<UserDashboard />} />
       
      </Routes>
    </Router>
  );
}

export default App;
