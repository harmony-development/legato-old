import React, { useRef, useEffect } from 'react';
import { TextField, Typography, Button } from '@material-ui/core';
import { useLoginStyles } from './LoginStyle';
import { harmonySocket } from '../../Root';
import { toast } from 'react-toastify';
import { useHistory } from 'react-router';

export const Login: React.FC = () => {
    const history = useHistory(); // history for better routing

    const [err, setErr] = React.useState<string | undefined>(undefined);
    const emailRef = useRef<HTMLInputElement | undefined>(undefined);
    const pwdRef = useRef<HTMLInputElement | undefined>(undefined);
    const classes = useLoginStyles();
    const onFormSubmit = (e: React.FormEvent<EventTarget>) => e.preventDefault();

    const login = () => {
        if (harmonySocket.conn.readyState === WebSocket.CLOSED) {
            toast.error("Can't login, trouble connecting to server");
        } else if (emailRef.current && pwdRef.current && emailRef.current.value.length && pwdRef.current.value) {
            harmonySocket.login(emailRef.current.value, pwdRef.current.value);
        } else {
            toast.error("Can't login, missing email or password");
        }
    };

    useEffect(() => {
        harmonySocket.events.addListener('loginerror', (raw: any) => {
            if (typeof raw['message'] === 'string') {
                setErr(raw['message']);
            }
        });
        harmonySocket.events.addListener('token', (raw: any) => {
            if (typeof raw['token'] === 'string') {
                localStorage.setItem('token', raw['token']);
                history.push('/app');
            }
        });
        return () => {
            harmonySocket.events.removeAllListeners('loginerror');
            harmonySocket.events.removeAllListeners('token');
        };
    }, [history]);

    return (
        <div className={classes.root}>
            <form onSubmit={onFormSubmit}>
                <TextField label='Email' type='email' name='email' autoComplete='email' margin='normal' fullWidth inputRef={emailRef} />
                <TextField label='Password' type='password' name='password' margin='normal' fullWidth inputRef={pwdRef} />
                {err ? (
                    <Typography variant='subtitle1' color={'error'}>
                        {err}
                    </Typography>
                ) : (
                    undefined
                )}
                <Button variant='contained' color='primary' className={classes.submitBtn} onClick={login} type='submit'>
                    Log In
                </Button>
            </form>
        </div>
    );
};
