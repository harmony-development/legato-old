import React, { useEffect } from 'react';
import NavBar from './NavBar/NavBar';
import ChangeNameDialog from './ProfileSettingsDialog/ProfileSettingsDialog';
import { useStyles } from './styles';
import Chat from './Chat/Chat';
import { Events } from '../../socket/socket';
import { useHistory } from 'react-router';
import { socketServer } from '../Root';

const App: React.FC<{}> = () => {
  const classes = useStyles();
  const history = useHistory();

  useEffect(() => {
    socketServer.connection.on(Events.INVALIDATE_SESSION, () => {
      localStorage.removeItem('token');
      history.push('/');
    });

    return (): void => {
      socketServer.connection.removeEventListener(Events.INVALIDATE_SESSION);
    };
  });

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
