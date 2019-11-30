import { makeStyles } from '@material-ui/core';
import { Theme } from '@material-ui/core';

export const useStyles = makeStyles((theme: Theme) => ({
    drawerButton: {
        marginRight: theme.spacing(2)
    },
    title: {
        flexGrow: 1
    }
}));
