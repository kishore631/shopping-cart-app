import axios from "axios";
import { useEffect, useState } from "react";
import "./App.css";

function Items() {
  const [items, setItems] = useState([]);
  const token = localStorage.getItem("token");

  // 🔐 Protect route + load items
  useEffect(() => {
    if (!token) {
      window.location.href = "/";
      return;
    }

    axios
      .get("http://localhost:8080/items", {
        headers: { Authorization: "Bearer " + token },
      })
      .then((res) => setItems(res.data))
      .catch(() => alert("Failed to load items"));
  }, []);

  // 🛒 Add to cart
  const addToCart = async (itemId) => {
    try {
      await axios.post(
        "http://localhost:8080/carts",
        { item_id: itemId },
        { headers: { Authorization: "Bearer " + token } }
      );
      alert("Added to cart ✅");
    } catch {
      alert("Item not found");
    }
  };

  // ✅ Checkout
  const checkout = async () => {
    try {
      await axios.post(
        "http://localhost:8080/orders",
        {},
        { headers: { Authorization: "Bearer " + token } }
      );
      alert("Order placed successfully ✅");
    } catch {
      alert("Checkout failed");
    }
  };

  // 📦 Order History
  const viewOrders = async () => {
    try {
      const res = await axios.get("http://localhost:8080/orders", {
        headers: { Authorization: "Bearer " + token },
      });

      if (res.data.length === 0) {
        alert("No orders placed yet");
        return;
      }

      const ids = res.data.map((o) => o.ID).join(", ");
      alert("Your Orders: " + ids);
    } catch {
      alert("Failed to fetch orders");
    }
  };

  // 🔓 Logout
  const logout = () => {
    localStorage.removeItem("token");
    window.location.href = "/";
  };

  return (
    <div className="page">
      {/* 🔹 TOP BAR */}
      <div className="top-bar">
        <button onClick={() => (window.location.href = "/cart")}>
          View Cart
        </button>

        <button onClick={viewOrders}>
          Order History
        </button>

        <button onClick={checkout}>
          Checkout
        </button>

        <button className="logout-btn" onClick={logout}>
          Logout
        </button>
      </div>

      <h2 className="products-title">Products</h2>

      {/* 🔹 PRODUCTS GRID */}
      <div className="products-grid">
        {items.map((item) => (
          <div key={item.ID} className="product-card">
            <h3>{item.Name}</h3>
            <p className="price">₹{item.Price}</p>
            <button onClick={() => addToCart(item.ID)}>
              Add to Cart
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Items;
