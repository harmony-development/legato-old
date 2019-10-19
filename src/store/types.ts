import { Color } from '@material-ui/core';

export enum Actions {
  INVERT_THEME,
  TOGGLE_NAME_DIALOG,
  UPDATE_USER
}

// this is the interface which lays out the types for each state item.
export interface IAppState {
  theme: {
    type: 'dark' | 'light';
    primary: Color;
    secondary: Color;
  };
  user: IUser;
  nameDialog: boolean;
}

export interface IUser {
  name: string;
  icon: string;
}

export interface IInvertTheme {
  type: Actions.INVERT_THEME;
}

export interface IShowChangeNameDialog {
  type: Actions.TOGGLE_NAME_DIALOG;
}

export interface IUpdateUser {
  type: Actions.UPDATE_USER;
  payload: IUser;
}

export type ActionTypes = IInvertTheme | IShowChangeNameDialog | IUpdateUser;
