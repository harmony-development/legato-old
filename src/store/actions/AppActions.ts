import { Actions, IInvertTheme, IChangeName, IShowChangeNameDialog as IToggleChangeNameDialog } from '../types';

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

export function changeName(payload: string): IChangeName {
  return {
    type: Actions.CHANGE_NAME,
    payload
  };
}
