import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { CircularProgress } from '@material-ui/core';

const rootElementStyle = {
	width: '100vw',
	height: '100vh',
	display: 'flex',
	justifyContent: 'center',
	alignItems: 'center',
};

export const Loading = () => {
	return (
		<div style={rootElementStyle}>
			<CircularProgress />
		</div>
	);
};
