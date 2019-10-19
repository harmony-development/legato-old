import {
  Actions,
  IInvertTheme,
  IShowChangeNameDialog as IToggleChangeNameDialog,
  IUpdateUser, IUser
} from '../types';

export function invertTheme(): IInvertTheme {
  return {
    type: Actions.INVERT_THEME
  };
}

export function toggleChangeNameDialog(): IToggleChangeNameDialog {
  return {
    type: Actions.TOGGLE_NAME_DIALOG
  };
}

export function updateUser(payload: IUser): IUpdateUser {
  return  {
    type: Actions.UPDATE_USER,
    payload
  }
}