import React from 'react';
import ReactDOM from 'react-dom/client';
import Chat from './chat/Chat';
import Registration from './registration/Registration';
import 'bootstrap/dist/css/bootstrap.min.css';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
        <Chat />
        {/*<Registration />*/}
    </React.StrictMode>
);