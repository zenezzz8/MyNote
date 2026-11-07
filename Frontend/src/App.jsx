import React, { use } from 'react';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import Login from './pages/auth/Login';
import Register from './pages/auth/Register';
import Dashboard from './pages/Dashboard';
import CreateNotes from './pages/notes/createNotes';
import EditNotes from './pages/notes/editNotes';

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/login" replace />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />

        <Route path="/notes" element={<Dashboard />} />

        <Route path="/create-notes" element={<CreateNotes />} />
        <Route path="/edit-notes/:noteId" element={<EditNotes />} />
      </Routes>
    </BrowserRouter>
  );
}