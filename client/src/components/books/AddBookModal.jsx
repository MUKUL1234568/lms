"use client"

import { useState } from "react"
import Modal from "./Modal"
import "./AddBookModal.css"

const AddBookModal = ({ onClose, onAdd }) => {
  const [newBook, setNewBook] = useState({ title: "", author: "", year: "" })

  const handleSubmit = (e) => {
    e.preventDefault()
    onAdd(newBook)
  }

  return (
    <Modal onClose={onClose}>
      <h2>Add New Book</h2>
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
          placeholder="Author"
          value={newBook.author}
          onChange={(e) => setNewBook({ ...newBook, author: e.target.value })}
          required
        />
        <input
          type="number"
          placeholder="Year"
          value={newBook.year}
          onChange={(e) => setNewBook({ ...newBook, year: e.target.value })}
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

