import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useGuildListStyle = makeStyles((theme: Theme) => ({
    guildlist: {
        width: '68px'
    },
    guildiconroot: {
        boxSizing: 'border-box',
        borderRadius: '50%',
        marginTop: theme.spacing(1),
        border: `2px solid transparent`
    },
    selectedguildicon: {
        border: `2px solid ${theme.palette.primary.light}`
    },
    guildicon: {
        width: '64px',
        height: '64px',
        borderRadius: '64px',
        backgroundColor: theme.palette.type === 'dark' ? theme.palette.grey[800] : theme.palette.grey[400]
    },
    joinGuild: {
        marginTop: theme.spacing(1),
        width: '64px',
        height: '64px',
        borderRadius: '64px',
        backgroundColor: theme.palette.secondary.dark
    }
}));
