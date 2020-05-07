import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useGuildSettingsStyle = makeStyles((theme: Theme) => ({
	menuEntry: {
		marginTop: theme.spacing(2),
	},
	clipboardbtn: {
		marginLeft: theme.spacing(3),
	},
}));
