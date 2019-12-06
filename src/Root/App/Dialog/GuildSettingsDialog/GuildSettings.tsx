import React, { useRef, useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { IState, Actions } from '../../../../types/redux';
import { Dialog, DialogContent, AppBar, Toolbar, IconButton, Typography, Button, TextField, Avatar, ButtonBase } from '@material-ui/core';
import CloseIcon from '@material-ui/icons/Close';
import { useGuildSettingsStyle } from './GuildSettingsStyle';
import axios from 'axios';
import { harmonySocket } from '../../../Root';

export const GuildSettings = () => {
    const open = useSelector((state: IState) => state.guildSettingsDialog);
    const selectedGuild = useSelector((state: IState) => state.selectedGuild);
    const guilds = useSelector((state: IState) => state.guildList);
    const dispatch = useDispatch();
    const guildIconUpload = useRef<HTMLInputElement | null>(null);
    const [guildName, setGuildName] = useState<string | undefined>(guilds[selectedGuild] ? guilds[selectedGuild].guildname : undefined);
    const [guildIcon, setGuildIcon] = useState<string | undefined>(guilds[selectedGuild] ? guilds[selectedGuild].picture : undefined);
    const classes = useGuildSettingsStyle();

    const onSaveChanges = () => {};

    const onGuildIconSelected = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.currentTarget.files && event.currentTarget.files.length > 0) {
            const file = event.currentTarget.files[0];
            if (file.type.startsWith('image/') && file.size < 33554432) {
                const fileReader = new FileReader();
                fileReader.readAsDataURL(file);
                fileReader.addEventListener('load', () => {
                    if (typeof fileReader.result === 'string') {
                        setGuildIcon(fileReader.result);
                    }
                });
            }
        }
    };

    return (
        <Dialog open={open} onClose={() => dispatch({ type: Actions.TOGGLE_GUILD_SETTINGS_DIALOG })} fullScreen>
            <AppBar style={{ position: 'relative' }}>
                <Toolbar>
                    <IconButton edge='start' color='inherit' onClick={() => dispatch({ type: Actions.TOGGLE_GUILD_SETTINGS_DIALOG })}>
                        <CloseIcon />
                    </IconButton>
                    <Typography variant='h6'>Guild Settings</Typography>
                </Toolbar>
            </AppBar>
            <DialogContent>
                <div style={{ width: '33%' }}>
                    <input type='file' id='file' multiple ref={guildIconUpload} style={{ display: 'none' }} onChange={onGuildIconSelected} />
                    <ButtonBase
                        style={{ borderRadius: '50%' }}
                        onClick={() => {
                            if (guildIconUpload.current) {
                                guildIconUpload.current.click();
                            }
                        }}
                    >
                        <Avatar className={classes.guildIcon} src={guildIcon}></Avatar>
                    </ButtonBase>
                    <TextField label='Guild Name' fullWidth className={classes.menuEntry} value={guildName} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setGuildName(e.currentTarget.value)} />
                    <Button variant='contained' color='secondary' className={classes.menuEntry} onClick={onSaveChanges}>
                        Save Changes
                    </Button>
                </div>
            </DialogContent>
        </Dialog>
    );
};
