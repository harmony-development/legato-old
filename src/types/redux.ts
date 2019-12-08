import { ITheme } from './theming';
import { Color } from '@material-ui/core';

export enum Actions {
    INVERT_THEME,
    TOGGLE_THEME_DIALOG,
    CHANGE_PRIMARY,
    CHANGE_SECONDARY,
    SET_CONNECTED,
    SET_GUILDS,
    SET_SELECTED_GUILD,
    ADD_MESSAGE,
    SET_MESSAGES,
    SET_INPUT_STYLE,
    SET_CHANNELS,
    SET_SELECTED_CHANNEL,
    TOGGLE_JOIN_GUILD_DIALOG,
    TOGGLE_GUILD_SETTINGS_DIALOG,
    SET_GUILD_PICTURE,
    SET_GUILD_NAME,
    SET_INVITES,
    SET_USERNAME
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
    usernames: {
        [key: string]: string;
    };
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

export interface ISetUsername {
    type: Actions.SET_USERNAME;
    payload: {
        userid: string;
        username: string;
    };
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
    | ISetUsername;
