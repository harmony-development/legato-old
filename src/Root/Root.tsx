/**
 * This file is intended for adding things such as Redux Providers and other things.
 */
import React, { useEffect } from 'react';
import { Provider, useDispatch } from 'react-redux';
import { store } from '../store/store';
import { useStyles } from './styles';
import { createMuiTheme, CssBaseline } from '@material-ui/core';
import { useSelector } from 'react-redux';
import { IAppState } from '../store/types';
import { ThemeProvider } from '@material-ui/styles';
import { Switch, Route, BrowserRouter as Router } from 'react-router-dom';
import EntryScreen from './EntryScreen/EntryScreen';
import App from './App/App';
import { HarmonyConnection } from '../socket/socket';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export const socketServer = new HarmonyConnection();
let previouslyDisconnected = false;

const Theme: React.FC<{}> = () => {
    useStyles();

    return <></>;
};

socketServer.connection.addEventListener('open', () => {
    if (previouslyDisconnected) {
        toast.success('You have reconnected to the server');
    }
});

socketServer.connection.addEventListener('close', () => {
    toast.error('You have lost connection to the server');
    previouslyDisconnected = true;
});

const Root: React.FC<{}> = () => {
    const { type, primary, secondary } = useSelector((state: IAppState) => state.theme);
    const dispatch = useDispatch();
    const theme = createMuiTheme({
        palette: {
            type,
            primary,
            secondary
        }
    });

    useEffect(() => {
        // TODO : ADD PROFILE SAVING
    }, [type, primary, secondary]);

    useEffect(() => {
        //TODO : ADD GETTING PROFILE FEATURES
    }, [dispatch]);

    return (
        <div className='app-container'>
            <ThemeProvider theme={theme}>
                <Theme />
                <CssBaseline />
                <Router>
                    <Switch>
                        <Route exact path='/' component={EntryScreen}></Route>
                        <Route exact path='/app' component={App}></Route>
                    </Switch>
                </Router>
                <ToastContainer />
            </ThemeProvider>
        </div>
    );
};

const ReduxRoot: React.FC<{}> = () => {
    return (
        <Provider store={store}>
            <Root />
        </Provider>
    );
};

export default ReduxRoot;
