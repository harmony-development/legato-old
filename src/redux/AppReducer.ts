import { IState, Action, Actions } from '../types/redux';
import { red, orange } from '@material-ui/core/colors';

const appState: IState = {
    theme: {
        type: 'dark',
        primary: red,
        secondary: orange
    },
    guildList: {},
    themeDialog: false,
    connected: false,
    selectedGuild: '',
    messages: [],
    inputStyle: 'standard'
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
        case Actions.SET_GUILDS: {
            return {
                ...state,
                guildList: action.payload
            };
        }
        case Actions.SET_SELECTED_GUILD: {
            return {
                ...state,
                selectedGuild: action.payload
            };
        }
        case Actions.ADD_MESSAGE: {
            return {
                ...state,
                messages: [...state.messages, action.payload]
            };
        }
        case Actions.SET_MESSAGES: {
            return {
                ...state,
                messages: action.payload
            };
        }
        case Actions.SET_INPUT_STYLE: {
            return {
                ...state,
                inputStyle: action.payload
            };
        }
        default: {
            return state;
        }
    }
}
