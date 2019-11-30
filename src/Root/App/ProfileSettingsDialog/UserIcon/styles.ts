import { makeStyles, Theme } from '@material-ui/core';

export const useStyles = makeStyles((theme: Theme) => ({
    iconRoot: {
        display: 'flex'
    },
    changeIconButton: {
        marginLeft: theme.spacing(1)
    }
}));
