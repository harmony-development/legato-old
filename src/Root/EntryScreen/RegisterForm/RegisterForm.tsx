import React, { useEffect } from 'react';
import { Button, TextField, Typography } from '@material-ui/core';
import { useStyles } from './styles';
import { socketServer } from '../../Root';
import { Events } from '../../../socket/socket';

const RegisterForm: React.FC<{}> = () => {
  const classes = useStyles();
  const [email, setEmail] = React.useState<string | undefined>(undefined);
  const [username, setUsername] = React.useState<string | undefined>(undefined);
  const [password, setPassword] = React.useState<string | undefined>(undefined);
  const [confirmPassword, setConfirmPassword] = React.useState<
    string | undefined
  >(undefined);
  const [error, setError] = React.useState<string | undefined>(undefined);

  const register = (): void => {
    if (confirmPassword !== password) {
      setError('Passwords do not match!');
    } else {
      if (email && username && password) {
        socketServer.register(email, username, password);
      } else {
        setError(`Missing email, username, or password`);
      }
    }
  };

  useEffect(() => {
    socketServer.connection.on(Events.REGISTER_ERROR, (error: string) => {
      setError(error);
    });
    socketServer.connection.on(Events.REGISTER, (token: string) => {
      console.log(token);
    });

    return (): void => {
      // cleanup event listeners
      socketServer.connection.removeListener(Events.REGISTER);
      socketServer.connection.removeListener(Events.REGISTER_ERROR);
    };
  }, []);

  return (
    <div className={classes.root}>
      <TextField
        label='Email'
        type='email'
        name='email'
        autoComplete='email'
        margin='normal'
        fullWidth
        onChange={(event): void => setEmail(event.target.value)}
      />
      <TextField
        label='Username'
        type='username'
        name='username'
        autoComplete='username'
        margin='normal'
        fullWidth
        onChange={(event): void => setUsername(event.target.value)}
      />
      <TextField
        label='Password'
        type='password'
        name='password'
        margin='normal'
        fullWidth
        onChange={(event): void => setPassword(event.target.value)}
      />
      <TextField
        label='Confirm Password'
        type='password'
        name='confirmpassword'
        margin='normal'
        fullWidth
        onChange={(event): void => setConfirmPassword(event.target.value)}
      />
      {error ? (
        <Typography variant='subtitle1' color={'error'}>
          {error}
        </Typography>
      ) : (
        undefined
      )}
      <Button
        variant='contained'
        color='primary'
        className={classes.submitButton}
        onClick={register}
      >
        Register
      </Button>
    </div>
  );
};

export default RegisterForm;
