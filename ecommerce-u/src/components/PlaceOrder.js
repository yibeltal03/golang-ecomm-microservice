import React, { useState } from 'react';

const PlaceOrder = () => {
  const [product, setProduct] = useState('');
  const [quantity, setQuantity] = useState(1);
  const [username, setUsername] = useState('');
  const [message, setMessage] = useState('');

  const placeOrder = async () => {
    const response = await fetch('http://localhost:8082/order', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ product, quantity, username })
    });
    const data = await response.text();
    setMessage(data);
  };

  return (
    <div>
      <h2>Place Order</h2>
      <input
        type="text"
        placeholder="Product"
        value={product}
        onChange={(e) => setProduct(e.target.value)}
      />
      <input
        type="number"
        placeholder="Quantity"
        value={quantity}
        onChange={(e) => setQuantity(e.target.value)}
      />
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <button onClick={placeOrder}>Place Order</button>
      <p>{message}</p>
    </div>
  );
};

export default PlaceOrder;
