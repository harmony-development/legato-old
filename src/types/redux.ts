import { ITheme } from './theming';
import { Color } from '@material-ui/core';

export enum Actions {
    TOGGLE_THEME_DIALOG,
    TOGGLE_JOIN_GUILD_DIALOG,
    TOGGLE_GUILD_SETTINGS_DIALOG,
    TOGGLE_USER_SETTINGS_DIALOG,

    INVERT_THEME,
    CHANGE_PRIMARY,
    CHANGE_SECONDARY,
    SET_INPUT_STYLE,

    SET_CONNECTED,

    SET_GUILDS,
    SET_SELECTED_GUILD,
    SET_CHANNELS,
    SET_SELECTED_CHANNEL,
    SET_GUILD_PICTURE,
    SET_GUILD_NAME,
    SET_INVITES,
    SET_USER,
    SET_MESSAGES,
    ADD_MESSAGE,

    FOCUS_CHAT_INPUT
}

interface IGuild {
    guildid: string;
    picture: string;
    guildname: string;
    owner: boolean;
}

interface ISetGuildPicturePayload {
    guild: string;
    picture: string;
}

interface ISetGuildNamePayload {
    guild: string;
    name: string;
}

export interface IChannels {
    [key: string]: string;
}

export interface IMessage {
    userid: string;
    createdat: number;
    guild: string;
    channel: string;
    message: string;
    messageid: string;
}

export interface IInvertTheme {
    type: Actions.INVERT_THEME;
}

export interface IToggleThemeDialog {
    type: Actions.TOGGLE_THEME_DIALOG;
}

export interface IChangePrimary {
    type: Actions.CHANGE_PRIMARY;
    payload: Color;
}

export interface IChangeSecondary {
    type: Actions.CHANGE_SECONDARY;
    payload: Color;
}

export interface ISetConnected {
    type: Actions.SET_CONNECTED;
    payload: boolean;
}

export interface ISetGuilds {
    type: Actions.SET_GUILDS;
    payload: {
        [key: string]: IGuild;
    };
}

export interface ISetSelectedGuild {
    type: Actions.SET_SELECTED_GUILD;
    payload: string;
}

export interface IAddMessage {
    type: Actions.ADD_MESSAGE;
    payload: IMessage;
}

export interface ISetMessages {
    type: Actions.SET_MESSAGES;
    payload: IMessage[];
}

export interface ISetInputStyle {
    type: Actions.SET_INPUT_STYLE;
    payload: 'standard' | 'filled' | 'outlined';
}

export interface ISetChannels {
    type: Actions.SET_CHANNELS;
    payload: {
        [key: string]: string;
    };
}

export interface ISetSelectedChannel {
    type: Actions.SET_SELECTED_CHANNEL;
    payload: string;
}

export interface IToggleJoinGuildDialog {
    type: Actions.TOGGLE_JOIN_GUILD_DIALOG;
}

export interface IToggleGuildSettingsDialog {
    type: Actions.TOGGLE_GUILD_SETTINGS_DIALOG;
}

export interface ISetGuildPicture {
    type: Actions.SET_GUILD_PICTURE;
    payload: ISetGuildPicturePayload;
}

export interface ISetGuildName {
    type: Actions.SET_GUILD_NAME;
    payload: ISetGuildNamePayload;
}

export interface ISetInvites {
    type: Actions.SET_INVITES;
    payload: {
        [key: string]: number;
    };
}

export interface ISetUser {
    type: Actions.SET_USER;
    payload: {
        userid: string;
        username: string;
        avatar: string;
    };
}

export interface IToggleUserSettingsDialog {
    type: Actions.TOGGLE_USER_SETTINGS_DIALOG;
}

export interface IFocusChatInput {
    type: Actions.FOCUS_CHAT_INPUT;
}

export interface IState {
    theme: ITheme;
    themeDialog: boolean;
    connected: boolean;
    guildList: {
        [key: string]: IGuild;
    };
    selectedGuild: string;
    messages: IMessage[];
    inputStyle: 'standard' | 'filled' | 'outlined';
    channels: {
        [key: string]: string;
    };
    selectedChannel: string | undefined;
    joinGuildDialog: boolean;
    guildSettingsDialog: boolean;
    invites: {
        [key: string]: number;
    };
    users: {
        [key: string]: {
            username: string;
            avatar: string;
        };
    };
    userSettingsDialog: boolean;
    chatInputFocus: boolean;
}

export type Action =
    | IInvertTheme
    | IToggleThemeDialog
    | IChangePrimary
    | IChangeSecondary
    | ISetConnected
    | ISetGuilds
    | ISetSelectedGuild
    | IAddMessage
    | ISetMessages
    | ISetInputStyle
    | ISetChannels
    | ISetSelectedChannel
    | IToggleJoinGuildDialog
    | IToggleGuildSettingsDialog
    | ISetGuildPicture
    | ISetGuildName
    | ISetInvites
    | ISetUser
    | IToggleUserSettingsDialog
    | IFocusChatInput;
