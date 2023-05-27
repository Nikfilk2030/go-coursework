import React from 'react';
import { Container } from 'react-bootstrap';

import './AboutPage.css';

function AboutPage() {
    return (
        <Container className="about-container">
            <div className="about-jumbotron">
                <h1>GigaChat</h1>
                <p>Добро пожаловать в GigaChat</p>
                <p>Разработано Никитой Шевердовым</p>
            </div>
        </Container>
    );
}

export default AboutPage;
