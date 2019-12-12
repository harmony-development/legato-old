import React, { useRef, useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { IState, Actions } from '../../../../types/redux';
import { Dialog, TextField, Typography, DialogContent, Button, Grid } from '@material-ui/core';
import { harmonySocket } from '../../../Root';

export const JoinGuild = () => {
    const open = useSelector((state: IState) => state.joinGuildDialog);
    const inputStyle = useSelector((state: IState) => state.theme.inputStyle);
    const [joinErr] = useState<string>('');
    const [createErr] = useState<string>('');
    const joinCodeRef = useRef<HTMLInputElement | null>(null);
    const guildNameRef = useRef<HTMLInputElement | null>(null);
    const dispatch = useDispatch();

    const createGuild = () => {
        if (guildNameRef.current && guildNameRef.current.value) {
            harmonySocket.createGuild(guildNameRef.current.value);
        }
    };

    const joinGuild = () => {
        if (joinCodeRef.current && joinCodeRef.current.value) {
            harmonySocket.joinGuild(joinCodeRef.current.value);
        }
    };

    return (
        <Dialog open={open} onClose={() => dispatch({ type: Actions.TOGGLE_JOIN_GUILD_DIALOG })}>
            <DialogContent>
                <Grid container spacing={1}>
                    <Grid item xs={6}>
                        <div>
                            <Typography variant='h5'>Join Guild</Typography>
                            <TextField label='Join Code' variant={inputStyle as any} fullWidth inputRef={joinCodeRef} />
                            <Typography color='error' variant='body2'>
                                {joinErr || <br />}
                            </Typography>
                            <Button onClick={joinGuild}>Join Guild</Button>
                        </div>
                    </Grid>
                    <Grid item xs={6}>
                        <div>
                            <Typography variant='h5'>Create Guild</Typography>
                            <TextField label='Guild Name' variant={inputStyle as any} fullWidth inputRef={guildNameRef} />
                            <Typography color='error' variant='body2'>
                                {createErr || <br />}
                            </Typography>
                            <Button onClick={createGuild}>Create Guild</Button>
                        </div>
                    </Grid>
                </Grid>
            </DialogContent>
        </Dialog>
    );
};
