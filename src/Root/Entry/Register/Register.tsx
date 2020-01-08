import React, { useRef, useEffect } from 'react';
import { TextField, Typography, Button } from '@material-ui/core';
import { toast } from 'react-toastify';
import { useHistory } from 'react-router';

import { harmonySocket } from '../../Root';

import { useRegisterStyles } from './RegisterStyle';

export const Register = () => {
	const classes = useRegisterStyles();
	const history = useHistory();
	const [err, setErr] = React.useState<string | undefined>(undefined);
	const emailRef = useRef<HTMLInputElement | undefined>(undefined);
	const usernameRef = useRef<HTMLInputElement | undefined>(undefined);
	const pwdRef = useRef<HTMLInputElement | undefined>(undefined);

	const register = () => {
		if (harmonySocket.conn.readyState === WebSocket.CLOSED) {
			toast.error("Can't register, trouble connecting to server");
		} else if (
			emailRef.current &&
			usernameRef.current &&
			pwdRef.current &&
			emailRef.current.value &&
			pwdRef.current.value &&
			usernameRef.current.value
		) {
			harmonySocket.register(emailRef.current.value, usernameRef.current.value, pwdRef.current.value);
		} else {
			toast.error("Can't register, missing email, username, or password");
		}
	};

	useEffect(() => {
		harmonySocket.events.addListener('registererror', (raw: any) => {
			if (typeof raw['message'] === 'string') {
				setErr(raw['message']);
			}
		});
		harmonySocket.events.addListener('token', (raw: any) => {
			if (typeof raw['token'] === 'string' && typeof raw['userid'] === 'string') {
				localStorage.setItem('token', raw['token']);
				localStorage.setItem('userid', raw['userid']);
				harmonySocket.getGuilds();
				history.push('/app');
			}
		});
		return () => {
			harmonySocket.events.removeAllListeners('registererror');
			harmonySocket.events.removeAllListeners('token');
		};
	}, [history]);

	return (
		<div className={classes.root}>
			<form onSubmit={(e: React.FormEvent<EventTarget>) => e.preventDefault()}>
				<TextField
					label="Email"
					type="email"
					name="email"
					autoComplete="email"
					margin="normal"
					fullWidth
					inputRef={emailRef}
				/>
				<TextField
					label="Username"
					type="text"
					name="usernamee"
					autoComplete="username"
					margin="normal"
					fullWidth
					inputRef={usernameRef}
				/>
				<TextField
					label="Password"
					type="password"
					name="password"
					autoComplete="new-password"
					margin="normal"
					fullWidth
					inputRef={pwdRef}
				/>
				<TextField
					label="Confirm Password"
					type="password"
					name="confirmpassword"
					autoComplete="none"
					margin="normal"
					fullWidth
				/>
				{err ? (
					<Typography variant="subtitle1" color={'error'}>
						{err}
					</Typography>
				) : (
					undefined
				)}
				<Button variant="contained" color="primary" className={classes.submitBtn} onClick={register} type="submit">
					Log In
				</Button>
			</form>
		</div>
	);
};
