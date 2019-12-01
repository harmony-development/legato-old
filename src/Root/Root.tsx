import React from 'react';
import { CssBaseline, createMuiTheme } from '@material-ui/core';
import { ThemeProvider } from '@material-ui/core/styles';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from 'react-toastify';
import { useSelector } from 'react-redux';
import { IState } from '../types/redux';
import HarmonySocket from '../socket/socket';
import { Switch, Route } from 'react-router';
import { BrowserRouter } from 'react-router-dom';
import './Root.css';
import { App } from './App/App';
import { Entry } from './Entry/Entry';

export const harmonySocket = new HarmonySocket();

const Root: React.FC = () => {
    const themeState = useSelector((state: IState) => state.theme);

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
                    </Switch>
                </BrowserRouter>
            </ThemeProvider>
        </div>
    );
};

export default Root;
