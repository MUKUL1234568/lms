import { useState, useEffect } from "react";
import axios from "axios";
import { Pie } from "react-chartjs-2";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import "./Dashboard.css";

ChartJS.register(ArcElement, Tooltip, Legend);

const Dashboard = () => {
  const [stats, setStats] = useState(null);
  const [loading, setLoading] = useState(true);

  // Retrieve the token from localStorage
  const token = localStorage.getItem("token");

  // Assuming libId is also stored or passed as needed
    // Replace with dynamic value if necessary

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/libraries/states/id`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        // Extracting states from the response
        const { states } = response.data;

        setStats(states);
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) {
    return <div>Loading...</div>; // Show loading message while fetching data
  }

  // Data for the Pie chart
  const bookData = {
    labels: ["Total Books", "Currently Issued Books"],
    datasets: [
      {
        data: [stats.total_books, stats.total_issued_book],
        backgroundColor: ["#36A2EB", "#FFCE56"],
        hoverBackgroundColor: ["#36A2EB", "#FFCE56"],
      },
    ],
  };

  return (
    <div className="dashboard">
      <h2>Dashboard</h2>
      <div className="dashboard-stats">
        <div className="stat-item">
          <h3>Total Books</h3>
          <p>{stats.total_books}</p>
        </div>
        <div className="stat-item">
          <h3>Total Users</h3>
          <p>{stats.total_users}</p>
        </div>
        <div className="stat-item">
          <h3>Total Issued Books</h3>
          <p>{stats.total_issued_book}</p>
        </div>
      </div>
      <div className="dashboard-charts">
        <div className="chart-item">
          <h3>Books Overview</h3>
          <Pie data={bookData} />
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
