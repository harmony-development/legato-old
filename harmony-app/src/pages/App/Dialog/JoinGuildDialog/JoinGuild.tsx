import React, { useRef, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Dialog, TextField, Typography, DialogContent, Button, Grid } from '@material-ui/core';

import { IState } from '../../../../types/redux';
import { harmonySocket } from '../../../../Root';
import { ToggleGuildDialog } from '../../../../redux/AppReducer';
import { AppDispatch } from '../../../../redux/store';

export const JoinGuild = () => {
	const dispatch = useDispatch<AppDispatch>();
	const [open, inputStyle] = useSelector((state: IState) => [state.guildDialog, state.theme.inputStyle]);
	const [joinErr] = useState<string>('');
	const [createErr] = useState<string>('');
	const joinCodeRef = useRef<HTMLInputElement | null>(null);
	const guildNameRef = useRef<HTMLInputElement | null>(null);

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
		<Dialog open={open} onClose={() => dispatch(ToggleGuildDialog())}>
			<DialogContent>
				<Grid container spacing={1}>
					<Grid item xs={6}>
						<div>
							<Typography variant="h5">Join Guild</Typography>
							<TextField label="Join Code" variant={inputStyle as any} fullWidth inputRef={joinCodeRef} />
							<Typography color="error" variant="body2">
								{joinErr || <br />}
							</Typography>
							<Button onClick={joinGuild}>Join Guild</Button>
						</div>
					</Grid>
					<Grid item xs={6}>
						<div>
							<Typography variant="h5">Create Guild</Typography>
							<TextField label="Guild Name" variant={inputStyle as any} fullWidth inputRef={guildNameRef} />
							<Typography color="error" variant="body2">
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
