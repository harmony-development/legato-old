import { ITheme } from './theming';
import { Color } from '@material-ui/core';

export enum Actions {
    INVERT_THEME,
    TOGGLE_THEME_DIALOG,
    CHANGE_PRIMARY,
    CHANGE_SECONDARY,
    SET_CONNECTED
}

export interface IState {
    theme: ITheme;
    themeDialog: boolean;
    connected: boolean;
}

export interface IInvertTheme {
    type: Actions.INVERT_THEME;
}

export interface IToggleThemeDialog {
    type: Actions.TOGGLE_THEME_DIALOG;
}

export interface IChangePrimary {
    type: Actions.CHANGE_PRIMARY;
    payload: Color;
}

export interface IChangeSecondary {
    type: Actions.CHANGE_SECONDARY;
    payload: Color;
}

export interface ISetConnected {
    type: Actions.SET_CONNECTED;
    payload: boolean;
}

export type Action = IInvertTheme | IToggleThemeDialog | IChangePrimary | IChangeSecondary | ISetConnected;
