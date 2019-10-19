import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { IAppState, IShowChangeNameDialog } from '../../../store/types';
import {toggleChangeNameDialog, updateUser} from '../../../store/actions/AppActions';
import { Dialog, DialogTitle, DialogContent, TextField, DialogActions, Button } from '@material-ui/core';
import UserIcon from './UserIcon/UserIcon';

const ChangeNameDialog: React.FC<{}> = () => {
  const dispatch = useDispatch();
  const [open, user] = useSelector((state: IAppState) => [state.nameDialog, state.user]);
  const [draftName, setDraftName] = useState(user.name);
  const [draftIcon, setDraftIcon] = useState(user.icon);

  const setProfileAndClose = (): void => {
    dispatch(updateUser({
        name: draftName,
        icon: draftIcon
    }));
    dispatch(toggleChangeNameDialog());
  };

  return (
    <Dialog open={open} onClose={(): IShowChangeNameDialog => dispatch(toggleChangeNameDialog())}>
      <DialogTitle>Change Username</DialogTitle>
      <DialogContent>
        <UserIcon setIcon={setDraftIcon} icon={draftIcon} />
      </DialogContent>
      <DialogContent>
        <TextField label='Username' value={draftName} onChange={(e): void => setDraftName(e.target.value)}/>
      </DialogContent>
      <DialogActions>
        <Button onClick={setProfileAndClose} color='primary' autoFocus>
          Apply
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default ChangeNameDialog;
