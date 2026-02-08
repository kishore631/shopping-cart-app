import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./Login";
import Items from "./Items";
import Cart from "./Cart";

ReactDOM.createRoot(document.getElementById("root")).render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<Login />} />
      <Route path="/items" element={<Items />} />
      <Route path="/cart" element={<Cart />} />
    </Routes>
  </BrowserRouter>
);
