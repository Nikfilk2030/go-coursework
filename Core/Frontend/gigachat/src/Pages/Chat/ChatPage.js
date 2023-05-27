import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Cookies from 'js-cookie';
import { Button, Container, Form } from 'react-bootstrap';

import { MESSAGES_URL, SEND_MESSAGE_URL } from '../../api/Endpoints';
import './ChatPage.css';

function ChatPage() {
    const [messages, setMessages] = useState([]);
    const token = Cookies.get('token');

    useEffect(() => {
        const getMessages = async () => {
            const response = await axios.get(MESSAGES_URL, {
                headers: { 'Authorization': `Bearer ${token}` }
            });
            setMessages(response.data);
        };
        getMessages();
    }, [token]);

    const sendMessage = async (e) => {
        e.preventDefault();
        const content = e.target.elements.content.value;
        await axios.post(SEND_MESSAGE_URL, {
            content: content
        }, {
            headers: { 'Authorization': `Bearer ${token}` }
        });
        e.target.elements.content.value = '';
    };

    return (
        <Container className="chat-container">
            <div className="messages-container">
                {messages.map((message, index) => (
                    <div className="message" key={index}>
                        <span className="username">{message.username}</span>
                        <span className="content">{message.content}</span>
                    </div>
                ))}
            </div>
            <Form onSubmit={sendMessage}>
                <Form.Group>
                    <Form.Control type="text" name="content" />
                </Form.Group>
                <Button type="submit">Send</Button>
            </Form>
        </Container>
    );
}

export default ChatPage;
