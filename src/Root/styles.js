import { makeStyles } from '@material-ui/core';
import { Theme } from '@material-ui/core';

export const useStyles = makeStyles((theme) => ({
    '@global': {
        '::-webkit-scrollbar': {
            width: '10px'
        },
        '::-webkit-scrollbar-thumb:hover': {
            backgroundColor: theme.palette.type === 'light' ? 'rgb(150, 150, 150)' : 'rgb(100, 100, 100)'
        },
        '::-webkit-scrollbar-track': {
            backgroundColor: theme.palette.type === 'light' ? 'rgb(245, 245, 245)' : 'rgb(46, 46, 46)'
        },
        '::-webkit-scrollbar-thumb': {
            backgroundColor: theme.palette.type === 'light' ? 'rgb(200, 200, 200)' : 'rgb(64, 64, 64)'
        },
        '::-webkit-scrollbar-corner': {
            backgroundColor: theme.palette.type === 'light' ? 'rgb(200, 200, 200)' : 'rgb(64, 64, 64)'
        },
        '*': {
            scrollbarColor: `${theme.palette.type === 'light' ? 'rgb(200, 200, 200)' : 'rgb(64, 64, 64)'} ${theme.palette.type === 'light' ? 'rgb(245, 245, 245)' : 'rgb(46, 46, 46)'}`
        }
    }
}));
