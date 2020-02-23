import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useRegisterStyles = makeStyles((theme: Theme) => ({
    root: {
        paddingLeft: theme.spacing(1),
        paddingRight: theme.spacing(1),
        paddingTop: theme.spacing(1),
        paddingBottom: theme.spacing(1)
    },
    submitBtn: {
        position: 'absolute',
        bottom: theme.spacing(1),
        left: theme.spacing(1)
    }
}));
