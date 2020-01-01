import React from 'react';
import { Paper, Tabs, Tab } from '@material-ui/core';

import { useEntryStyles } from './EntryStyle';
import { Login } from './Login/Login';
import { Register } from './Register/Register';

export const Entry = () => {
	const classes = useEntryStyles();
	const [tabIDX, setTabIDX] = React.useState(0);

	return (
		<div className={classes.root}>
			<Paper className={classes.form}>
				<Tabs
					value={tabIDX}
					onChange={(_event: any, newValue: number): void => setTabIDX(newValue)}
					variant="fullWidth"
					indicatorColor={'primary'}
				>
					<Tab label="Login" id="form-tab-0" />
					<Tab label="Register" id="form-tab-1" />
				</Tabs>
				{tabIDX === 0 ? <Login /> : <Register />}
			</Paper>
		</div>
	);
};
