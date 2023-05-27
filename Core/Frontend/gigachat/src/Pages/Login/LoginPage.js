import React, { useState } from 'react';
import { Button, Form } from 'react-bootstrap';
import axios from 'axios';
import { useNavigate } from 'react-router';
import { DOMAIN } from '../../api/Endpoints';
import "./LoginPage.css"

const LoginPage = () => {
    const navigate = useNavigate();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = async (event) => {
        event.preventDefault();
        const result = await axios.post(DOMAIN.LOGIN_URL, { username, password });
        if (result.data.success) {
            alert('LoginPage successful');
            navigate("/chat");
        } else {
            alert('LoginPage failed');
        }
    };

    return (
        <div className="main-container">
            <div className="form-container">
                <Form onSubmit={handleSubmit}>
                    <Form.Group controlId="formBasicUsername" className="form-group">
                        <Form.Label className="form-label">Username </Form.Label>
                        <Form.Control type="text" placeholder="Enter username" className="form-control" onChange={(e) => setUsername(e.target.value)} />
                    </Form.Group>

                    <Form.Group controlId="formBasicPassword" className="form-group">
                        <Form.Label className="form-label">Password </Form.Label>
                        <Form.Control type="password" placeholder="Password" className="form-control" onChange={(e) => setPassword(e.target.value)} />
                    </Form.Group>
                    <Button variant="primary" type="submit" className="sign-button">
                        Submit
                    </Button>
                </Form>
            </div>
        </div>
    );
}

export default LoginPage;
