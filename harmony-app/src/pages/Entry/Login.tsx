import React, { useState } from 'react';
import { TextField, Typography, Button, makeStyles, Theme } from '@material-ui/core';
import { useHistory } from 'react-router';
import { useTranslation } from 'react-i18next';

import { AuthAPI } from '../../api/Auth';
import { useDialog } from '../../component/Dialog/CommonDialogContext';

const loginStyles = makeStyles((theme: Theme) => ({
	root: {
		paddingLeft: theme.spacing(1),
		paddingRight: theme.spacing(1),
		paddingTop: theme.spacing(1),
		paddingBottom: theme.spacing(1),
	},
	submitBtn: {
		marginTop: theme.spacing(2),
	},
}));

export const Login = () => {
	const history = useHistory();
	const classes = loginStyles();
	const dialog = useDialog();
	const { t } = useTranslation('entry');
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');

	const onEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => setEmail(e.currentTarget.value);
	const onPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => setPassword(e.currentTarget.value);

	const onSubmit = async (e: React.ChangeEvent<EventTarget>) => {
		e.preventDefault();
		try {
			const resp = await AuthAPI.login(email, password);
			localStorage.setItem('authsession', resp.session);
			history.push('/app');
		} catch (err) {
			dialog({
				type: 'alert',
				title: t('common:error'),
				description: err.message,
			});
			console.error('error logging in', err);
		}
	};

	return (
		<div className={classes.root}>
			<form onSubmit={onSubmit}>
				<TextField
					label={t('entry:email')}
					onChange={onEmailChange}
					type="email"
					name="email"
					autoComplete="email"
					margin="normal"
					fullWidth
				/>
				<TextField
					label={t('entry:password')}
					onChange={onPasswordChange}
					type="password"
					name="password"
					margin="normal"
					fullWidth
				/>
				<Button variant="contained" color="secondary" className={classes.submitBtn} type="submit" fullWidth>
					{t('entry:log-in')}
				</Button>
			</form>
		</div>
	);
};
