import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useUserSettingsStyle = makeStyles((theme: Theme) => ({
    guildIcon: {
        width: '100px',
        height: '100px'
    },
    menuEntry: {
        marginTop: theme.spacing(2)
    },
    clipboardbtn: {
        marginLeft: theme.spacing(3)
    }
}));
