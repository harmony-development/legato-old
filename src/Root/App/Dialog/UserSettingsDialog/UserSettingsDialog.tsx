import React, { useState, useRef } from 'react';
import { Dialog, AppBar, Toolbar, IconButton, Typography, DialogContent, TextField, ButtonBase, Avatar, Button } from '@material-ui/core';
import CloseIcon from '@material-ui/icons/Close';
import { useSelector, useDispatch } from 'react-redux';
import { IState, Actions } from '../../../../types/redux';
import { useUserSettingsStyle } from './UserSettingsStyle';

export const UserSettingsDialog = () => {
    const [open, inputStyle] = useSelector((state: IState) => [state.userSettingsDialog, state.userSettingsDialog]);
    const userAvatarUpload = useRef<HTMLInputElement | null>(null);
    const [username, setUsername] = useState<string>('');
    const dispatch = useDispatch();
    const classes = useUserSettingsStyle();

    return (
        <Dialog open={open} onClose={() => dispatch({ type: Actions.TOGGLE_USER_SETTINGS_DIALOG })} fullScreen>
            <AppBar style={{ position: 'relative' }}>
                <Toolbar>
                    <IconButton edge='start' onClick={() => dispatch({ type: Actions.TOGGLE_USER_SETTINGS_DIALOG })}>
                        <CloseIcon />
                    </IconButton>
                    <Typography variant='h6'>User Settings</Typography>
                </Toolbar>
            </AppBar>
            <DialogContent>
                <div style={{ width: '33%' }}>
                    <input type='file' id='file' multiple ref={userAvatarUpload} style={{ display: 'none' }} />
                    <ButtonBase
                        style={{ borderRadius: '50%' }}
                        onClick={() => {
                            if (userAvatarUpload.current) {
                                userAvatarUpload.current.click();
                            }
                        }}
                    >
                        <Avatar className={classes.guildIcon}></Avatar>
                    </ButtonBase>
                    <TextField
                        label='Username'
                        fullWidth
                        variant={inputStyle as any}
                        className={classes.menuEntry}
                        value={username}
                        onChange={(e: React.ChangeEvent<HTMLInputElement>) => setUsername(e.currentTarget.value)}
                    />
                    <Button variant='contained' color='secondary' className={classes.menuEntry}>
                        Save Changes
                    </Button>
                </div>
            </DialogContent>
        </Dialog>
    );
};
