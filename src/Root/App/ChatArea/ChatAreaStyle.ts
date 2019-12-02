import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useChatAreaStyles = makeStyles((theme: Theme) => ({
    root: {
        flex: 1,
        display: 'flex',
        flexDirection: 'column',
        flexGrow: 1,
        width: '100%',
        height: '100%'
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
