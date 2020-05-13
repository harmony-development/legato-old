import React, { useEffect } from 'react';
import { Drawer, List, ListItem, Typography, Theme } from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import { useTranslation } from 'react-i18next';

import { AuthAPI } from '../../../api/Auth';

const instanceListStyles = makeStyles((theme: Theme) => ({
	title: {
		padding: theme.spacing(1),
	},
	list: {
		width: 250,
	},
}));

export const InstanceList = React.memo(() => {
	const classes = instanceListStyles();
	const { t } = useTranslation('instancelist');

	useEffect(() => {
		(async () => {
			try {
				const resp = await AuthAPI.listServers();
				console.log(resp);
			} catch (e) {
				console.log(e);
			}
		})();
	}, []);

	return (
		<Drawer open={true}>
			<Typography variant="h4" className={classes.title}>
				{t('instancelist:title')}
			</Typography>
			<List className={classes.list}>
				<ListItem button>hi</ListItem>
			</List>
		</Drawer>
	);
});
