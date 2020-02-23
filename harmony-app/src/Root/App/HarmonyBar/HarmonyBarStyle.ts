import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useHarmonyBarStyles = makeStyles((theme: Theme) => ({
    leftMenuBtn: {
        marginRight: theme.spacing(1)
    },
    title: {
        flexGrow: 1
    }
}));
