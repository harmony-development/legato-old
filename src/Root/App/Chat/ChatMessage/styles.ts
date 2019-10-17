import { makeStyles } from '@material-ui/core';
import { Theme } from '@material-ui/core';

export const useStyles = makeStyles((theme: Theme) => ({
  messageContainer: {
    minHeight: '60px',
    paddingLeft: theme.spacing(1),
    paddingRight: theme.spacing(1),
    backgroundColor: 'rgb(240, 240, 240)'
  },
  inline: {
    display: 'inline'
  },
  messageLight: {
    backgroundColor: 'rgb(133, 133, 133, 0.1)'
  },
  messageText: {
    whiteSpace: 'pre-line'
  },
  userText: {
    marginRight: theme.spacing(1)
  }
}));
