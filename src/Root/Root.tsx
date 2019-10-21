/**
 * This file is intended for adding things such as Redux Providers and other things.
 */
import React from 'react';
import { Provider } from 'react-redux';
import { store } from '../store/store';
import { useStyles } from './styles';
import { createMuiTheme, CssBaseline } from '@material-ui/core';
import { useSelector } from 'react-redux';
import { IAppState } from '../store/types';
import { ThemeProvider } from '@material-ui/styles';
import { Switch, Route, BrowserRouter as Router } from 'react-router-dom';
import EntryScreen from './EntryScreen/EntryScreen';
import App from './App/App';

const Theme: React.FC<{}> = () => {
  const classes = useStyles();

  return <></>;
};

const Root: React.FC<{}> = () => {
  const { type, primary, secondary } = useSelector((state: IAppState) => state.theme);
  const theme = createMuiTheme({
    palette: {
      type,
      primary,
      secondary
    }
  });

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
