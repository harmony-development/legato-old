import React, { useEffect, useState, useCallback } from 'react';
import { useParams, useHistory } from 'react-router';
import { Typography, Button } from '@material-ui/core';

import { harmonySocket } from '../../Root';

import { useInviteHandlerStyle } from './InviteHandlerStyle';

export const InviteHandler = () => {
	const { invitecode } = useParams();
	const history = useHistory();
	const [errorMessage, setErrorMessage] = useState<string | null>(null);
	const classes = useInviteHandlerStyle();

	const handleJoinGuild = useCallback(
		(raw: any) => {
			if (!raw['message']) {
				setErrorMessage(null);
				history.push('/app');
			} else {
				setErrorMessage(raw['message']);
			}
			harmonySocket.events.removeCurrentListener();
		},
		[history]
	);

	useEffect(() => {
		harmonySocket.events.addListener('joinguild', handleJoinGuild);
	}, [handleJoinGuild]);

	useEffect(() => {
		harmonySocket.events.addListener('error', (raw: any) => {
			setErrorMessage(raw['message']);
		});
	}, []);

	useEffect(() => {
		if (invitecode) {
			if (harmonySocket.conn.readyState === WebSocket.OPEN) {
				harmonySocket.joinGuild(invitecode);
			} else {
				harmonySocket.events.addListener('open', () => {
					harmonySocket.joinGuild(invitecode);
				});
			}
		}
	}, [invitecode]);

	return (
		<div>
			{errorMessage ? (
				<div className={classes.errorRoot}>
					<Typography variant="h1" align="center" className={classes.errorMsg}>
						404
						<br />
						{errorMessage}
					</Typography>
					<Button variant="contained" color="secondary" className={classes.errorBtn} onClick={() => history.push('/')}>
						Return To Login
					</Button>
				</div>
			) : (
				undefined
			)}
		</div>
	);
};
