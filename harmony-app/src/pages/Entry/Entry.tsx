import React from 'react';
import { Paper, Tabs, Tab, makeStyles, Container } from '@material-ui/core';
import { useTranslation } from 'react-i18next';

import { Login } from './Login';
import { Register } from './Register';

const entryStyles = makeStyles({
	form: {
		width: '60%',
		height: '60%',
		position: 'relative',
	},
	root: {
		width: '100vw',
		height: '100vh',
		display: 'flex',
		alignItems: 'center',
		justifyContent: 'center',
	},
});

export const Entry = () => {
	const classes = entryStyles();
	const [tabIDX, setTabIDX] = React.useState(0);
	const { t } = useTranslation('entry');

	return (
		<div className={classes.root}>
			<Container maxWidth="sm">
				<Paper elevation={5}>
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
			</Container>
		</div>
	);
};
