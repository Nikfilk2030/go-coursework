import React, { useState, useEffect, useRef } from 'react';
import './App.css';

function App() {
    const [messages, setMessages] = useState([]);
    const [inputValue, setInputValue] = useState('');

    const socket = useRef(null);

    const [socketStatus, setSocketStatus] = useState('disconnected');

    useEffect(() => {
        const socket = new WebSocket('ws://localhost:8080/ws');

        socket.onopen = () => {
            console.log('WebSocket connection opened');
            setSocketStatus('connected');
        };

        socket.onmessage = (event) => {
            const newMessage = JSON.parse(event.data);
            setMessages((prevMessages) => [...prevMessages, newMessage]);
        };

        socket.onclose = (event) => {
            console.log('WebSocket connection closed with code:', event.code);
            setSocketStatus('disconnected');
        };

        socket.onerror = (error) => {
            console.error('WebSocket error:', error);
            setSocketStatus('disconnected');
        };

        return () => {
            socket.close();
        };
    }, []);


    const handleSubmit = (event) => {
        event.preventDefault();
        if (!inputValue.trim()) return;
        if (socket.current && socket.current.readyState === WebSocket.OPEN) {
            const message = { author: 'Anonymous', text: inputValue };
            socket.current.send(JSON.stringify(message));
            setInputValue('');
        } else {
            console.warn('WebSocket connection is not open');
        }
    };

    return (
        <div className="container">
            <div className="chat">
                {messages.map((message, index) => (
                    <div className="message" key={index}>
                        <div className="author">{message.author}</div>
                        <div className="text">{message.text}</div>
                    </div>
                ))}
            </div>
            <form onSubmit={handleSubmit}>
                <div className="input-container">
                    <input
                        type="text"
                        value={inputValue}
                        onChange={(event) => setInputValue(event.target.value)}
                        className="input"
                        placeholder="Type a message"
                    />
                    <button type="submit" className="send-button">
                        Send
                    </button>
                </div>
            </form>
        </div>
    );
}

export default App;
