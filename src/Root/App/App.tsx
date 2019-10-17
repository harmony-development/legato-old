import React from 'react';
import NavBar from './NavBar/NavBar';
import Chat from './Chat/Chat';
import { useStyles } from './styles';
import { createMuiTheme, CssBaseline } from '@material-ui/core';
import { useSelector } from 'react-redux';
import { IAppState } from '../../store/types';
import { ThemeProvider } from '@material-ui/styles';
import ChangeNameDialog from './ChangeNameDialog/ChangeNameDialog';

const App = () => {
  const classes = useStyles();
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
        <CssBaseline />
        <NavBar />
        <div className={classes.navbarSpacer} />
        <ChangeNameDialog />
        <Chat />
      </ThemeProvider>
    </div>
  );
};

export default App;
