import "./Registration.css";
import {useState} from "react";

function Registration() {
    const [login, setLogin] = useState("");
    const [password, setPassword] = useState("");

    const onLoginChanged = (event) => {
        setLogin(event.target.value);
    };

    const onPasswordChanged = (event) => {
        setPassword(event.target.value);
    };

    const onButtonClick = () => {
        console.log(login);
        console.log(password);

        fetch("http://localhost:3003/login", {
            method: "POST",
            mode: "no-cors",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                login: login,
                password: password,
            }),
        })
            .then((response) => response.json())
            .then((data) => {
                console.log(data);
            })
            .catch((error) => {
                console.log(error);
            });

        setLogin("");
        setPassword("");
    };

    return (
        <div className="main-container">
            <div className="form-container">
                <h1 className="title">Sign In</h1>
                <form>
                    <div className="form-group">
                        <label className="form-label">Username</label>
                        <input
                            type="text"
                            placeholder="Enter username"
                            className="form-control"
                            value={login}
                            onChange={onLoginChanged}
                        />
                    </div>
                    <div className="form-group">
                        <label className="form-label">Password</label>
                        <input
                            type="password"
                            placeholder="Enter password"
                            className="form-control"
                            value={password}
                            onChange={onPasswordChanged}
                        />
                    </div>
                    <button type="button" className="sign-button" onClick={onButtonClick}>
                        Sign In
                    </button>
                </form>
            </div>
        </div>

    );
}

export default Registration;
