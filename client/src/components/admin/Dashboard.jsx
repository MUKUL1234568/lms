import { Bar, Pie } from "react-chartjs-2"
import { Chart as ChartJS, ArcElement, Tooltip, Legend, CategoryScale, LinearScale, BarElement } from "chart.js"
import "./Dashboard.css"

ChartJS.register(ArcElement, Tooltip, Legend, CategoryScale, LinearScale, BarElement)

const Dashboard = ({ books, issuedBooks, users }) => {
  // Calculate total number of copies
  const totalCopies = books.reduce((sum, book) => sum + (book.total_copies || 0), 0)

  const bookData = {
    labels: ["Total Books", "Issued Books"],
    datasets: [
      {
        data: [totalCopies, issuedBooks.length],
        backgroundColor: ["#36A2EB", "#FFCE56"],
        hoverBackgroundColor: ["#36A2EB", "#FFCE56"],
      },
    ],
  }

  return (
    <div className="dashboard">
      <h2>Dashboard</h2>
      <div className="dashboard-stats">
        <div className="stat-item">
          <h3>Total Books</h3>
          <p>{totalCopies}</p>
        </div>
        <div className="stat-item">
          <h3>Issued Books</h3>
          <p>{issuedBooks.length}</p>
        </div>
        <div className="stat-item">
          <h3>Total Users</h3>
          <p>{users.length}</p>
        </div>
      </div>
      <div className="dashboard-charts">
        <div className="chart-item">
          <h3>Books Overview</h3>
          <Pie data={bookData} />
        </div>
      </div>
    </div>
  )
}

export default Dashboard
