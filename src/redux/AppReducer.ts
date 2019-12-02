import { IState, Action, Actions } from '../types/redux';
import { red, orange } from '@material-ui/core/colors';

const appState: IState = {
    theme: {
        type: 'dark',
        primary: red,
        secondary: orange
    },
    themeDialog: false,
    connected: false
};

export default function AppReducer(state = appState, action: Action): IState {
    switch (action.type) {
        case Actions.INVERT_THEME: {
            return {
                ...state,
                theme: {
                    ...state.theme,
                    type: state.theme.type === 'dark' ? 'light' : 'dark'
                }
            };
        }
        case Actions.TOGGLE_THEME_DIALOG: {
            return {
                ...state,
                themeDialog: !state.themeDialog
            };
        }
        case Actions.CHANGE_PRIMARY: {
            return {
                ...state,
                theme: {
                    ...state.theme,
                    primary: action.payload
                }
            };
        }
        case Actions.CHANGE_SECONDARY: {
            return {
                ...state,
                theme: {
                    ...state.theme,
                    secondary: action.payload
                }
            };
        }
        case Actions.SET_CONNECTED: {
            return {
                ...state,
                connected: action.payload
            };
        }
        default: {
            return state;
        }
    }
}
