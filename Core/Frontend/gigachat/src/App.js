import React from 'react';
import {
  BrowserRouter as Router,
  Routes,
  Route
} from 'react-router-dom';
import { Container } from 'react-bootstrap';

import Header from './Components/Header/Header';
import LoginPage from './Pages/Login/LoginPage';
import RegisterPage from './Pages/Register/RegisterPage';
import AboutPage from './Pages/About/AboutPage';
import ChatPage from './Pages/Chat/ChatPage';
import './App.css';

function App() {
  return (
      <Router>
        <Container className="App">
          <Header />
          <Routes>
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/about" element={<AboutPage />} />
            <Route path="/chat" element={<ChatPage />} />
          </Routes>
        </Container>
      </Router>
  );
}

export default App;
