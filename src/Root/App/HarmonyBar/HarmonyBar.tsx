import React from 'react';
import { AppBar, Toolbar, IconButton, Typography } from '@material-ui/core';
import MenuIcon from '@material-ui/icons/Menu';
import PaletteIcon from '@material-ui/icons/Palette';
import { useHarmonyBarStyles } from './HarmonyBarStyle';
import { useDispatch } from 'react-redux';
import { Actions } from '../../../types/redux';

export const HarmonyBar = () => {
    const classes = useHarmonyBarStyles();
    const dispatch = useDispatch();

    return (
        <AppBar position='static'>
            <Toolbar>
                <IconButton edge='start' color='inherit' className={classes.leftMenuBtn}>
                    <MenuIcon />
                </IconButton>
                <Typography variant='h6' className={classes.title}>
                    Harmony
                </Typography>
                <IconButton edge='end' color='inherit' onClick={() => dispatch({ type: Actions.TOGGLE_THEME_DIALOG })}>
                    <PaletteIcon />
                </IconButton>
            </Toolbar>
        </AppBar>
    );
};
