import React, { useRef, useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import axios from 'axios';
import { IState, Actions } from '../../../../types/redux';
import { Dialog, DialogContent, AppBar, Toolbar, IconButton, Typography, Button, TextField, Avatar, ButtonBase, Table, TableHead, TableRow, TableCell, TableBody } from '@material-ui/core';
import AddIcon from '@material-ui/icons/Add';
import CloseIcon from '@material-ui/icons/Close';
import { useGuildSettingsStyle } from './GuildSettingsStyle';
import { toast } from 'react-toastify';
import { harmonySocket } from '../../../Root';

export const GuildSettings = () => {
    const open = useSelector((state: IState) => state.guildSettingsDialog);
    const selectedGuild = useSelector((state: IState) => state.selectedGuild);
    const inputStyle = useSelector((state: IState) => state.inputStyle);
    const guilds = useSelector((state: IState) => state.guildList);
    const invites = useSelector((state: IState) => state.invites);
    const dispatch = useDispatch();
    const guildIconUpload = useRef<HTMLInputElement | null>(null);
    const [guildName, setGuildName] = useState<string | undefined>(guilds[selectedGuild] ? guilds[selectedGuild].guildname : undefined);
    const [guildIconFile, setGuildIconFile] = useState<File | null>(null);
    const [guildIcon, setGuildIcon] = useState<string | undefined>(guilds[selectedGuild] ? guilds[selectedGuild].picture : undefined);
    const classes = useGuildSettingsStyle();

    const onSaveChanges = () => {
        if (guilds[selectedGuild]) {
            if (guildIcon !== guilds[selectedGuild].picture && guildIconFile) {
                const guildIconUpload = new FormData();
                guildIconUpload.append('file', guildIconFile);
                axios
                    .post(`http://${window.location.hostname}:2288/api/rest/fileupload`, guildIconUpload, {})
                    .then((res) => {
                        if (res.data) {
                            const uploadID = res.data;
                            harmonySocket.sendGuildPictureUpdate(selectedGuild, `http://${window.location.hostname}:2288/filestore/${uploadID}`);
                        }
                    })
                    .catch(() => {
                        toast.error('Failed to update guild icon');
                    });
            }
            if (guilds[selectedGuild].guildname !== guildName && guildName) {
                harmonySocket.sendGuildNameUpdate(selectedGuild, guildName);
            }
        }
    };

    const onGuildIconSelected = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.currentTarget.files && event.currentTarget.files.length > 0) {
            const file = event.currentTarget.files[0];
            setGuildIconFile(file);
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
                    <IconButton
                        edge='start'
                        color='inherit'
                        onClick={() =>
                            dispatch({
                                type: Actions.TOGGLE_GUILD_SETTINGS_DIALOG
                            })
                        }
                    >
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
                    <TextField
                        label='Guild Name'
                        fullWidth
                        variant={inputStyle as any}
                        className={classes.menuEntry}
                        value={guildName}
                        onChange={(e: React.ChangeEvent<HTMLInputElement>) => setGuildName(e.currentTarget.value)}
                    />
                    <Button variant='contained' color='secondary' className={classes.menuEntry} onClick={onSaveChanges}>
                        Save Changes
                    </Button>
                    <Typography variant='h4' className={classes.menuEntry}>
                        Join Codes
                    </Typography>
                    <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell>Join Code</TableCell>
                                <TableCell>Amount Used</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {Object.keys(invites).map((key) => {
                                return (
                                    <TableRow key={key}>
                                        <TableCell component='th' scope='row'>
                                            {key}
                                        </TableCell>
                                        <TableCell component='th' scope='row'>
                                            {invites[key]}
                                        </TableCell>
                                    </TableRow>
                                );
                            })}
                        </TableBody>
                        <Button fullWidth>
                            <AddIcon />
                        </Button>
                    </Table>
                </div>
            </DialogContent>
        </Dialog>
    );
};
