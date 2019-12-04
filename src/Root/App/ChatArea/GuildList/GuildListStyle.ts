import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useGuildListStyle = makeStyles((theme: Theme) => ({
    guildlist: {
        width: '68px'
    },
    guildiconroot: {
        borderRadius: '64px',
        marginTop: theme.spacing(1)
    },
    guildicon: {
        width: '64px',
        height: '64px',
        borderRadius: '64px',
        backgroundColor: theme.palette.type === 'dark' ? theme.palette.grey[800] : theme.palette.grey[400]
    }
}));
