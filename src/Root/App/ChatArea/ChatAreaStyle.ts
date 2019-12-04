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
        backgroundColor: theme.palette.type === 'dark' ? theme.palette.grey[900] : theme.palette.grey[200],
        padding: theme.spacing(1)
    },
    channellist: {
        backgroundColor: theme.palette.background.paper
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
        flexFlow: 'column',
        width: '100%',
        height: '100%'
    },
    messages: {
        width: '100%',
        flex: '1 1 auto',
        overflow: 'auto'
    },
    input: {
        width: '100%'
    }
}));
