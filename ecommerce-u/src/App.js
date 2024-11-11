import React from 'react';
import './App.css'; // Import the global CSS file
import RegisterUser from './components/RegisterUser';
import PlaceOrder from './components/PlaceOrder';
import ViewOrders from './components/ViewOrders';

function App() {
  return (
    <div className="App">
      <h1>E-commerce App</h1>
      <div className="card">
        <RegisterUser />
      </div>
      <div className="card">
        <PlaceOrder />
      </div>
      <div className="card">
        <ViewOrders />
      </div>
    </div>
  );
}

export default App;
