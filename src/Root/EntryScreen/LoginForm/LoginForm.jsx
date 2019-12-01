import React, { useEffect } from 'react';
import { TextField, Button, Typography } from '@material-ui/core';
import { useStyles } from './styles';
import { socketServer } from '../../Root';
import { Events } from '../../../socket/socket';
import { useHistory } from 'react-router';
import { useDispatch } from 'react-redux';
import { updateUser } from '../../../store/actions/AppActions';
import { toast } from 'react-toastify';

const LoginForm = () => {
    const classes = useStyles();
    const [email, setEmail] = React.useState(undefined);
    const [password, setPassword] = React.useState(undefined);
    const [error, setError] = React.useState(undefined);
    const dispatch = useDispatch();
    const history = useHistory();

    const login = () => {
        if (!socketServer.connection.connected) {
            toast.error('Unable to login. No connection to server.');
        }
        if (email && password) {
            socketServer.login(email, password);
        } else {
            setError('Missing username or password');
        }
    };

    const onFormSubmit = (e) => {
        e.preventDefault();
    };

    useEffect(() => {
        socketServer.connection.on(Events.LOGIN_ERROR, (error) => {
            setError(error);
        });

        socketServer.connection.on(Events.LOGIN, (response) => {
            history.push('/app');
            localStorage.setItem('token', response.token);
            dispatch(updateUser({ username: response.username, avatar: response.avatar }));
        });

        return () => {
            // cleanup event listeners
            socketServer.connection.removeListener(Events.LOGIN);
            socketServer.connection.removeListener(Events.LOGIN_ERROR);
        };
    }, [dispatch, history]);

    return (
        <div className={classes.root}>
            <form onSubmit={onFormSubmit}>
                <TextField label='Email' type='email' name='email' autoComplete='email' margin='normal' fullWidth onChange={(event) => setEmail(event.target.value)} />
                <TextField label='Password' type='password' name='password' margin='normal' fullWidth onChange={(event) => setPassword(event.target.value)} />
                {error ? (
                    <Typography variant='subtitle1' color={'error'}>
                        {error}
                    </Typography>
                ) : (
                    undefined
                )}
                <Button variant='contained' color='primary' className={classes.submitButton} onClick={login} type='submit'>
                    Log In
                </Button>
            </form>
        </div>
    );
};

export default LoginForm;
