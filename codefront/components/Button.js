import { useState } from 'react';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import { sendToPrinter } from '../api/print';

function Print() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const onClick = () => {
        const params = {
            "username": username,
            "password": password
        };
        sendToPrinter(params)
            .then((res) => {
                console.log(res);
            })
            .catch((err) => {
                console.log(err.code);
                switch (err) {
                    case 401:
                    case 404:
                        alert("用户名或密码错误");
                        break;
                    default:
                        alert("未知错误");
                        break;
                }
            })
    };

    return (
        <div>
            <TextField
                className="relative left-0 bg-slate-50"
                id="username"
                label="Username"
                variant="filled"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
            />
            <TextField
                className="relative left-5 bg-slate-50"
                id="password"
                label="Password"
                variant="filled"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <Button
                className="relative left-10 text-white sm:px-8 py-2 sm:py-3 bg-sky-700 hover:bg-sky-800"
                onClick={onClick}
            >
                Print it!
            </Button>
        </div>
    );
}

export default Print;