import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import List from "./components/List";
import EditPost from "./components/EditPost";
import AddPost from "./components/AddPost";

const App = () => {
  return (
    <div>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<List />}></Route>
          <Route path="/article/edit/:id" element={<EditPost />}></Route>
          <Route path="/article/new" element={<AddPost />}></Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
};

export default App;
