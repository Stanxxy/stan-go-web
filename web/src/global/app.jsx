import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import HomePage from '../pages/homePage';
import MenuPage from '../pages/menuPage';
import CartPage from '../pages/cartPage';
import CheckoutPage from '../pages/checkoutPage';
import DiscoverPage from '../pages/discoverPage';
import Login from '../pages/loginPage';
import Navigation from './navigation';
import Footer from './footer';



export default function App() {
  return (
    <Router>
      <Navigation />
      <Routes>
        <Route exact path="/"  element={<HomePage />} />
        <Route exact path="/menu" element={<MenuPage />} />
        <Route exact path="/cart" element={<CartPage />} />
        <Route exact path="/checkout" element={<CheckoutPage />} />
        <Route exact path="/discover" element={<DiscoverPage />} />
        <Route exact path="/login" element={<Login />} />
      </Routes>
      <Footer />
    </Router>
  );
}
