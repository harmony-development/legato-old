import React, { useState, useEffect } from 'react';
import {
	Dialog,
	DialogTitle,
	DialogContent,
	DialogActions,
	Button,
	Color,
	FormControlLabel,
	FormControl,
	FormLabel,
	RadioGroup,
	Radio,
	Typography,
	Select,
	MenuItem,
} from '@material-ui/core';
import { useDispatch, useSelector } from 'react-redux';

import { IState } from '../../../types/redux';
import { AppDispatch } from '../../../redux/store';
import { SetInputStyle, SetPrimary, SetSecondary, ToggleThemeDialog, InvertTheme } from '../../../redux/AppReducer';

import { ColorPicker } from './ColorPicker';

export const ThemeDialog = () => {
	const dispatch = useDispatch<AppDispatch>();
	const [open, themeType, globalPrimary, globalSecondary, inputStyle] = useSelector((state: IState) => [
		state.themeDialog,
		state.theme.type,
		state.theme.primary,
		state.theme.secondary,
		state.theme.inputStyle,
	]);
	const [primary, setPrimary] = useState<Color>(globalPrimary);
	const [secondary, setSecondary] = useState<Color>(globalSecondary);

	const changeInputStyle = (ev: React.ChangeEvent<{ value: unknown }>) => {
		if (ev.target.value === 'standard' || ev.target.value === 'outlined' || ev.target.value === 'filled') {
			dispatch(SetInputStyle(ev.target.value));
		}
	};

	const handleInvertTheme = () => {
		dispatch(InvertTheme());
	};

	useEffect(() => {
		localStorage.setItem('inputstyle', inputStyle);
	}, [inputStyle]);
	useEffect(() => {
		localStorage.setItem('themetype', themeType);
	}, [themeType]);
	useEffect(() => {
		dispatch(SetPrimary(primary));
		localStorage.setItem('primary', JSON.stringify(primary));
	}, [primary]);
	useEffect(() => {
		dispatch(SetSecondary(secondary));
		localStorage.setItem('secondary', JSON.stringify(secondary));
	}, [secondary]);

	return (
		<Dialog open={open} onClose={() => dispatch(ToggleThemeDialog())}>
			<DialogTitle>Customize Theme</DialogTitle>
			<DialogContent>
				<FormControl component="fieldset">
					<FormLabel component="legend">Theme Type</FormLabel>
					<RadioGroup value={themeType} row onChange={handleInvertTheme}>
						<FormControlLabel value="light" control={<Radio color="secondary" />} label="Light" labelPlacement="end" />
						<FormControlLabel value="dark" control={<Radio color="secondary" />} label="Dark" labelPlacement="end" />
					</RadioGroup>
				</FormControl>
				<ColorPicker color={primary} setColor={setPrimary} label={'Primary Color'} />
				<ColorPicker color={secondary} setColor={setSecondary} label={'Secondary Color'} />
				<Typography>Text Input Style</Typography>
				<Select value={inputStyle || 'standard'} onChange={changeInputStyle} variant={inputStyle as any} fullWidth>
					<MenuItem value={'standard'}>Standard</MenuItem>
					<MenuItem value={'filled'}>Filled</MenuItem>
					<MenuItem value={'outlined'}>Outlined</MenuItem>
				</Select>
			</DialogContent>
			<DialogActions>
				<Button color="primary" onClick={() => dispatch(ToggleThemeDialog())}>
					Close
				</Button>
			</DialogActions>
		</Dialog>
	);
};
