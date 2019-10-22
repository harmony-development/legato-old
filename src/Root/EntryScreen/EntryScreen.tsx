import React, { useEffect } from 'react';
import { useStyles } from './styles';
import { Paper, Tabs, Tab } from '@material-ui/core';
import LoginForm from './LoginForm/LoginForm';
import RegisterForm from './RegisterForm/RegisterForm';
import { socketServer } from '../Root';
import { Events } from '../../socket/socket';
import { useHistory } from 'react-router';

const EntryScreen: React.FC<{}> = () => {
  const classes = useStyles();
  const [value, setValue] = React.useState(0);
  const history = useHistory();

  useEffect(() => {
    if (typeof localStorage.getItem('token') === 'string') {
      socketServer.getUserData();
      socketServer.connection.on(Events.GET_USER_DATA, () => {
        history.push('/app');
      });
    }
    return (): void => {
      socketServer.connection.removeEventListener(Events.GET_USER_DATA);
    };
  });

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
