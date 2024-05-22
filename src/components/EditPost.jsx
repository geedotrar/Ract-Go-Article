import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";

const EditPost = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [category, setCategory] = useState("");
  const [status, setStatus] = useState("");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios
      .get(`http://localhost:8080/article/posts/${id}`)
      .then((response) => {
        const { title, content, category, status } = response.data;
        setTitle(title);
        setContent(content);
        setCategory(category);
        setStatus(status);
        setLoading(false);
      })
      .catch((error) => {
        setError(error.message);
        setLoading(false);
      });
  }, [id]);

  const handleSubmit = (e) => {
    e.preventDefault();
    const postData = { title, content, category, status };
    axios
      .put(`http://localhost:8080/article/posts/${id}`, postData)
      .then((response) => {
        console.log("Post updated", response.data);
        alert("Post updated successfully!");
        navigate("/");
      })
      .catch((error) => {
        // Menangkap pesan kesalahan dari backend dan menampilkan alert
        if (error.response && error.response.data && error.response.data.error) {
          alert(`Error updating post: ${error.response.data.error}`);
        } else {
          alert("Error updating post");
        }
        console.error("Error updating post", error);
      });
  };

  const handleCancel = () => {
    navigate("/");
  };

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Edit Post</h1>
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label className="block text-gray-700">Judul</label>
          <input type="text" value={title} onChange={(e) => setTitle(e.target.value)} className="mt-1 p-2 border rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Konten</label>
          <textarea value={content} onChange={(e) => setContent(e.target.value)} className="mt-1 p-2 border rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Kategori</label>
          <input type="text" value={category} onChange={(e) => setCategory(e.target.value)} className="mt-1 p-2 border rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Status</label>
          <select value={status} onChange={(e) => setStatus(e.target.value)} className="mt-1 p-2 border rounded w-full">
            <option value="publish">Publish</option>
            <option value="draft">Draft</option>
            <option value="thrash">Thrash</option>
          </select>
        </div>
        {error && <div className="mb-4 text-red-500">{error}</div>}
        <button type="submit" className="bg-blue-500 text-white p-2 rounded mr-2">
          Submit
        </button>
        <button type="button" onClick={handleCancel} className="bg-gray-500 text-white p-2 rounded">
          Cancel
        </button>
      </form>
    </div>
  );
};

export default EditPost;
