import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useAppStyles = makeStyles((theme: Theme) => ({
	root: {
		display: 'flex',
		height: '100%',
		flexDirection: 'column',
	},
	leftMenuBtn: {
		marginRight: theme.spacing(1),
	},
	title: {
		flexGrow: 1,
	},
	navFill: {
		...theme.mixins.toolbar,
	},
}));
