import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useEntryStyles = makeStyles((theme: Theme) => ({
    form: {
        width: '60%',
        height: '60%',
        position: 'relative'
    },
    root: {
        width: '100vw',
        height: '100vh',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center'
    }
}));
