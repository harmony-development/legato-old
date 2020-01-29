import { createReducer, createAction } from '@reduxjs/toolkit';
import { Color } from '@material-ui/core';
import { red, orange } from '@material-ui/core/colors';

import { IMessage, IChannels, IState, IGuild } from '../types/redux';

const appState: IState = {
	theme: {
		type: 'dark',
		primary: red,
		secondary: orange,
		inputStyle: 'filled',
	},
	guildList: {},
	guildMembers: {},
	themeDialog: false,
	connected: false,
	messages: [],
	invites: {},
	channels: {},
	currentGuild: undefined,
	currentChannel: undefined,
	guildDialog: false,
	guildSettingsDialog: false,
	users: {},
	self: {},
	userSettingsDialog: false,
	chatInputFocus: false,
};

function WithPayload<T>() {
	return (t: T) => ({ payload: t });
}

export const SetConnected = createAction('SET_CONNECTED', WithPayload<boolean>());
export const SetMessages = createAction('SET_MESSAGES', WithPayload<IMessage[]>());
export const AddMessage = createAction('ADD_MESSAGE', WithPayload<IMessage>());
export const SetCurrentChannel = createAction('SET_CURRENT_CHANNEl', WithPayload<string | undefined>());
export const SetCurrentGuild = createAction('SET_CURRENT_GUILD', WithPayload<string | undefined>());
export const SetChannels = createAction('SET_CHANNELS', WithPayload<IChannels>());
export const SetGuilds = createAction(
	'SET_GUILDS',
	WithPayload<{
		[key: string]: IGuild;
	}>()
);
export const RemoveGuild = createAction(
	'REMOVE_GUILD',
	WithPayload<{
		guild: string;
	}>()
);
export const SetGuildMembers = createAction(
	'SET_GUILD_MEMBERS',
	WithPayload<{
		guild: string;
		members: string[];
	}>()
);
export const FocusChatInput = createAction('FOCUS_CHAT_INPUT');
export const ToggleThemeDialog = createAction('TOGGLE_THEME_DIALOG');
export const InvertTheme = createAction('INVERT_THEME');
export const SetPrimary = createAction('SET_PRIMARY', WithPayload<Color>());
export const SetSecondary = createAction('SET_SECONDARY', WithPayload<Color>());
export const SetInputStyle = createAction('SET_INPUT_STYLE', WithPayload<'standard' | 'filled' | 'outlined'>());
export const ToggleGuildDialog = createAction('TOGGLE_GUILD_DIALOG');
export const ToggleGuildSettingsDialog = createAction('TOGGLE_GUILD_SETTINGS_DIALOG');
export const ToggleUserSettingsDialog = createAction('TOGGLE_USER_SETTINGS_DIALOG');
export const SetGuildPicture = createAction(
	'SET_GUILD_PICTURE',
	WithPayload<{
		guild: string;
		picture: string;
	}>()
);
export const SetGuildName = createAction(
	'SET_GUILD_NAME',
	WithPayload<{
		guild: string;
		name: string;
	}>()
);
export const SetInvites = createAction(
	'SET_INVITES',
	WithPayload<{
		[key: string]: number;
	}>()
);
export const SetUser = createAction(
	'SET_USER',
	WithPayload<{
		userid: string;
		username: string;
		avatar: string;
	}>()
);
export const SetAvatar = createAction(
	'SET_AVATAR',
	WithPayload<{
		userid: string;
		avatar: string;
	}>()
);
export const SetUsername = createAction(
	'SET_USERNAME',
	WithPayload<{
		userid: string;
		username: string;
	}>()
);
export const SetSelf = createAction(
	'SET_SELF',
	WithPayload<{
		userid: string;
		username: string;
		avatar: string;
	}>()
);
export const AppReducer = createReducer(appState, builder =>
	builder
		.addCase(SetConnected, (state, action) => ({
			...state,
			connected: action.payload,
		}))
		.addCase(SetMessages, (state, action) => ({
			...state,
			messages: action.payload,
		}))
		.addCase(AddMessage, (state, action) => ({
			...state,
			messages: [...state.messages, action.payload],
		}))
		.addCase(SetCurrentChannel, (state, action) => ({
			...state,
			currentChannel: action.payload,
		}))
		.addCase(SetCurrentGuild, (state, action) => ({
			...state,
			currentGuild: action.payload,
		}))
		.addCase(SetChannels, (state, action) => ({
			...state,
			channels: action.payload,
		}))
		.addCase(SetGuilds, (state, action) => ({
			...state,
			guildList: action.payload,
		}))
		.addCase(RemoveGuild, (state, action) => {
			const newGuildList = state.guildList;
			const deleteID = action.payload.guild;
			delete newGuildList[deleteID];
		})
		.addCase(SetGuildMembers, (state, action) => ({
			...state,
			guildMembers: {
				...state.guildMembers,
				[action.payload.guild]: action.payload.members,
			},
		}))
		.addCase(FocusChatInput, state => ({
			...state,
			chatInputFocus: true,
		}))
		.addCase(ToggleThemeDialog, state => ({
			...state,
			themeDialog: !state.themeDialog,
		}))
		.addCase(InvertTheme, state => ({
			...state,
			theme: {
				...state.theme,
				type: state.theme.type === 'dark' ? 'light' : 'dark',
			},
		}))
		.addCase(SetPrimary, (state, action) => ({
			...state,
			theme: {
				...state.theme,
				primary: action.payload,
			},
		}))
		.addCase(SetSecondary, (state, action) => ({
			...state,
			theme: {
				...state.theme,
				secondary: action.payload,
			},
		}))
		.addCase(SetInputStyle, (state, action) => ({
			...state,
			theme: {
				...state.theme,
				inputStyle: action.payload,
			},
		}))
		.addCase(ToggleGuildDialog, state => ({
			...state,
			guildDialog: !state.guildDialog,
		}))
		.addCase(ToggleGuildSettingsDialog, state => ({
			...state,
			guildSettingsDialog: !state.guildSettingsDialog,
		}))
		.addCase(ToggleUserSettingsDialog, state => ({
			...state,
			userSettingsDialog: !state.userSettingsDialog,
		}))
		.addCase(SetGuildPicture, (state, action) => ({
			...state,
			guildList: {
				...state.guildList,
				[action.payload.guild]: {
					...state.guildList[action.payload.guild],
					picture: action.payload.picture,
				},
			},
		}))
		.addCase(SetGuildName, (state, action) => ({
			...state,
			guildList: {
				...state.guildList,
				[action.payload.guild]: {
					...state.guildList[action.payload.guild],
					guildname: action.payload.name,
				},
			},
		}))
		.addCase(SetInvites, (state, action) => ({
			...state,
			invites: action.payload,
		}))
		.addCase(SetUser, (state, action) => ({
			...state,
			users: {
				...state.users,
				[action.payload.userid]: {
					avatar: action.payload.avatar,
					username: action.payload.username,
				},
			},
		}))
		.addCase(SetAvatar, (state, action) => ({
			...state,
			users: {
				...state.users,
				[action.payload.userid]: {
					...state.users[action.payload.userid],
					avatar: action.payload.avatar,
				},
			},
		}))
		.addCase(SetUsername, (state, action) => ({
			...state,
			users: {
				...state.users,
				[action.payload.userid]: {
					...state.users[action.payload.userid],
					username: action.payload.username,
				},
			},
		}))
		.addCase(SetSelf, (state, action) => ({
			...state,
			self: action.payload,
		}))
);
