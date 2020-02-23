import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useInviteHandlerStyle = makeStyles((theme: Theme) => ({
    errorRoot: {
        textAlign: 'center'
    },
    errorMsg: {
        paddingTop: theme.spacing(10)
    },
    errorBtn: {
        marginTop: theme.spacing(2)
    }
}));
