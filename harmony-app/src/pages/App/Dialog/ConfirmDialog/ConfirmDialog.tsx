import React from 'react';
import { Button, Dialog, DialogTitle, DialogContent, DialogContentText, DialogActions } from '@material-ui/core';

interface Props {
	open: boolean;
	title: string;
	description: string;
	onSubmit: () => void;
	onClose: () => void;
}

export const ConfirmDialog: React.FC<Props> = (props: Props) => {
	return (
		<Dialog open={props.open}>
			<DialogTitle>{props.title}</DialogTitle>
			<DialogContent>
				<DialogContentText>{props.description}</DialogContentText>
			</DialogContent>
			<DialogActions>
				<Button color="primary" variant="outlined" onClick={props.onClose} autoFocus>
					No
				</Button>
				<Button color="primary" variant="contained" onClick={props.onSubmit}>
					Yes
				</Button>
			</DialogActions>
		</Dialog>
	);
};
