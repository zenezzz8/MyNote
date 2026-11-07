import { useState } from 'react';
import API from '../../api/axios';
import { useNavigate } from 'react-router-dom';

export default function CreateNotes() {
    const [type, setType] = useState('');
    const [body, setBody] = useState('');
    const navigate = useNavigate();

    const handleCreateNote = async (e) => {
        e.preventDefault();
        try {
            await API.post('/api/notes/', { type, body, check: false });
            setType('');
            setBody('');
            navigate('/notes'); 
        } catch (err) {
            console.error('Gagal membuat catatan', err);
        }
    };

    return (    
        <div className="min-h-screen bg-gray-50 p-6">
            <div className="max-w-4xl mx-auto">
                <form onSubmit={handleCreateNote} className="bg-white p-6 rounded-xl shadow-md mb-8 space-y-4 border border-gray-100">
                <h2 className="text-xl font-bold text-gray-700">Tambah Catatan Baru</h2>
                <div>
                    <input
                    type="text"
                    placeholder="Judul Catatan"
                    value={type}
                    onChange={(e) => setType(e.target.value)}
                    required
                    className="w-full px-4 py-2 border rounded-lg focus:ring-blue-500 focus:border-blue-500 border-gray-300 outline-none"
                    />
                </div>
                <div>
                    <textarea
                    placeholder="Tulis isi catatan di sini..."
                    value={body}
                    onChange={(e) => setBody(e.target.value)}
                    required
                    rows="3"
                    className="w-full px-4 py-2 border rounded-lg focus:ring-blue-500 focus:border-blue-500 border-gray-300 outline-none"
                    ></textarea>
                </div>
                <button
                    type="submit"
                    className="px-6 py-2 bg-green-600 text-white font-semibold rounded-lg hover:bg-green-700 transition shadow">
                    Simpan Catatan
                </button>
                </form>
            </div>
        </div>
    );
} 