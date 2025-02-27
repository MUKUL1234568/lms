import Header from "../components/layout/Header";
import Footer from "../components/layout/Footer";
import "./Home.css"; // Import Home-specific styles

function Home() {
  return (
    <>
      <Header />
      <div className="home-container">
        <div className="overlay">
          <h1>📚 Welcome to the Library Management System</h1>
         
        </div>
      </div>
      <Footer />
    </>
  );
}

export default Home;
