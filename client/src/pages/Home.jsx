import { useState, useEffect } from "react";
import Header from "../components/layout/Header";
import Footer from "../components/layout/Footer";
import "./Home.css"; // Import Home-specific styles
import axios from "axios";
 
function Home() {
  const [currentImage, setCurrentImage] = useState(0);
   
  const [activeTab, setActiveTab] = useState("catalog");
  const [isFormOpen, setIsFormOpen] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [formData, setFormData] = useState({
    libraryName: "",
    userName: "",
    email: "",
    contact: "",
    password: ""
  });
  const [formErrors, setFormErrors] = useState({});
  const [animatedStats, setAnimatedStats] = useState({
    libraries: 0,
    users: 0,
    books: 0,
    countries: 0
  });
  const [statsInView, setStatsInView] = useState(false);

  // Images for hero section carousel
  const images = [
    "/assets/library1.jpg",
    "/assets/library2.jpg",
    "/assets/library3.jpg"
  ];

  // Features data
  const features = [
    {
      id: "catalog",
      title: "Smart Cataloging",
      description: "Organize your books with intelligent categorization and tagging systems.",
      icon: "📚",
      image:  "/assets/Library.jpeg",
      story: "The LibraEase Public Library reduced cataloging time by 60% after implementing our smart system."
    },
    {
      id: "discover",
      title: "Powerful Discovery",
      description: "Help readers find exactly what they're looking for with advanced search capabilities.",
      icon: "🔍",
      image: "/assets/discovery.png",
      story: "University of Michigan students reported finding research materials 3x faster with our discovery tools."
    },
    {
      id: "share",
      title: "Community Sharing",
      description: "Build a community around your library with sharing and recommendation features.",
      icon: "🔄",
      image: "/assets/com.jpg",
      story: "A small town library increased membership by 45% after implementing our community features."
    },
    {
      id: "analytics",
      title: "Insightful Analytics",
      description: "Gain valuable insights into reading habits and collection usage.",
      icon: "📊",
      image: "/assets/ha.png",
      story: "Boston College optimized their acquisition budget by 30% using our analytics dashboard."
    },
   
    {
      id: "security",
      title: "Advanced Security",
      description: "Keep your library data safe with enterprise-grade security features.",
      icon: "🔒",
      image: "/assets/security.JPG",
      story: "Harvard University chose our platform for its robust security features, protecting millions of records."
    }
  ];

  const [statsData, setStatsData] = useState({
    total_libraries: 0,
    total_users: 0,
    total_books: 0,
  });

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const response = await axios.get("http://localhost:8080/libraries/states");
        const data = response.data.states;

        setStatsData({
          total_libraries: data.total_libraries,
          total_users: data.total_users,
          total_books: data.total_books,
        });
      } catch (error) {
        console.error("Error fetching stats:", error);
      }
    };

    fetchStats();
  }, []);

  const stats = [
    {
      icon: "📚",
      value: statsData.total_libraries,
      label: "Libraries Created",
      suffix: "+"
    },
    {
      icon: "👥",
      value: statsData.total_users,
      label: "Active Users",
      suffix: "+"
    },
    {
      icon: "📖",
      value: statsData.total_books,
      label: "Books Cataloged",
      suffix: "+"
    },
    {
      icon: "🌎",
      value: 1,  // Keeping this static as it's not provided by the API
      label: "Countries",
      suffix: "+"
    }
  ];

  // Image carousel effect
  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentImage((prev) => (prev + 1) % images.length);
    }, 5000);
    return () => clearInterval(interval);
  }, [images.length]);

  // Testimonial carousel effect
  
  // Stats animation when in view
   // Replace the existing useEffect for stats animation with this simpler version
useEffect(() => {
  // Function to check if an element is in viewport
  const isInViewport = (element) => {
    const rect = element.getBoundingClientRect();
    return (
      rect.top <= (window.innerHeight || document.documentElement.clientHeight) &&
      rect.bottom >= 0
    );
  };

  // Function to animate stats
  const animateStats = () => {
    const statsSection = document.getElementById('stats-section');
    
    if (statsSection && isInViewport(statsSection) && !statsInView) {
      setStatsInView(true);
      
      // Start animation for each stat
      let startTime = null;
      const duration = 1000; // 5 seconds
      
      const animate = (timestamp) => {
        if (!startTime) startTime = timestamp;
        const progress = Math.min((timestamp - startTime) / duration, 1);
        
        setAnimatedStats({
          libraries: Math.floor(progress * stats[0].value),
          users: Math.floor(progress * stats[1].value),
          books: Math.floor(progress * stats[2].value),
          countries: Math.floor(progress * stats[3].value)
        });
        
        if (progress < 1) {
          requestAnimationFrame(animate);
        }
      };
      
      requestAnimationFrame(animate);
    }
  };
  
  // Add scroll event listener
  window.addEventListener('scroll', animateStats);
  // Check once on mount in case stats are already in view
  animateStats();
  
  return () => window.removeEventListener('scroll', animateStats);
}, [statsInView, stats]);

// Remove the second useEffect for stats animation since we've combined them
  // Form handling
  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value
    });
  };

  const validateForm = () => {
    const errors = {};
    
    if (!formData.libraryName || formData.libraryName.length < 3) {
      errors.libraryName = "Library name must be at least 3 characters";
    }
    
    if (!formData.userName || formData.userName.length < 2) {
      errors.userName = "Name must be at least 2 characters";
    } else if (!/^[A-Za-z\s]+$/.test(formData.userName)) {
      errors.userName = "Name must contain only letters and spaces";
    }
    
    
    if (!formData.email || !/\S+@\S+\.\S+/.test(formData.email)) {
      errors.email = "Please enter a valid email address";
    }
    
    // if (!formData.password || formData.password.length < 8 || !/(?=.*[a-z])(?=.*[A-Z])(?=.*\d)/.test(formData.password)) {
    //   errors.password = "Password must be at least 8 characters with uppercase, lowercase, and numbers";
    // }
    
    if (!formData.contact || formData.contact.length < 10) {
      errors.contact = "Please enter a valid phone number";
    }
    
    setFormErrors(errors);
    return Object.keys(errors).length === 0;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
  
    if (validateForm()) {
      setIsSubmitting(true);
  
      const libraryData = {
        owner_email: formData.email,
        owner_name: formData.userName,
        owner_password: formData.password,
        owner_contact: formData.contact,
        library_name: formData.libraryName
      };
  
      try {
        const response = await axios.post("http://localhost:8080/libraries/", libraryData, {
          headers: {
            "Content-Type": "application/json"
          }
        });
  
        console.log("Library Created:", response.data);
        alert(`${formData.libraryName} has been successfully created!`);
  
        // Reset form and close modal
        setIsFormOpen(false);
        setFormData({
          libraryName: "",
          userName: "",
          email: "",
          contact: "",
          password: ""
        });
  
      } catch (error) {
        // Check if error.response is available (API error response)
        const errorMessage = error.response && error.response.data
          ? error.response.data.error || "Failed to create library. Please try again."
          : "Failed to create library. Please try again.";
          
        console.error("Error creating library:", error);
        alert(errorMessage);  // Display the error message from the API
      } finally {
        setIsSubmitting(false);
      }
    }
};


  // Format large numbers with commas
  const formatNumber = (num) => {
    return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
  };

  return (
    <>
      <Header />
      <div className="home-container">
        {/* Hero Section */}
        <section className="hero-section">
          <div className="hero-content">
            <h1>Your Gateway to <span className="highlight">Knowledge</span></h1>
            <p>Create and manage your own digital library with powerful tools for cataloging, sharing, and discovering books.</p>
            <div className="hero-buttons">
              <button className="primary-btn" onClick={() => setIsFormOpen(true)}>
                <span className="icon">➕</span> Create Library
              </button>
              <button className="secondary-btn">
                Learn more <span className="icon">→</span>
              </button>
            </div>
            <div className="hero-stats">
              <div className="stat-item">
                <span className="icon">📚</span>
                <span>10,000+ Libraries</span>
              </div>
              <div className="stat-item">
                <span className="icon">👥</span>
                <span>1M+ Users</span>
              </div>
              <div className="stat-item">
                <span className="icon">📖</span>
                <span>50M+ Books</span>
              </div>
            </div>
          </div>
          <div className="hero-image-container">
            {images.map((src, index) => (
              <div
                key={index}
                className={`hero-image ${currentImage === index ? "active" : ""}`}
                style={{ backgroundImage: `url(${src || "/assets/libraryBooks.jpg"})` }}
              >
                <div className="image-overlay"></div>
              </div>
            ))}
          </div>
        </section>

        {/* Features Section */}
        <section className="features-section" id="features">
          <div className="section-header">
            <h2>Powerful Features for Modern Libraries</h2>
            <p>Discover how our platform transforms the way libraries operate and engage with their communities.</p>
          </div>

          <div className="features-tabs">
            <div className="tabs-header">
              {features.map((feature) => (
                <button
                  key={feature.id}
                  className={`tab-button ${activeTab === feature.id ? "active" : ""}`}
                  onClick={() => setActiveTab(feature.id)}
                >
                  <span className="tab-icon">{feature.icon}</span>
                  <span className="tab-title">{feature.title}</span>
                </button>
              ))}
            </div>
            
            <div className="tabs-content">
              {features.map((feature) => (
                <div 
                  key={feature.id} 
                  className={`tab-content ${activeTab === feature.id ? "active" : ""}`}
                >
                  <div className="feature-details">
                    <h3>{feature.title}</h3>
                    <p>{feature.description}</p>
                    
                    <div className="feature-story">
                      <h4>Success Story</h4>
                      <p>{feature.story}</p>
                    </div>
                  </div>
                  <div className="feature-image">
                    <img 
                      src={feature.image ||  "../assets/Library.jpeg"} 
                      alt={feature.title} 
                    />
                  </div>
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* Call to Action */}
        <section className="cta-section">
          <h2>Ready to Start Your Library Journey?</h2>
          <p>Join thousands of libraries worldwide and create your own digital space for knowledge sharing.</p>
          <button className="primary-btn" onClick={() => setIsFormOpen(true)}>
            <span className="icon">➕</span> Create Library
          </button>
        </section>

        {/* Stats Section */}
        <section className="stats-section" id="stats-section">
          <div className="section-header">
            <h2>Our Global Impact</h2>
            <p>Transforming libraries and communities worldwide</p>
          </div>
          
          <div className="stats-container">
            {stats.map((stat, index) => (
              <div key={index} className="stat-card">
                <div className="stat-icon">{stat.icon}</div>
                <div className="stat-value">
                  {formatNumber(
                    index === 0 ? animatedStats.libraries :
                    index === 1 ? animatedStats.users :
                    index === 2 ? animatedStats.books :
                    animatedStats.countries
                  )}
                  {stat.suffix}
                </div>
                <div className="stat-label">{stat.label}</div>
              </div>
            ))}
          </div>
        </section>

        

        {/* Create Library Form Modal */}
        {isFormOpen && (
          <div className="modal-overlay">
            <div className="modal-container">
              <div className="modal-header">
                <h2>Create Your Library</h2>
                <button className="close-btn" onClick={() => setIsFormOpen(false)}>×</button>
              </div>
              
              <div className="modal-body">
                <form onSubmit={handleSubmit}>
                  <div className="form-group">
                    <label htmlFor="libraryName">Library Name</label>
                    <input
                      type="text"
                      id="libraryName"
                      name="libraryName"
                      value={formData.libraryName}
                      onChange={handleInputChange}
                      placeholder="My Amazing Library"
                    />
                    {formErrors.libraryName && <div className="error-message">{formErrors.libraryName}</div>}
                    <div className="form-hint">This is how your library will appear to others.</div>
                  </div>
                  
                  <div className="form-group">
                    <label htmlFor="userName">Your Name</label>
                    <input
                      type="text"
                      id="userName"
                      name="userName"
                      value={formData.userName}
                      onChange={handleInputChange}
                      placeholder="John Doe"
                    />
                    {formErrors.userName && <div className="error-message">{formErrors.userName}</div>}
                  </div>
                  
                  <div className="form-row">
                    <div className="form-group">
                      <label htmlFor="email">Email</label>
                      <input
                        type="email"
                        id="email"
                        name="email"
                        value={formData.email}
                        onChange={handleInputChange}
                        placeholder="you@example.com"
                      />
                      {formErrors.email && <div className="error-message">{formErrors.email}</div>}
                    </div>
                    
                    <div className="form-group">
                      <label htmlFor="contact">Contact Number</label>
                      <input
                        type="tel"
                        id="contact"
                        name="contact"
                        value={formData.contact}
                        onChange={handleInputChange}
                        placeholder="+1 (555) 123-4567"
                      />
                      {formErrors.contact && <div className="error-message">{formErrors.contact}</div>}
                    </div>
                  </div>
                  
                  <div className="form-group">
                    <label htmlFor="password">Password</label>
                    <input
                      type="password"
                      id="password"
                      name="password"
                      value={formData.password}
                      onChange={handleInputChange}
                      placeholder="••••••••"
                    />
                    {formErrors.password && <div className="error-message">{formErrors.password}</div>}
                    <div className="form-hint">Must be at least 8 characters with uppercase, lowercase, and numbers.</div>
                  </div>
                  
                  <div className="form-actions">
                    <button 
                      type="button" 
                      className="secondary-btn" 
                      onClick={() => setIsFormOpen(false)}
                      disabled={isSubmitting}
                    >
                      Cancel
                    </button>
                    <button 
                      type="submit" 
                      className="primary-btn"
                      disabled={isSubmitting}
                    >
                      {isSubmitting ? "Creating..." : "Create Library"}
                    </button>
                  </div>
                </form>
              </div>
            </div>
          </div>
        )}
      </div>
      <Footer />
    </>
  );
}

export default Home;