import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const AddPost = () => {
  const navigate = useNavigate();
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [category, setCategory] = useState("");
  const [status, setStatus] = useState("draft");

  const handleSubmit = () => {
    axios
      .post("http://localhost:8080/article/posts", { title, content, category, status })
      .then((response) => {
        console.log("Post created", response.data);
        navigate("/");
      })
      .catch((error) => {
        console.error("Error creating post", error);
        alert(`Error: ${error.response?.data?.error || error.message}`);
      });
  };

  const handleCancel = () => {
    navigate("/");
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Add New Post</h1>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          handleSubmit();
        }}
      >
        <div className="mb-4">
          <label className="block text-gray-700">Title</label>
          <input type="text" value={title} onChange={(e) => setTitle(e.target.value)} className="mt-1 p-2 border rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Content</label>
          <textarea value={content} onChange={(e) => setContent(e.target.value)} className="mt-1 p-2 border rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Category</label>
          <input type="text" value={category} onChange={(e) => setCategory(e.target.value)} className="mt-1 p-2 border rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Status</label>
          <select value={status} onChange={(e) => setStatus(e.target.value)} className="mt-1 p-2 border rounded w-full">
            <option value="draft">Draft</option>
            <option value="publish">Publish</option>
          </select>
        </div>
        <div className="flex space-x-4">
          <button type="submit" className="bg-blue-500 text-white p-2 rounded">
            Submit
          </button>
          <button type="button" onClick={handleCancel} className="bg-gray-500 text-white p-2 rounded">
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
};

export default AddPost;
