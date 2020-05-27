import { CircularProgress, Paper, Typography } from '@material-ui/core';
import React from 'react';

export const DisconnectedStatus = () => {
	return (
		<Paper>
			<div style={{ display: 'flex', alignItems: 'center' }}>
				<CircularProgress />
				<Typography style={{ marginLeft: '16px' }}>Reconnecting To Server...</Typography>
			</div>
		</Paper>
	);
};
