import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { IAppState } from '../../../store/types';
import { toggleChangeNameDialog, changeName } from '../../../store/actions/AppActions';
import { Dialog, DialogTitle, DialogContent, DialogContentText, Input, TextField, DialogActions, Button } from '@material-ui/core';

const ChangeNameDialog = () => {
  const dispatch = useDispatch();
  const [open, name] = useSelector((state: IAppState) => [state.nameDialog, state.name]);
  const [draftName, setDraftName] = useState(name);

  const setNameAndClose = () => {
    dispatch(changeName(draftName));
    dispatch(toggleChangeNameDialog());
  };

  return (
    <Dialog open={open} onClose={() => dispatch(toggleChangeNameDialog())}>
      <DialogTitle>Change Username</DialogTitle>
      <DialogContent>
        <TextField label='Username' value={draftName} onChange={(e) => setDraftName(e.target.value)}></TextField>
      </DialogContent>
      <DialogActions>
        <Button onClick={setNameAndClose} color='primary' autoFocus>
          Apply
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default ChangeNameDialog;
