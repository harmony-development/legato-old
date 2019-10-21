import React from 'react';
import { Button, TextField } from '@material-ui/core';
import { useStyles } from './styles';

const RegisterForm: React.FC<{}> = () => {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <TextField label='Email' type='email' name='email' autoComplete='email' margin='normal' fullWidth />
      <TextField label='Username' type='username' name='username' autoComplete='username' margin='normal' fullWidth />
      <TextField label='Password' type='password' name='password' margin='normal' fullWidth />
      <TextField label='Confirm Password' type='password' name='confirmpassword' margin='normal' fullWidth />
      <Button variant='contained' color='primary' className={classes.submitButton}>
        Register
      </Button>
    </div>
  );
};

export default RegisterForm;
