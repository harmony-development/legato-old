import { Color } from '@material-ui/core';

export enum Actions {
  INVERT_THEME,
  TOGGLE_PROFILE_SETTINGS_DIALOG,
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
  username: string;
  avatar: string;
}

export interface IInvertTheme {
  type: Actions.INVERT_THEME;
}

export interface IToggleProfileSettingsDialog {
  type: Actions.TOGGLE_PROFILE_SETTINGS_DIALOG;
}

export interface IUpdateUser {
  type: Actions.UPDATE_USER;
  payload: IUser;
}

export type ActionTypes = IInvertTheme | IToggleProfileSettingsDialog | IUpdateUser;
