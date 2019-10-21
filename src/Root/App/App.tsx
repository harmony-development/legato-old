import React from 'react';
import NavBar from './NavBar/NavBar';
import ChangeNameDialog from './ChangeNameDialog/ChangeNameDialog';
import { Chat } from '@material-ui/icons';
import { useStyles } from './styles';

const App: React.FC<{}> = () => {
  const classes = useStyles();

  return (
    <>
      <NavBar />
      <div className={classes.navbarSpacer} />
      <ChangeNameDialog />
      <Chat />
    </>
  );
};

export default App;
