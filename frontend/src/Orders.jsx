import axios from "axios";
import { useEffect, useState } from "react";

function Orders() {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    axios.get("http://localhost:8080/orders").then((res) => {
      setOrders(res.data);
    });
  }, []);

  return (
    <div>
      <h2>Orders</h2>
      {orders.map((o, i) => (
        <ul key={i}>
          {o.Items.map((item, j) => (
            <li key={j}>{item.Name}</li>
          ))}
        </ul>
      ))}
    </div>
  );
}

export default Orders;
