import React from 'react';
import {Button, Container, Form} from 'react-bootstrap';
import axios from 'axios';
import {useNavigate} from 'react-router-dom';

import {REGISTER_URL} from '../../api/Endpoints';
import './RegisterPage.css';

function RegisterPage() {
    const navigate = useNavigate();

    const register = async (e) => {
        e.preventDefault();
        const username = e.target.elements.username.value;
        const password = e.target.elements.password.value;
        await axios.post(REGISTER_URL, {
            username: username,
            password: password
        });
        navigate('/login');
    };

    return (
        <div className="main-container">
            <div className="form-container">
                <Form onSubmit={register}>
                    <Form.Group controlId="formBasicUsername" className="form-group">
                        <Form.Label>Username </Form.Label>
                        <Form.Control type="text" name="username"/>
                    </Form.Group>

                    <Form.Group controlId="formBasicUsername" className="form-group">
                        <Form.Label>Password </Form.Label>
                        <Form.Control type="password" name="password"/>
                    </Form.Group>
                    <Button variant="primary" type="submit" className="sign-button">
                        Register
                    </Button>
                </Form>
            </div>
        </div>
    );
}

export default RegisterPage;
