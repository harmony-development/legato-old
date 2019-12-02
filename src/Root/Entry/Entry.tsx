import React, { useEffect } from 'react';
import { Paper, Tabs, Tab } from '@material-ui/core';
import { useEntryStyles } from './EntryStyle';
import { Login } from './Login/Login';
import { Register } from './Register/Register';
import { useHistory } from 'react-router';
import { harmonySocket } from '../Root';

export const Entry = () => {
    const history = useHistory();
    const classes = useEntryStyles();
    const [tabIDX, setTabIDX] = React.useState(0);

    useEffect(() => {
        if (typeof localStorage.getItem('token') === 'string' && harmonySocket.conn.readyState === WebSocket.OPEN) {
            history.push('/app');
        }
    }, [history]);

    return (
        <div className={classes.root}>
            <Paper className={classes.form}>
                <Tabs value={tabIDX} onChange={(event: any, newValue: number): void => setTabIDX(newValue)} variant='fullWidth' indicatorColor={'primary'}>
                    <Tab label='Login' id='form-tab-0' />
                    <Tab label='Register' id='form-tab-1' />
                </Tabs>
                {tabIDX === 0 ? <Login /> : <Register />}
            </Paper>
        </div>
    );
};
