import React, { useEffect } from 'react';
import { CssBaseline, createMuiTheme, Button } from '@material-ui/core';
import { ThemeProvider } from '@material-ui/core/styles';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast } from 'react-toastify';
import { useSelector, useDispatch } from 'react-redux';
import { IState, Actions } from '../types/redux';
import HarmonySocket from '../socket/socket';
import { Switch, Route } from 'react-router';
import { BrowserRouter } from 'react-router-dom';
import './Root.css';
import { App } from './App/App';
import { Entry } from './Entry/Entry';

export const harmonySocket = new HarmonySocket();
export let previouslyDisconnected = false;

const Root: React.FC = () => {
    const themeState = useSelector((state: IState) => state.theme);
    const dispatch = useDispatch();

    useEffect(() => {
        harmonySocket.events.addListener('close', () => {
            // lol plz no spahm
            if (!previouslyDisconnected) {
                toast.error('You have lost connection to the server');
                dispatch({ type: Actions.SET_CONNECTED, payload: false });
                previouslyDisconnected = true;
            }
            setTimeout(harmonySocket.connect, 3000);
        });
        harmonySocket.events.addListener('open', () => {
            if (previouslyDisconnected) toast.success('You have reconnected to the server');
            dispatch({ type: Actions.SET_CONNECTED, payload: true });
        });
        return () => {
            harmonySocket.events.removeAllListeners('close'); // cleanup all socket events registered here
            harmonySocket.events.removeAllListeners('open');
        };
    }, [dispatch]);

    const theme = createMuiTheme({
        palette: {
            primary: themeState.primary,
            secondary: themeState.secondary,
            type: themeState.type
        }
    });

    return (
        <div className='root'>
            <ThemeProvider theme={theme}>
                <CssBaseline />
                <ToastContainer />
                <BrowserRouter>
                    <Switch>
                        <Route exact path='/'>
                            <Entry />
                        </Route>
                        <Route exact path='/app'>
                            <App />
                        </Route>
                        <Route exact path='/bruh'>
                            <Button onClick={()=>{toast.info("GET BRUHED ON KID")}}>Bruh Buttyon</Button>
                        </Route>
                    </Switch>
                </BrowserRouter>
            </ThemeProvider>
        </div>
    );
};

export default Root;
