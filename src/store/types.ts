import { Color } from '@material-ui/core';

export enum Actions {
  INVERT_THEME,
  TOGGLE_NAME_DIALOG,
  CHANGE_NAME
}

// this is the interface which lays out the types for each state item.
export interface IAppState {
  theme: {
    type: 'dark' | 'light';
    primary: Color;
    secondary: Color;
  };
  nameDialog: boolean;
  name: string;
}

export interface IInvertTheme {
  type: Actions.INVERT_THEME;
}

export interface IShowChangeNameDialog {
  type: Actions.TOGGLE_NAME_DIALOG;
}

export interface IChangeName {
  type: Actions.CHANGE_NAME;
  payload: string;
}

export const Events = {
  USERNAME_UPDATE: 'USERNAME_UPDATE',
  MESSAGE: 'MESSAGE',
  LOGIN: 'LOGIN',
  DISCONNECT: 'DISCONNECT'
};

export type ActionTypes = IInvertTheme | IShowChangeNameDialog | IChangeName;
