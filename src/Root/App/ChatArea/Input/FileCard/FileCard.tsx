import React from 'react';
import { Box, IconButton } from '@material-ui/core';
import { Delete } from '@material-ui/icons';

import { useFileCardStyles } from './FileCardStyle';

export const FileCard = (props: { image: string; removeFromQueue: (index: number) => void; index: number }) => {
	const classes = useFileCardStyles();

	return (
		<Box className={classes.root}>
			<div className={classes.overlay}>
				<IconButton onClick={() => props.removeFromQueue(props.index)}>
					<Delete />
				</IconButton>
			</div>
			<img src={props.image} className={classes.thumbnail} alt=""></img>
		</Box>
	);
};
