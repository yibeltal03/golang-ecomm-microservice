import React, { useState } from 'react';

function ViewOrders() {
  const [username, setUsername] = useState('');
  const [orders, setOrders] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const fetchOrders = async () => {
    setLoading(true);
    setError('');
    
    try {
      const response = await fetch(`http://localhost:8082/orders/${username}`);
      
      if (!response.ok) {
        throw new Error("Failed to fetch orders");
      }

      const data = await response.json();
      setOrders(data);
    } catch (error) {
      setError(error.message);
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (username.trim()) {
      fetchOrders();
    } else {
      setError("Please enter a valid username.");
    }
  };

  return (
    <div>
      <h2>View Orders</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          placeholder="Enter username"
          required
        />
        <button type="submit">View Orders</button>
      </form>

      {loading && <p>Loading orders...</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}

      <div>
        <h3>Orders for {username}:</h3>
        {orders.length > 0 ? (
          <ul>
            {orders.map((order, index) => (
              <li key={index}>
                Product: {order.product}, Quantity: {order.quantity}
              </li>
            ))}
          </ul>
        ) : (
          !loading && <p>No orders found.</p>
        )}
      </div>
    </div>
  );
}

export default ViewOrders;
