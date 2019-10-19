import { makeStyles, Theme } from '@material-ui/core';

export const useStyles = makeStyles((theme: Theme) => ({
  thumbnail: {
    height: '100px'
  },
  root: {
    position: 'relative',
    '&:hover $overlay': {
      opacity: 100,
      visibility: 'visible'
    }
  },
  overlay: {
    opacity: 0,
    visibility: 'hidden',
    display: 'flex',
    position: 'absolute',
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'rgba(0, 0, 0, 0.4)',
    top: 0,
    left: 0,
    height: '100%',
    width: '100%',
    transition: 'all 0.2s ease-in'
  },
  deleteButton: {}
}));
