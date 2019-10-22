import React from 'react';
import { Menu, InvertColors, AccountCircle } from '@material-ui/icons';
import { AppBar, Toolbar, IconButton, Typography } from '@material-ui/core';
import { useStyles } from './styles';
import { useDispatch } from 'react-redux';
import { invertTheme, toggleProfileSettingsDialog } from '../../../store/actions/AppActions';
import { IInvertTheme, IToggleProfileSettingsDialog } from '../../../store/types';

const NavBar: React.FC<{}> = () => {
  const classes = useStyles();
  const dispatch = useDispatch();

  return (
    <AppBar>
      <Toolbar>
        <IconButton edge='start' className={classes.drawerButton} color='inherit' aria-label='menu'>
          <Menu />
        </IconButton>
        <Typography variant='h6' className={classes.title}>
          #general
        </Typography>
        <IconButton onClick={(): IToggleProfileSettingsDialog => dispatch(toggleProfileSettingsDialog())}>
          <AccountCircle />
        </IconButton>
        <IconButton onClick={(): IInvertTheme => dispatch(invertTheme())}>
          <InvertColors />
        </IconButton>
      </Toolbar>
    </AppBar>
  );
};

export default NavBar;
