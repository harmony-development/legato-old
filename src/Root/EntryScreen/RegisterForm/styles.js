import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useStyles = makeStyles((theme) => ({
    root: {
        paddingLeft: theme.spacing(1),
        paddingRight: theme.spacing(1),
        paddingBottom: theme.spacing(1)
    },
    submitButton: {
        position: 'absolute',
        bottom: theme.spacing(1),
        left: theme.spacing(1)
    }
}));
