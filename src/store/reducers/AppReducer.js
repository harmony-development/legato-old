/**
 * This file is the controller for Redux.
 *
 * It handles dispatches from ../actions/AppActions.ts and returns a new state.
 *
 * The initial state is never mutated. It always returns a copy of the original state.
 */
import { red, purple } from '@material-ui/core/colors';
import { socketServer } from '../../Root/Root';

const initialState = {
    theme: {
        type: 'dark',
        primary: red,
        secondary: purple
    },
    nameDialog: false,
    user: {
        username: '',
        avatar: ''
    }
};

export default function AppReducer(state = initialState, action) {
    switch (action.type) {
        case 'INVERT_THEME': {
            socketServer.saveProfile({ theme: state.theme, token: localStorage.getItem('token') });
            return {
                ...state,
                theme: {
                    ...state.theme,
                    type: state.theme.type === 'dark' ? 'light' : 'dark'
                }
            };
        }
        case 'TOGGLE_PROFILE_SETTINGS_DIALOG': {
            return {
                ...state,
                nameDialog: !state.nameDialog
            };
        }
        case 'UPDATE_USER': {
            return {
                ...state,
                user: action.payload
            };
        }
        default:
            return state;
    }
}
