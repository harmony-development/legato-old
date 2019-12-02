import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useChatAreaStyles = makeStyles((theme: Theme) => ({
    root: {
        flex: 1,
        display: 'flex',
        flexDirection: 'row',
        flexGrow: 1,
        width: '100%',
        height: '100%'
    },
    guildlist: {
        backgroundColor: theme.palette.grey[900],
        padding: theme.spacing(1)
    },
    guildiconroot: {
        borderRadius: '64px'
    },
    guildicon: {
        width: '64px',
        height: '64px'
    },
    chatArea: {
        display: 'flex',
        flexDirection: 'column',
        width: '100%'
    },
    messages: {
        width: '100%',
        flex: 1,
        flexGrow: 1
    },
    input: {
        width: '100%'
    }
}));
