

import { useEffect, useState } from "react"
import Header from "../components/layout/Header"
import Sidebar from "../components/admin/Sidebar"
import BookList from "../components/books/BookList"
import axios from "axios"
// import RequestList from "./components/RequestList"
// import UserList from "./components/UserList"
// import IssuedBookList from "./components/IssuedBookList"
import AddBookModal from "../components/books/AddBookModal"
// import EditBookModal from "./components/EditBookModal"
import "./admin.css"

const admin = () => {
  const [activeTab, setActiveTab] = useState("books")
  const [showAddBookModal, setShowAddBookModal] = useState(false)
  const [showEditBookModal, setShowEditBookModal] = useState(false)
  const [editingBook, setEditingBook] = useState(null)
  const [books, setBooks] = useState([])


  const fetchbook= async ()=>{
    const token =localStorage.getItem("token");
  
    try{
      const response =await axios.get("http://localhost:8080/books/lib",{
        headers:{
            Authorization:` Bearer ${token}`,
        },
      });
     
      if(response.status==200)
      {
        setBooks(response.data.books)
      }
      else{
        console.log("error in fetching the books from database")
      }
    } catch(error){
        console.error("error in fetching the ",error)
    }
  };
  useEffect(()=>{
      fetchbook();

  },[]);

  useEffect(() => {
    console.log("Books have been updated:", books);
  }, [books]);

  const [requests, setRequests] = useState([
    { id: 1, bookId: 1, userId: 1, status: "pending" },
    { id: 2, bookId: 2, userId: 2, status: "pending" },
  ])
  const [users, setUsers] = useState([
    { id: 1, name: "John Doe", email: "john@example.com" },
    { id: 2, name: "Jane Smith", email: "jane@example.com" },
  ])
  const [issuedBooks, setIssuedBooks] = useState([
    { id: 1, bookId: 1, userId: 1, issueDate: "2023-05-01", returnDate: "2023-05-15" },
  ])

  const addBook = (newBook) => {
    setBooks([...books, { id: books.length + 1, ...newBook }])
    setShowAddBookModal(false)
  }

  const updateBook = (updatedBook) => {
    setBooks(books.map((book) => (book.id === updatedBook.id ? updatedBook : book)))
    setShowEditBookModal(false)
  }

  const removeBook = (id) => {
    setBooks(books.filter((book) => book.id !== id))
  }

  const approveRequest = (id) => {
    const request = requests.find((r) => r.id === id)
    setRequests(requests.map((r) => (r.id === id ? { ...r, status: "approved" } : r)))
    setIssuedBooks([
      ...issuedBooks,
      {
        id: issuedBooks.length + 1,
        bookId: request.bookId,
        userId: request.userId,
        issueDate: new Date().toISOString().split("T")[0],
        returnDate: new Date(Date.now() + 14 * 24 * 60 * 60 * 1000).toISOString().split("T")[0],
      },
    ])
  }

  const updateIssuedBook = (id, status) => {
    setIssuedBooks(issuedBooks.map((book) => (book.id === id ? { ...book, status } : book)))
  }

  return (
    <><Header />
    <div className="app">
      <div className="main-content">
        <Sidebar activeTab={activeTab} setActiveTab={setActiveTab} />
        <div className="content">
          {activeTab === "books" && (
            <>
              <div className="content-header">
                <h2>Book List</h2>
                <div className="add-book-cantainer">
                <button className="add-book-btn" onClick={() => setShowAddBookModal(true)}>
                  Add New Book
                </button>
                </div>
              </div>
              <BookList
                books={books}
                onEdit={(book) => {
                  setEditingBook(book)
                  setShowEditBookModal(true)
                }}
                onRemove={removeBook}
              />
            </>
          )}
          {/* {activeTab === "requests" && (
            <>
              <h2>Pending Requests</h2>
              <RequestList requests={requests} books={books} users={users} onApprove={approveRequest} />
            </>
          )}
          {activeTab === "users" && (
            <>
              <h2>User List</h2>
              <UserList users={users} />
            </>
          )}
          {activeTab === "issued" && (
            <>
              <h2>Issued Books</h2>
              <IssuedBookList issuedBooks={issuedBooks} books={books} users={users} onUpdateStatus={updateIssuedBook} />
            </>
          )} */}
        </div>
      </div>
      {showAddBookModal && <AddBookModal onClose={() => setShowAddBookModal(false)} onAdd={addBook} />}
      {showEditBookModal && (<EditBookModal book={editingBook} onClose={() => setShowEditBookModal(false)} onUpdate={updateBook} />
      )}
    </div>

    </>
  )
}

export default admin

