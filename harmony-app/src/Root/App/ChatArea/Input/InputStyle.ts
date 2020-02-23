import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useInputStyles = makeStyles(() => ({
	inputRoot: {
		display: 'flex',
		flexDirection: 'row',
	},
	fileQueue: {
		overflowX: 'scroll',
		height: 'auto',
		display: 'flex',
	},
}));
