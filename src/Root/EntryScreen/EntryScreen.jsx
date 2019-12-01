import React, { useEffect } from 'react';
import { useStyles } from './styles';
import { Paper, Tabs, Tab } from '@material-ui/core';
import LoginForm from './LoginForm/LoginForm';
import RegisterForm from './RegisterForm/RegisterForm';
import { socketServer } from '../Root';
import { useHistory } from 'react-router';

const EntryScreen = () => {
    const classes = useStyles();
    const [value, setValue] = React.useState(0);
    const history = useHistory();

    useEffect(() => {
        if (typeof localStorage.getItem('token') === 'string') {
            socketServer.emitter.addListener('getservers', () => {
                history.push('/app');
            });
        }
        return () => {
            socketServer.emitter.removeAllListeners('getservers');
        };
    });

    return (
        <div className={classes.root}>
            <Paper className={classes.form}>
                <Tabs value={value} onChange={(event, newValue) => setValue(newValue)} variant='fullWidth' indicatorColor={'primary'}>
                    <Tab label='Login' id='form-tab-0' />
                    <Tab label='Register' id='form-tab-1' />
                </Tabs>
                {value === 0 ? <LoginForm /> : <RegisterForm />}
            </Paper>
        </div>
    );
};

export default EntryScreen;
