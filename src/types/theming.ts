import { Color } from '@material-ui/core';

export interface ITheme {
	type: 'light' | 'dark';
	primary: Color;
	secondary: Color;
	inputStyle: 'standard' | 'filled' | 'outlined';
}
