import React from 'react';
import NavBar from './NavBar/NavBar';
import ChangeNameDialog from './ChangeNameDialog/ChangeNameDialog';
import { useStyles } from './styles';
import Chat from './Chat/Chat';

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
