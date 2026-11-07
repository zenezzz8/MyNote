import { useEffect, useState } from 'react';
import API from '../../api/axios';
import { useNavigate, useParams } from 'react-router-dom';

export default function CreateNotes() {
    const [type, setType] = useState('');
    const [body, setBody] = useState('');
    const {noteId} = useParams();
    const navigate = useNavigate();

    useEffect(() => {
        if (noteId) {
            const fetchNoteDetail = async () => {
                try {
                    const response = await API.get(`/api/notes/${noteId}`);
                    const noteData = response.data.note;
                    setType(noteData.type || noteData.Type || '');
                    setBody(noteData.body || noteData.Body || '');
                } catch (err) {
                    console.error('Gagal memuat detail catatan', err);
                }
            };
            fetchNoteDetail();
        }
    }, [noteId]);

    const handleEditNote = async (e) => {
        e.preventDefault();
        try {
            await API.put(`/api/notes/${noteId}`, { type, body, check: false });
            setType('');
            setBody('');
            navigate('/notes');
        } catch (err) {
            console.error('Gagal membuat catatan', err);
        }
    };

    const handleDeleteNote = async (e) => {
        e.preventDefault();
        try {
            await API.delete(`/api/notes/${noteId}`);
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
                <form onSubmit={handleEditNote} className="bg-white p-6 rounded-xl shadow-md mb-8 space-y-4 border border-gray-100">
                <h2 className="text-xl font-bold text-gray-700">Edit Catatan</h2>
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
                <button
                    type="button"
                    onClick={handleDeleteNote}
                    className="px-6 py-2 bg-red-600 text-white font-semibold rounded-lg hover:bg-red-700 transition shadow">
                    Hapus Catatan
                </button>
                </form>
            </div>
        </div>
    );
} 