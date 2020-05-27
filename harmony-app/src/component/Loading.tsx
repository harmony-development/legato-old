import React from 'react';
import { CircularProgress } from '@material-ui/core';

const rootElementStyle = {
	width: '100vw',
	height: '100vh',
	display: 'flex',
	justifyContent: 'center',
	alignItems: 'center',
};

export const Loading = React.memo(() => {
	return (
		<div style={rootElementStyle}>
			<CircularProgress />
		</div>
	);
});
