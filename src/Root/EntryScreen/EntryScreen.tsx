import React from 'react';
import { useStyles } from './styles';
import { Paper, Button, Tabs, Tab } from '@material-ui/core';
import LoginForm from './LoginForm/LoginForm';
import RegisterForm from './RegisterForm/RegisterForm';

const EntryScreen: React.FC<{}> = () => {
  const classes = useStyles();
  const [value, setValue] = React.useState(0);

  return (
    <div className={classes.root}>
      <Paper className={classes.form}>
        <Tabs value={value} onChange={(event: React.ChangeEvent<{}>, newValue: number): void => setValue(newValue)} variant='fullWidth' indicatorColor={'primary'}>
          <Tab label='Login' id='form-tab-0' />
          <Tab label='Register' id='form-tab-1' />
        </Tabs>
        {value === 0 ? <LoginForm /> : <RegisterForm />}
      </Paper>
    </div>
  );
};

export default EntryScreen;
