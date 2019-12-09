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
    invites: {},
    inputStyle: 'filled',
    channels: {},
    selectedChannel: undefined,
    joinGuildDialog: false,
    guildSettingsDialog: false,
    usernames: {},
    userSettingsDialog: false
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
        case Actions.SET_CHANNELS: {
            return {
                ...state,
                channels: action.payload
            };
        }
        case Actions.SET_SELECTED_CHANNEL: {
            return {
                ...state,
                selectedChannel: action.payload
            };
        }
        case Actions.TOGGLE_JOIN_GUILD_DIALOG: {
            return {
                ...state,
                joinGuildDialog: !state.joinGuildDialog
            };
        }
        case Actions.TOGGLE_GUILD_SETTINGS_DIALOG: {
            return {
                ...state,
                guildSettingsDialog: !state.guildSettingsDialog
            };
        }
        case Actions.SET_GUILD_PICTURE: {
            return {
                ...state,
                guildList: {
                    ...state.guildList,
                    [action.payload.guild]: {
                        ...state.guildList[action.payload.guild],
                        picture: action.payload.picture
                    }
                }
            };
        }
        case Actions.SET_GUILD_NAME: {
            return {
                ...state,
                guildList: {
                    ...state.guildList,
                    [action.payload.guild]: {
                        ...state.guildList[action.payload.guild],
                        guildname: action.payload.name
                    }
                }
            };
        }
        case Actions.SET_INVITES: {
            return {
                ...state,
                invites: action.payload
            };
        }
        case Actions.SET_USERNAME: {
            return {
                ...state,
                usernames: {
                    ...state.usernames,
                    [action.payload.userid]: action.payload.username
                }
            };
        }
        case Actions.TOGGLE_USER_SETTINGS_DIALOG: {
            return {
                ...state,
                userSettingsDialog: !state.userSettingsDialog
            };
        }
        default: {
            return state;
        }
    }
}
