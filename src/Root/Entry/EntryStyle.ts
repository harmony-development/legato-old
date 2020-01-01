import { makeStyles } from '@material-ui/styles';

export const useEntryStyles = makeStyles(() => ({
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
}));
