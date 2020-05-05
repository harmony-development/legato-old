import React from 'react';
import { Paper, Tabs, Tab } from '@material-ui/core';
import { useTranslation } from 'react-i18next';

import { useEntryStyles } from './EntryStyle';
import { Login } from './Login/Login';
import { Register } from './Register/Register';

export const Entry = () => {
	const classes = useEntryStyles();
	const [tabIDX, setTabIDX] = React.useState(0);
	const { t } = useTranslation('entry');

	return (
		<div className={classes.root}>
			<Paper className={classes.form} elevation={5}>
				<Tabs
					value={tabIDX}
					onChange={(_event, newValue: number): void => setTabIDX(newValue)}
					variant="fullWidth"
					indicatorColor={'primary'}
				>
					<Tab label={t('entry:login')} />
					<Tab label={t('entry:register')} />
				</Tabs>
				{tabIDX === 0 ? <Login /> : <Register />}
			</Paper>
		</div>
	);
};
