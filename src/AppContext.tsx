import { createContext } from 'react';
import { Color } from '@material-ui/core';
import { red, orange } from '@material-ui/core/colors';

interface IAppContext {
	themeType: 'dark' | 'light';
	themePrimary: Color;
	themeSecondary: Color;
	inputStyle: 'standard' | 'filled' | 'outlined';
	connected: boolean;

	themeDialogOpen: boolean;
}

export const AppContext = createContext<IAppContext>({
	themeType: 'dark',
	themePrimary: red,
	themeSecondary: orange,
	inputStyle: 'filled',
	connected: false,

	themeDialogOpen: false,
});
