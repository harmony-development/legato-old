import React from 'react';
import './Root.css';
import { CssBaseline, createMuiTheme } from '@material-ui/core';
import { ThemeProvider } from '@material-ui/core/styles';
import { ToastContainer } from 'react-toastify';
import { useSelector } from 'react-redux';
import { IState } from '../types/redux';
import HarmonySocket from '../socket/socket';
import { Switch, Route } from 'react-router';
import { BrowserRouter } from 'react-router-dom';
import { Login } from './Entry/Login/Login';
import { Register } from './Entry/Register/Register';

export const harmonySocket = new HarmonySocket();

const App: React.FC = () => {
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
                        <Route exact path='/login'>
                            <Login />
                        </Route>
                        <Route exact path='/register'>
                            <Register />
                        </Route>
                        <Route exact path='/app'></Route>
                    </Switch>
                </BrowserRouter>
            </ThemeProvider>
        </div>
    );
};

export default App;
