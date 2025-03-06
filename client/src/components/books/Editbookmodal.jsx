 "use client"

import { useState, useEffect } from "react"
import Modal from "./Modal"
import "./EditBookModal.css"

const EditBookModal = ({ book, onClose, onUpdate }) => {
  const [editedBook, setEditedBook] = useState(book)
  const [error, setError] = useState("")
    // console.log(book)
  useEffect(() => {
    setEditedBook(book)
  }, [book])

  const validateForm = () => {
    if (!/^\d{13}$/.test(editedBook.isbn)) {
      setError("ISBN must be exactly 13 digits.")
      return false
    }
  
    if (Number(editedBook.available_copies) > Number(editedBook.total_copies)) {
      setError("Available copies cannot be greater than total copies.")
      return false
    }
  
    const issuedBooksCount = editedBook.issue_records
      ? editedBook.issue_records.filter(record => record.status.toLowerCase() === "issued").length
      : 0
         console.log(issuedBooksCount)
         console.log(Number(editedBook.total_copies))
    if (issuedBooksCount > Number(editedBook.total_copies)) {
      setError(`Cannot reduce total copies below ${issuedBooksCount}. ${issuedBooksCount} book(s) are currently issued to users.`)
      return false
    }
  
    setError("")
    return true
  }
  

  const handleSubmit = (e) => {
    e.preventDefault()
    if (validateForm()) {
      onUpdate(editedBook)
    }
  }

  return (
    <Modal onClose={onClose}>
      <h2>Edit Book</h2>
      {error && <p className="error-message">{error}</p>}
      <form onSubmit={handleSubmit} className="edit-book-form">
        <input
          type="text"
          placeholder="Title"
          value={editedBook.title}
          onChange={(e) => setEditedBook({ ...editedBook, title: e.target.value })}
          required
        />
        <input
          type="text"
          placeholder="ISBN (13 digits)"
          value={editedBook.isbn}
          onChange={(e) => setEditedBook({ ...editedBook, isbn: e.target.value })}
          required
        />
        <input
          type="text"
          placeholder="Author"
          value={editedBook.authors}
          onChange={(e) => setEditedBook({ ...editedBook, authors: e.target.value })}
          required
        />
        <input
          type="text"
          placeholder="Publisher"
          value={editedBook.publisher}
          onChange={(e) => setEditedBook({ ...editedBook, publisher: e.target.value })}
          required
        />
        <input
          type="text"
          placeholder="Version"
          value={editedBook.version}
          onChange={(e) => setEditedBook({ ...editedBook, version: e.target.value })}
          required
        />
        <input
          type="number"
          placeholder="Total Copies"
          value={editedBook.total_copies}
          onChange={(e) => setEditedBook({ ...editedBook, total_copies: e.target.value })}
          required
        />
        <input
          type="number"
          placeholder="Available Copies"
          value={editedBook.available_copies}
          onChange={(e) => setEditedBook({ ...editedBook, available_copies: e.target.value })}
          required
        />
        <button type="submit" className="submit-btn">
          Update Book
        </button>
      </form>
    </Modal>
  )
}

export default EditBookModal
