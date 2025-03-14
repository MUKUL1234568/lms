"use client"

import { useState, version } from "react"
import Modal from "./Modal"
import "./AddBookModal.css"

const AddBookModal = ({ onClose, onAdd }) => {
  const [newBook, setNewBook] = useState({ isbn:"",title: "", authors: "",publisher:"",version:"" ,total_copies: ""})
  const [error, setError] = useState("")
  const validateForm = () => {
    if (!/^\d{13}$/.test(newBook.isbn)) {
      setError("ISBN must be exactly 13 digits.")
      return false
    }
    
    if (!/^[A-Za-z\s,]+$/.test(newBook.authors)) {
      setError("Author name must contain only letters, spaces, and commas.")
      return false
    }
  
    setError("")
    return true
  }
  

  const handleSubmit = (e) => {
    e.preventDefault()
    if(validateForm()){
    onAdd(newBook)
    }
  }

  return (
    <Modal onClose={onClose}>
      <h2>Add New Book</h2>
      {error && <p className="error-message">{error}</p>}
      <form onSubmit={handleSubmit} className="add-book-form">
        <input
          type="text"
          placeholder="Title"
          value={newBook.title}
          onChange={(e) => setNewBook({ ...newBook, title: e.target.value })}
          required
        />
        <input
          type="text"
          placeholder="ISBN"
          value={newBook.isbn}
          onChange={(e) => setNewBook({ ...newBook,isbn: e.target.value })}
          required
        />
        <input
          type="text"
          placeholder="Publisher"
          value={newBook.publisher}
          onChange={(e) => setNewBook({ ...newBook, publisher: e.target.value })}
          required
        />
        <input
          type="text"
          placeholder="Version"
          value={newBook.version}
          onChange={(e) => setNewBook({ ...newBook, version: e.target.value })}
          required
        />
        <input
          type="text"
          placeholder="Author"
          value={newBook.authors}
          onChange={(e) => setNewBook({ ...newBook, authors: e.target.value })}
          required
        />
        <input
          type="number"
          placeholder="Total Copies"
          value={newBook.total_copies}
          onChange={(e) => setNewBook({ ...newBook, total_copies: e.target.value })}
          required
        />
        <button type="submit" className="submit-btn">
          Add Book
        </button>
      </form>
    </Modal>
  )
}

export default AddBookModal

