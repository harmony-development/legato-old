import { makeStyles } from '@material-ui/core';
import { Theme } from '@material-ui/core';

export const useStyles = makeStyles((theme: Theme) => ({
    container: {
        display: 'flex',
        overflow: 'auto',
        flex: 1,
        flexFlow: 'column',
        flexDirection: 'column'
    },
    messageBoxContainer: {
        flex: '0 0 auto',
        display: 'flex',
        paddingRight: theme.spacing(1)
    },
    fileQueue: {
        overflowX: 'scroll',
        height: 'auto'
    },
    valign: {
        display: 'flex',
        alignItems: 'center'
    },
    chatBoxContainer: {
        flex: '1 1 auto',
        overflow: 'auto'
    },
    chatBox: {
        flex: 1
    },
    messageBox: {
        width: '100%'
    }
}));
