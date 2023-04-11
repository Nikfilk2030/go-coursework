import React, {useEffect, useState} from 'react';
import './Chat.css';
import initialPosts from './posts';

function Chat() {
    const [inputValue, setInputValue] = useState("");
    const [messages, setMessages] = useState(initialPosts);

    // const handleSubmit = (event) => {
    //     console.log("handleSubmit");
    //
    //     event.preventDefault();
    //     if (!inputValue.trim()) return;
    //     fetch("http://localhost:8080/messages", {
    //         method: "POST",
    //         mode: "no-cors",
    //         headers: {
    //             "Content-Type": "application/json",
    //         },
    //         body: JSON.stringify({
    //             author: "anon",
    //             text: inputValue,
    //         }),
    //     })
    //         .then((response) => response.json())
    //         .then((data) => {
    //             setMessages([...messages, {author: "anon", text: inputValue}]);
    //             setInputValue('');
    //             console.log(data);
    //         })
    //         .catch((error) => {
    //             console.log(error);
    //         });
    // };

    useEffect(() => {
        const fetchMessages = async () => {
            try {
                const response = await fetch('http://localhost:8080/messages');
                const data = await response.json();
                setMessages(data);
            } catch (error) {
                console.error(error);
            }
        };
        fetchMessages().then(r => console.log("123"));
    }, []);

    const onButtonClick = (event) => {
        // event.preventDefault();
        console.log(inputValue);
        console.log("onButtonClick");

        fetch("http://localhost:8080/messages", {
            method: "POST",
            mode: "no-cors",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                author: "anon",
                text: inputValue,
            }),
        })
            .then((response) => response.json())
            .then((data) => {
                setMessages([...messages, {author: "anon", text: inputValue}]);
                setInputValue('');
                console.log(data);
            })
            .catch((error) => {
                console.log(error);
            });
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
            {/*<form onSubmit={handleSubmit}>*/}
            <form>
                <div className="input-container">
                    <input
                        type="text"
                        value={inputValue}
                        onChange={(event) => setInputValue(event.target.value)}
                        className="input"
                        placeholder="Type a message"
                    />
                    <button type="submit" className="send-button" onClick={onButtonClick}>
                        Send
                    </button>
                </div>
            </form>
        </div>
    );
}

export default Chat;
