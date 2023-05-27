import React from 'react';
import { Container, Navbar, Nav } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import './Header.css';

function Header() {
    return (
        <Navbar expand="lg" className="header-container">
            <Container fluid>
                <Navbar.Brand className="brand">
                    <Link to="/about">GigaChat</Link>
                    <p></p>
                </Navbar.Brand>
                <Navbar.Collapse className="justify-content-end">
                    <Nav className="nav-links">
                        <Link className="nav-link" to="/login">Login</Link>
                        <Link className="nav-link" to="/register">Register</Link>
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default Header;
