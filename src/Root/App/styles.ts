import { makeStyles } from '@material-ui/core';
import { Theme } from '@material-ui/core';

export const useStyles = makeStyles((theme: Theme) => ({
  navbarSpacer: theme.mixins.toolbar,
  '@global': {
    '::-webkit-scrollbar': {
      width: '10px'
    },
    '::-webkit-scrollbar-track': {
      backgroundColor: theme.palette.type === 'dark' ? 'rgb(46, 46, 46)' : 'rgb(245, 245, 245)'
    },
    '::-webkit-scrollbar-thumb': {
      backgroundColor: theme.palette.type === 'dark' ? 'rgb(64, 64, 64)' : 'rgb(200, 200, 200)'
    }
  }
}));
