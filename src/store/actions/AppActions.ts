import { Actions, IInvertTheme, IToggleProfileSettingsDialog, IUpdateUser, IUser } from '../types';

export function invertTheme(): IInvertTheme {
    return {
        type: Actions.INVERT_THEME
    };
}

export function toggleProfileSettingsDialog(): IToggleProfileSettingsDialog {
    return {
        type: Actions.TOGGLE_PROFILE_SETTINGS_DIALOG
    };
}

export function updateUser(payload: IUser): IUpdateUser {
    return {
        type: Actions.UPDATE_USER,
        payload
    };
}
