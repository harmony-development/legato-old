import { Color } from '@material-ui/core';

export enum Actions {
    INVERT_THEME
}

export interface IState {
    theme: {
        type: 'light' | 'dark';
        primary: Color;
        secondary: Color;
    };
}

export interface IInvertTheme {
    type: Actions.INVERT_THEME;
}

export type Action = IInvertTheme;
