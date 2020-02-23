import React from 'react';
import { AppBar, Toolbar, IconButton, Typography } from '@material-ui/core';
import MenuIcon from '@material-ui/icons/Menu';
import PaletteIcon from '@material-ui/icons/Palette';
import UserIcon from '@material-ui/icons/AccountCircle';
import { useDispatch } from 'react-redux';

import { AppDispatch } from '../../../redux/store';
import { ToggleUserSettingsDialog, ToggleThemeDialog } from '../../../redux/AppReducer';

import { useHarmonyBarStyles } from './HarmonyBarStyle';

export const HarmonyBar = () => {
	const dispatch = useDispatch<AppDispatch>();
	const classes = useHarmonyBarStyles();

	return (
		<AppBar position="absolute">
			<Toolbar>
				<IconButton edge="start" color="inherit" className={classes.leftMenuBtn}>
					<MenuIcon />
				</IconButton>
				<Typography variant="h6" className={classes.title}>
					Harmony
				</Typography>
				<IconButton
					edge="end"
					color="inherit"
					onClick={() => {
						dispatch(ToggleThemeDialog());
					}}
				>
					<PaletteIcon />
				</IconButton>
				<IconButton edge="end" color="inherit" onClick={() => dispatch(ToggleUserSettingsDialog())}>
					<UserIcon />
				</IconButton>
			</Toolbar>
		</AppBar>
	);
};
