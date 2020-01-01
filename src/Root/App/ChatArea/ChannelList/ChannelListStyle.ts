import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

import { HarmonyDark } from '../../HarmonyColor';

export const useChannelListStyle = makeStyles((theme: Theme) => ({
	selectedChannel: {
		backgroundColor: theme.palette.type === 'dark' ? HarmonyDark[400] : undefined,
	},
	nested: {
		paddingLeft: theme.spacing(4),
	},
	newChannelInput: {
		marginLeft: theme.spacing(2),
		marginRight: theme.spacing(2),
	},
}));
