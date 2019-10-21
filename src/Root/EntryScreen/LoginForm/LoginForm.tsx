import React, { useEffect } from 'react';
import { TextField, Button, Typography } from '@material-ui/core';
import { useStyles } from './styles';
import { socketServer } from '../../Root';
import { Events } from '../../../socket/socket';

const LoginForm: React.FC<{}> = () => {
  const classes = useStyles();
  const [email, setEmail] = React.useState<string | undefined>(undefined);
  const [password, setPassword] = React.useState<string | undefined>(undefined);
  const [error, setError] = React.useState<string | undefined>(undefined);

  const login = (): void => {
    if (email && password) {
      socketServer.login(email, password);
    } else {
      setError('Missing username or password');
    }
  };

  useEffect(() => {
    socketServer.connection.on(Events.LOGIN_ERROR, (error: string) => {
      setError(error);
    });
    socketServer.connection.on(Events.LOGIN, (token: string) => {
      console.log(token);
    });
  }, []);

  return (
    <div className={classes.root}>
      <TextField label='Email' type='email' name='email' autoComplete='email' margin='normal' fullWidth onChange={(event): void => setEmail(event.target.value)} />
      <TextField label='Password' type='password' name='password' margin='normal' fullWidth onChange={(event): void => setPassword(event.target.value)} />
      {error ? (
        <Typography variant='subtitle1' color={'error'}>
          {error}
        </Typography>
      ) : (
        undefined
      )}
      <Button variant='contained' color='primary' className={classes.submitButton} onClick={login}>
        Log In
      </Button>
    </div>
  );
};

export default LoginForm;
