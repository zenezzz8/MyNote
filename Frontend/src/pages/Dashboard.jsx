import React, { useEffect, useState } from 'react';
import API from '../api/axios';
import { useNavigate } from 'react-router-dom';

export default function Notes() {
    const [notes, setNotes] = useState([]);
    const [type, setType] = useState('');
    const [body, setBody] = useState('');
    const navigate = useNavigate();

    const fetchNotes = async () => {
        try {
        const response = await API.get('/api/notes/');
        setNotes(response.data.notes || []);
        } catch (err) {
        console.error('Gagal memuat catatan', err);
        if (err.response?.status === 401) {
            navigate('/login');
        }
        }
    };

    useEffect(() => {
        fetchNotes();
    }, []);

    const handleLogout = async () => {
        try {
            await API.post('/api/logout');
        } catch (err) {
            console.error('Gagal melakukan logout:', err);
        } finally {
            localStorage.removeItem('token');
            navigate('/login');
        }
    };

    return (
        <div className="min-h-screen bg-gray-50 p-6">
        <div className="max-w-4xl mx-auto">
            <div className="flex justify-between items-center mb-8">
                <h1 className="text-3xl font-extrabold text-gray-800">Manajemen Catatan</h1>

                <div className='flex gap-3'>
                    <button
                        onClick={() => navigate('/create-notes')}
                        className="px-4 py-2 text-sm font-medium text-white bg-green-600 rounded-lg hover:bg-green-700 transition shadow">
                        Create
                    </button>

                    <button
                        onClick={handleLogout}
                        className="px-4 py-2 text-sm font-medium text-white bg-red-600 rounded-lg hover:bg-red-700 transition shadow">
                        Logout
                    </button>
                </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                {notes.length > 0 ? (
                    notes.map((note) => (
                    <div onClick={() => navigate(`/edit-notes/${ note.ID}`)} key={note.ID} className="bg-white p-5 rounded-xl shadow-md border border-gray-100 flex flex-col justify-between">
                        <div>
                        <h3 className="text-lg font-bold text-gray-800 mb-2">{note.type}</h3>
                        <p className="text-gray-600 whitespace-pre-line">{note.body}</p>
                        </div>
                    </div>
                    ))
                ) : (
                    <p className="text-gray-500 col-span-2 text-center py-6">Belum ada catatan tersedia.</p>
                )}
            </div>
        </div>
        </div>
    );
}