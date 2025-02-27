

import { useEffect, useState } from "react"
import Header from "../components/layout/Header"
import Sidebar from "../components/admin/Sidebar"
import BookList from "../components/books/BookList"
import axios from "axios"
import Dashboard from "../components/admin/Dashboard"
 import RequestList from "../components/admin/RequestList"  
 import UserList from "../components/admin/UserList"
 import IssuedBookList from "../components/books/IssuedBookList" 
import AddBookModal from "../components/books/AddBookModal"
// import EditBookModal from "./components/EditBookModal"
import "./admin.css"

const admin = () => {
  const [activeTab, setActiveTab] = useState("dashboard")
  const [showAddBookModal, setShowAddBookModal] = useState(false)
  const [showEditBookModal, setShowEditBookModal] = useState(false)
  const [editingBook, setEditingBook] = useState(null)
  const [books, setBooks] = useState([])
  const [users, setUsers] = useState([])
   const [requests, setRequests] = useState([])
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
  const fetchuser= async ()=>{
    const token =localStorage.getItem("token");
  
    try{
      const response =await axios.get("http://localhost:8080/users/",{
        headers:{
            Authorization:` Bearer ${token}`,
        },
      });
     
      if(response.status==200)
      {
        setUsers(response.data.users)
      }
      else{
        console.log("error in fetching the users from database")
      }
    } catch(error){
        console.error("error in fetching the user ",error)
    }
  };
 
  const fetchrequest= async ()=>{
    const token =localStorage.getItem("token");
  
    try{
      const response =await axios.get("http://localhost:8080/requests/allreq",{
        headers:{
            Authorization:` Bearer ${token}`,
        },
      });
     
      if(response.status==200)
      {
        console.log(response.data.requests)
        setRequests(response.data.requests)
      }
      else{
        console.log("error in fetching the req from database")
      }
    } catch(error){
        console.error("error in fetching the req ",error)
    }
  };
  useEffect(()=>{
    fetchbook();
    fetchuser();
    fetchrequest();
},[]);

   

 
  const [issuedBooks, setIssuedBooks] = useState([
    { id: 1, bookId: "977-0234190440", userId: 1, issueDate: "2023-05-01", returnDate: "2023-05-15" },
  ])

 
 

  const addBook =  async (newBook) => {
    
    const token=localStorage.getItem("token")
    try{
      const response= await axios.post("http://localhost:8080/books/",{...newBook, total_copies: Number(newBook.total_copies)} ,{
        headers:{ "Content-Type":"application/json",
          Authorization:`Bearer ${token}`
                  
        }
        
      });
         if(response.status==201){
          setBooks([...books, { ...newBook,available_copies:newBook.total_copies }])
          console.log("book added")
         }
         else{
          console.error("faild to add book ")
         }

    } catch(error){
            console.error("erroe in adding the book",error)
    }
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
    const request = requests.find((r) => r.req_id === id)
    setRequests(requests.map((r) => (r.req_id === id ? { ...r, status: "approved" } : r)))
    setIssuedBooks([
      ...issuedBooks,
      {
        id: issuedBooks.length + 1,
        bookId: request.book_id,
        userId: request.reader_id,
        issueDate: new Date().toISOString().split("T")[0],
        returnDate: new Date(Date.now() + 14 * 24 * 60 * 60 * 1000).toISOString().split("T")[0],
      },
    ])
  }

  const rejectRequest = (id) => {
    setRequests(requests.map((r) => (r.req_id === id ? { ...r, status: "rejected" } : r)))
  }
  // const updateIssuedBook = (id, status) => {
  //   setIssuedBooks(issuedBooks.map((book) => (book.id === id ? { ...book, status } : book)))
  // }

  return (
    <><Header />
    <div className="app">
      <div className="main-content">
        <Sidebar activeTab={activeTab} setActiveTab={setActiveTab} />
        <div className="content">
        {activeTab === "dashboard" && <Dashboard books={books} issuedBooks={issuedBooks} users={users} />}
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
           {activeTab === "requests" && (
            <>
              <h2>Pending Requests</h2>
              <RequestList requests={requests} books={books} users={users} onApprove={approveRequest} onReject={rejectRequest}
              />
            </>
          )}
          {activeTab === "users" && (
            <>
              <h2>User List</h2>
              <UserList users={users} issuedBooks={issuedBooks} books={books}/>
            </>
          )}
          {activeTab === "issued" && (
            <>
              <h2>Issued Books</h2>
              <IssuedBookList issuedBooks={issuedBooks} books={books} users={users}   />
            </>
          )}
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

