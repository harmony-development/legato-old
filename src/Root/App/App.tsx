import React, { useEffect } from 'react';
import { useHistory, useParams } from 'react-router';
import { useDispatch, useSelector } from 'react-redux';

import { harmonySocket } from '../Root';
import { IState } from '../../types/redux';
import { SetMessages, SetSelectedChannel, SetCurrentGuild } from '../../redux/Dispatches';
import { useSocketHandler } from '../SocketHandler';

import { HarmonyBar } from './HarmonyBar/HarmonyBar';
import { ThemeDialog } from './Dialog/ThemeDialog';
import { useAppStyles } from './AppStyle';
import { ChatArea } from './ChatArea/ChatArea';
import { JoinGuild } from './Dialog/JoinGuildDialog/JoinGuild';
import { GuildSettings } from './Dialog/GuildSettingsDialog/GuildSettings';
import { UserSettingsDialog } from './Dialog/UserSettingsDialog/UserSettingsDialog';

export const App = (): JSX.Element => {
	const classes = useAppStyles();
	const dispatch = useDispatch();
	const { selectedguildparam: selectedGuildParam, selectedchannelparam: selectedChannelParam } = useParams();
	const [
		channels,
		currentGuild,
		selectedChannel,
		themeDialogOpen,
		joinDialogOpen,
		guildSettingsDialogOpen,
		userSettingsDialogOpen,
	] = useSelector((state: IState) => [
		state.channels,
		state.currentGuild,
		state.selectedChannel,
		state.themeDialog,
		state.joinGuildDialog,
		state.guildSettingsDialog,
		state.userSettingsDialog,
	]);
	const history = useHistory();
	useSocketHandler(harmonySocket, history);

	useEffect(() => {
		if (selectedChannelParam) {
			dispatch(SetSelectedChannel(selectedChannelParam));
		}
	}, [selectedChannelParam]);

	useEffect(() => {
		if (selectedGuildParam) {
			dispatch(SetCurrentGuild(selectedGuildParam));
		}
	}, [selectedGuildParam]);

	useEffect(() => {
		if (currentGuild) {
			history.push(`/app/${currentGuild}/${selectedChannel || ''}`);
			dispatch(SetMessages([]));
			dispatch(SetSelectedChannel(undefined));
			harmonySocket.exec(() => harmonySocket.getChannels(currentGuild));
		}
	}, [currentGuild]);

	useEffect(() => {
		if (currentGuild && selectedChannel) {
			document.title = `Harmony - ${channels[selectedChannel] || 'FOSS Chat Client'}`;
			history.push(`/app/${currentGuild}/${selectedChannel}`);
		}
	}, [selectedChannel, history, channels, currentGuild]);

	return (
		<div className={classes.root}>
			{themeDialogOpen ? <ThemeDialog /> : undefined}
			{joinDialogOpen ? <JoinGuild /> : undefined}
			{guildSettingsDialogOpen ? <GuildSettings /> : undefined}
			{userSettingsDialogOpen ? <UserSettingsDialog /> : undefined}
			<HarmonyBar />
			<div className={classes.navFill} /> {/* this fills the area where the navbar is*/}
			<ChatArea />
		</div>
	);
};
