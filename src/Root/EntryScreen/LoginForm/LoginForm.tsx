import React from 'react';
import { TextField, Button } from '@material-ui/core';
import { useStyles } from './styles';

const LoginForm: React.FC<{}> = () => {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <TextField label='Email' type='email' name='email' autoComplete='email' margin='normal' fullWidth />
      <TextField label='Password' type='password' name='password' margin='normal' fullWidth />
      <Button variant='contained' color='primary' className={classes.submitButton}>
        Log In
      </Button>
    </div>
  );
};

export default LoginForm;
