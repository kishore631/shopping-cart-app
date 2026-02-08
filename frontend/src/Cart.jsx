import axios from "axios";
import { useEffect, useState } from "react";
import "./App.css";

function Cart() {
  const [cart, setCart] = useState([]);
  const token = localStorage.getItem("token");

  useEffect(() => {
    axios
      .get("http://localhost:8080/carts", {
        headers: { Authorization: "Bearer " + token },
      })
      .then((res) => setCart(res.data))
      .catch(() => alert("Failed to load cart"));
  }, []);

  return (
<div className="cart-page">
  <div className="cart-card">
    <h2>Your Cart</h2>

    {cart.length === 0 && (
      <p className="cart-empty">Cart is empty</p>
    )}

    {cart.map((c, i) => (
      <div key={i} className="cart-item">
        <span>{c.Item.Name}</span>
        <span>₹{c.Item.Price}</span>
      </div>
    ))}

    <button onClick={() => window.location.href = "/items"}>
      Back to Items
    </button>
  </div>
</div>

  );
}

export default Cart;
