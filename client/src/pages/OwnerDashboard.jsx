

import { useEffect, useState } from "react"
import Header from "../components/layout/Header"
import Sidebar from "../components/admin/OwnerSidebar"
import BookList from "../components/books/BookList"
import axios from "axios"
import Dashboard from "../components/admin/Dashboard"
 import RequestList from "../components/admin/RequestList"  
 import UserList from "../components/admin/UserList"
 import IssuedBookList from "../components/books/IssuedBookList" 
import AddBookModal from "../components/books/AddBookModal"
 import EditBookModal from "../components/books/Editbookmodal"
import "./OwnerDashboard.css"
import AdminManagement from "../components/admin/Adminmanagement"
import ALLRequestList from "../components/admin/AllRequestList"

const admin = () => {
  const [activeTab, setActiveTab] = useState("dashboard")
  const [showAddBookModal, setShowAddBookModal] = useState(false)
  const [showEditBookModal, setShowEditBookModal] = useState(false)
  const [editingBook, setEditingBook] = useState(null)
  const [books, setBooks] = useState([])
  const [users, setUsers] = useState([])
   const [requests, setRequests] = useState([])
   const [issueregistry,setissueregistry]=useState([])

   const fetchbook= async ()=>{
    const token =localStorage.getItem("token");
  
    try{
      const response =await axios.get("http://localhost:8080/book/",{
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

 const fetchIssueregistry=async ()=>{
  const token =localStorage.getItem("token");
  
  try{
    const response =await axios.get("http://localhost:8080/issueregistry/",{
      headers:{
          Authorization:` Bearer ${token}`,
      },
    });
   
    if(response.status==200)
    {
      setissueregistry(response.data.issued_books)
    }
    else{
      console.log("error in fetching the issueregistry from database")
    }
  } catch(error){
      console.error("error in fetching the ",error)
  }
 };

  
  const fetchuser= async ()=>{
    const token =localStorage.getItem("token");
  
    try{
      const response =await axios.get("http://localhost:8080/user/",{
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
      const response =await axios.get("http://localhost:8080/request/",{
        headers:{
            Authorization:` Bearer ${token}`,
        },
      });
     
      if(response.status==200)
      {
        // console.log(response.data.requests)
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
    fetchIssueregistry();
},[]);

   

 
   

 
 

const addBook = async (newBook) => {
  const token = localStorage.getItem("token");

  try {
    const response = await axios.post(
      "http://localhost:8080/book/",
      { ...newBook, total_copies: Number(newBook.total_copies) },
      {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      }
    );

    if (response.status === 201) {
      // New book added, so just append it to the state
      setBooks([...books, { ...newBook, available_copies: newBook.total_copies }]);
      alert("Book added successfully!");
    } else if (response.status === 200) {
      // Book already exists, update its total and available copies
      setBooks((prevBooks) =>
        prevBooks.map((book) =>
          book.isbn === newBook.isbn
            ? {
                ...book,
                total_copies: book.total_copies + Number(newBook.total_copies),
                available_copies: book.available_copies + Number(newBook.total_copies),
              }
            : book
        )
      );
      alert("Book already exists! Updated total and available copies.");
    } else {
      console.error("Failed to add book");
    }
  } catch (error) {
    console.error("Error adding the book", error);
    alert("Failed to add the book. Please try again.");
  }

  setShowAddBookModal(false);
};



  const updateBook = async (updatedBook) => {
    // console.log(updatedBook)
    const token =localStorage.getItem("token");
        console.log(updatedBook.isbn)
    try{
       const response=await axios.patch(`http://localhost:8080/book/${updatedBook.isbn}`,{...updatedBook,total_copies: Number(updatedBook.total_copies),available_copies: Number(updatedBook.available_copies)},
       {
        headers:{"Content-Type":"application/json",
          Authorization:`Bearer ${token}`}
       }); 
       if(response.status==200){
        setBooks(books.map((book) => (book.isbn === updatedBook.isbn ? updatedBook : book)))
        setShowEditBookModal(false)
       } 
       else
       {
        console.error("faild to update the book")
       }
    } catch(error){
      console.error("error in updating the book",error)
}
         setShowAddBookModal(false)
  }

  const removeBook = async(isbn) => {
    const token =localStorage.getItem("token");
    
try{
   const response=await axios.delete(`http://localhost:8080/book/${isbn}`,
   {
    headers:{"Content-Type":"application/json",
      Authorization:`Bearer ${token}`}
   }); 
   if(response.status==200){
    setBooks(books.filter((book) => book.isbn !== isbn))
   } 
   else
   {
    console.error("faild to delete the book")
   }
} catch(error){
  console.error("error in deleting the book",error)
}
    
  }


  

  const approveRequest = async (id) => {
    const token = localStorage.getItem("token");
    
    try {
      // Call the approve API to approve the request with the token
      const response = await axios.put(
        `http://localhost:8080/request/${id}`,
        { approve: true }, // Pass the approval flag
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`, // Pass the token in the header
          },
        }
      );
  
      // After successful approval, update the local state
      if (response.status === 200) {
        setRequests(
          requests.map((r) =>
            r.req_id === id ? { ...r, status: "approved" } : r
          )
        );
      } else {
        console.error("Failed to approve the request");
      }
    } catch (error) {
      console.error("Error approving the request", error);
    }
  };
  

  const rejectRequest = async(id) => {
    const token = localStorage.getItem("token");
    
    try {
      // Call the approve API to approve the request with the token
      const response = await axios.put(
        `http://localhost:8080/request/${id}`,
        { approve: false }, // Pass the approval flag
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`, // Pass the token in the header
          },
        }
      );
  
      // After successful approval, update the local state
      if (response.status === 200) {
        setRequests(
          requests.map((r) =>
            r.req_id === id ? { ...r, status: "approved" } : r
          )
        );
      } else {
        console.error("Failed to approve the request");
      }
    } catch (error) {
      console.error("Error approving the request", error);
    }
  };
  // const updateIssuedBook = (id, status) => {
  //   setIssuedBooks(issuedBooks.map((book) => (book.id === id ? { ...book, status } : book)))
  // }

  const updateUserRole = async (userId, newRole) => {
    const token = localStorage.getItem("token")
    try {
      const response = await axios.patch(
        `http://localhost:8080/user/${userId}`,
        { role: newRole },
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        },
      )
      if (response.status === 200) {
        setUsers(users.map((user) => (user.id === userId ? { ...user, role: newRole } : user)))
      } else {
        console.error("Failed to update user role")
      }
    } catch (error) {
      console.error("Error updating user role", error)
    }
  }

  return (
    <><Header />
    <div className="app">
      <div className="main-content">
        <Sidebar activeTab={activeTab} setActiveTab={setActiveTab} />
        <div className="content">
        {activeTab === "dashboard" && <Dashboard />}
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
                onEdit={(book) => {
                  setEditingBook(book)
                  setShowEditBookModal(true)
                }}
                onRemove={removeBook}
                books={books}
              />
            </>
          )}
           {activeTab === "requests" && (
            <>
              {/* <h2>Pending Requests</h2> */}
              <RequestList requests={requests}  onApprove={approveRequest} onReject={rejectRequest}
              />
            </>
          )}
          {activeTab === "users" && (
            <>
              <h2>User List</h2>
              <UserList users={users}  />
            </>
          )}
          {activeTab === "issued" && (
            <>
               
              <IssuedBookList  issueregistry={issueregistry}  />
            </>
          )}
            {activeTab === "allrequests"  && <ALLRequestList requests={requests}  />}
           {activeTab === "adminmanagement" &&(
            <>
             
             <AdminManagement users={users} onUpdateUserRole={updateUserRole} />
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

