import React, { useState, useEffect } from "react";
import axios from "axios";
import { Link } from "react-router-dom";

const List = () => {
  const [data, setData] = useState([]);
  const [error, setError] = useState(null);
  const [activeTab, setActiveTab] = useState("publish");

  useEffect(() => {
    axios
      .get("http://localhost:8080/article/posts")
      .then((response) => {
        setData(response.data);
      })
      .catch((error) => {
        setError(error);
      });
  }, []);

  const handleTrash = (id) => {
    axios
      .put(`http://localhost:8080/article/posts/${id}`, { status: "thrash" })
      .then((response) => {
        const updatedData = data.map((post) => (post.id === id ? { ...post, status: "thrash" } : post));
        setData(updatedData);
      })
      .catch((error) => {
        console.error("Error updating post status:", error);
        setError(error);
      });
  };

  const filteredData = data.filter((post) => post.status === activeTab);

  return (
    <div className="container mx-auto p-4">
      <div className="tabs mb-4">
        {["publish", "draft", "thrash"].map((tab) => (
          <button key={tab} className={`tab px-4 py-2 rounded-lg ${activeTab === tab ? "bg-blue-500 text-white" : "bg-gray-200 text-gray-700"} mx-1`} onClick={() => setActiveTab(tab)}>
            {tab}
          </button>
        ))}
        <Link to="/article/new" className="bg-green-500 text-white p-2 rounded">
          Add New
        </Link>
      </div>
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white shadow-md rounded-lg overflow-hidden">
          <thead className="bg-gray-100">
            <tr>
              <th className="py-2 px-4 text-left text-gray-600 font-semibold">ID</th>
              <th className="py-2 px-4 text-left text-gray-600 font-semibold">Title</th>
              <th className="py-2 px-4 text-left text-gray-600 font-semibold">Category</th>
              <th className="py-2 px-4 text-left text-gray-600 font-semibold">Action</th>
            </tr>
          </thead>
          <tbody>
            {filteredData.map((item) => (
              <tr key={item.id} className="border-t border-gray-200">
                <td className="py-2 px-4">{item.id}</td>
                <td className="py-2 px-4">{item.title}</td>
                <td className="py-2 px-4">{item.category}</td>
                <td className="py-2 px-4">
                  <Link to={`/article/edit/${item.id}`} className="mr-2 text-blue-500 hover:underline">
                    Edit
                  </Link>
                  <button onClick={() => handleTrash(item.id)} className="text-red-500 hover:underline">
                    Thrash
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default List;
