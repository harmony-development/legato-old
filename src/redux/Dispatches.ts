import { Actions, IMessage, IChannels } from '../types/redux';
import { IGuildData } from '../types/socket';

/**
 * A function that sets the connected state to a value
 * @param payload The connection state to set it to
 */
export function SetConnected(payload: boolean) {
    return {
        type: Actions.SET_CONNECTED,
        payload
    };
}

/**
 * A function that sets the messages (displayed in the chat area)
 * @param payload The messages to assign
 */
export function SetMessages(payload: IMessage[]) {
    return {
        type: Actions.SET_MESSAGES,
        payload
    };
}

export function AddMessage(payload: IMessage) {
    return {
        type: Actions.ADD_MESSAGE,
        payload
    };
}

/**
 * A function that sets the selected channel. Used for the channel list
 * @param payload What to set the selected channel to
 */
export function SetSelectedChannel(payload: string | undefined) {
    return {
        type: Actions.SET_SELECTED_CHANNEL,
        payload
    };
}

/**
 * A function that sets the selected guild. Used for the guild list
 * @param payload The guild ID to set the selection to
 */
export function SetSelectedGuild(payload: string | undefined) {
    return {
        type: Actions.SET_SELECTED_GUILD,
        payload
    };
}

export function SetChannels(payload: IChannels) {
    return {
        type: Actions.SET_CHANNELS,
        payload
    };
}

export function SetGuilds(payload: IGuildData[]) {
    return {
        type: Actions.SET_GUILDS,
        payload
    };
}

export function ToggleThemeDialog() {
    return {
        type: Actions.TOGGLE_THEME_DIALOG
    };
}

export function ToggleGuildDialog() {
    return {
        type: Actions.TOGGLE_JOIN_GUILD_DIALOG
    };
}

export function ToggleGuildSettingsDialog() {
    return {
        type: Actions.TOGGLE_GUILD_SETTINGS_DIALOG
    };
}

export function SetGuildPicture(guild: string, picture: string) {
    return {
        type: Actions.SET_GUILD_PICTURE,
        payload: {
            guild,
            picture
        }
    };
}

interface IInvites {
    [key: string]: number;
}

export function SetInvites(invites: IInvites) {
    return {
        type: Actions.SET_INVITES,
        payload: invites
    };
}
