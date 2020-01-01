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
import { useSelector } from 'react-redux';
import { orange, red } from '@material-ui/core/colors';

import { IState } from '../../../types/redux';
import { store } from '../../../redux/store';
import { SetInputStyle, SetPrimary, SetSecondary, ToggleThemeDialog, InvertTheme } from '../../../redux/AppReducer';

import { ColorPicker } from './ColorPicker';

export const ThemeDialog = () => {
	const [open, themeType, inputStyle] = useSelector((state: IState) => [
		state.themeDialog,
		state.theme.type,
		state.theme.inputStyle,
	]);
	const [primary, setPrimary] = useState<Color>(red);
	const [secondary, setSecondary] = useState<Color>(orange);

	const changeInputStyle = (ev: React.ChangeEvent<{ value: unknown }>) => {
		if (ev.target.value === 'standard' || ev.target.value === 'outlined' || ev.target.value === 'filled') {
			store.dispatch(SetInputStyle(ev.target.value));
		}
	};

	useEffect(() => {
		store.dispatch(SetPrimary(primary));
	}, [primary]);
	useEffect(() => {
		store.dispatch(SetSecondary(secondary));
	}, [secondary]);

	return (
		<Dialog open={open} onClose={() => store.dispatch(ToggleThemeDialog)}>
			<DialogTitle>Customize Theme</DialogTitle>
			<DialogContent>
				<FormControl component="fieldset">
					<FormLabel component="legend">Theme Type</FormLabel>
					<RadioGroup value={themeType} row onChange={() => store.dispatch(InvertTheme)}>
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
				<Button color="primary" onClick={() => store.dispatch(ToggleThemeDialog)}>
					Close
				</Button>
			</DialogActions>
		</Dialog>
	);
};
