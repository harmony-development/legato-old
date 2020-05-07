import React from 'react';
import { Button, Dialog, DialogTitle, DialogContent, DialogContentText, DialogActions } from '@material-ui/core';
import { useTranslation } from 'react-i18next';

interface Props {
	open: boolean;
	title: string;
	description: string;
	type: 'confirm' | 'alert';
	onSubmit: () => void;
	onClose: () => void;
}

export const CommonDialog = (props: Props) => {
	const { t } = useTranslation(['common']);

	return (
		<Dialog open={props.open} onClose={props.onClose}>
			<DialogTitle>{props.title}</DialogTitle>
			<DialogContent>
				<DialogContentText>{props.description}</DialogContentText>
			</DialogContent>
			<DialogActions>
				{props.type === 'confirm' ? (
					<>
						<Button color="primary" variant="outlined" onClick={props.onClose} autoFocus>
							{t('common:no')}
						</Button>
						<Button color="primary" variant="contained" onClick={props.onSubmit}>
							{t('common:yes')}
						</Button>
					</>
				) : (
					<Button color="primary" variant="contained" onClick={props.onSubmit}>
						{t('common:ok')}
					</Button>
				)}
			</DialogActions>
		</Dialog>
	);
};
