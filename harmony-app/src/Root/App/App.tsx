import React, { useEffect } from 'react';
import { useHistory, useParams } from 'react-router';
import { useSelector, useDispatch } from 'react-redux';

import { IState } from '../../types/redux';
import { SetCurrentGuild, SetCurrentChannel } from '../../redux/AppReducer';

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
	const {
		channels,
		currentGuild,
		currentChannel,
		themeDialog,
		guildDialog,
		guildSettingsDialog,
		userSettingsDialog,
	} = useSelector((state: IState) => state);
	const history = useHistory();
	const { selectedguildparam: selectedGuildParam, selectedchannelparam: selectedChannelParam } = useParams();

	useEffect(() => {
		if (selectedGuildParam) {
			dispatch(SetCurrentChannel(selectedChannelParam));
		}
	}, [selectedChannelParam]);

	useEffect(() => {
		if (selectedGuildParam) {
			dispatch(SetCurrentGuild(selectedGuildParam));
		}
	}, [selectedGuildParam]);

	useEffect(() => {
		if (currentGuild && currentChannel) {
			document.title = `Harmony - ${channels[currentChannel] || 'FOSS Chat Client'}`;
			history.push(`/app/${currentGuild}/${currentChannel}`);
		}
	}, [currentChannel, history, channels, currentGuild]);

	return (
		<div className={classes.root}>
			{themeDialog ? <ThemeDialog /> : undefined}
			{guildDialog ? <JoinGuild /> : undefined}
			{guildSettingsDialog ? <GuildSettings /> : undefined}
			{userSettingsDialog ? <UserSettingsDialog /> : undefined}
			<HarmonyBar />
			<div className={classes.navFill} /> {/* this fills the area where the navbar is*/}
			<ChatArea />
		</div>
	);
};
