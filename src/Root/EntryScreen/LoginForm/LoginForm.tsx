import React, { useEffect } from 'react';
import { TextField, Button, Typography } from '@material-ui/core';
import { useStyles } from './styles';
import { socketServer } from '../../Root';
import { Events } from '../../../socket/socket';
import { useHistory } from 'react-router';
import { ILoginDetails } from '../../types';
import { useDispatch } from 'react-redux';
import { updateUser } from '../../../store/actions/AppActions';
import { toast } from 'react-toastify';

const LoginForm: React.FC<{}> = () => {
  const classes = useStyles();
  const [email, setEmail] = React.useState<string | undefined>(undefined);
  const [password, setPassword] = React.useState<string | undefined>(undefined);
  const [error, setError] = React.useState<string | undefined>(undefined);
  const dispatch = useDispatch();
  const history = useHistory();

  const login = (): void => {
    if (!socketServer.connection.connected) {
      toast.error('Unable to login. No connection to server.');
    }
    if (email && password) {
      socketServer.login(email, password);
    } else {
      setError('Missing username or password');
    }
  };

  const onFormSubmit = (e: React.FormEvent<EventTarget>): void => {
    e.preventDefault();
  };

  useEffect(() => {
    socketServer.connection.on(Events.LOGIN_ERROR, (error: string) => {
      setError(error);
    });

    socketServer.connection.on(Events.LOGIN, (response: ILoginDetails) => {
      history.push('/app');
      localStorage.setItem('token', response.token);
      dispatch(updateUser({ username: response.username, avatar: response.avatar }));
    });

    return (): void => {
      // cleanup event listeners
      socketServer.connection.removeListener(Events.LOGIN);
      socketServer.connection.removeListener(Events.LOGIN_ERROR);
    };
  }, [dispatch, history]);

  return (
    <div className={classes.root}>
      <form onSubmit={onFormSubmit}>
        <TextField label='Email' type='email' name='email' autoComplete='email' margin='normal' fullWidth onChange={(event): void => setEmail(event.target.value)} />
        <TextField label='Password' type='password' name='password' margin='normal' fullWidth onChange={(event): void => setPassword(event.target.value)} />
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
