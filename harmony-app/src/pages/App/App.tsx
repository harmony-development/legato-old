import React, { useEffect } from 'react';
import { useHistory, useParams } from 'react-router';
import { useSelector, useDispatch } from 'react-redux';
import { makeStyles, Theme } from '@material-ui/core';

import { SetCurrentGuild, SetCurrentChannel } from '../../redux/AppReducer';
import { RootState } from '../../redux/store';

import { HarmonyBar } from './HarmonyBar/HarmonyBar';
import { ThemeDialog } from './Dialog/ThemeDialog';
import { ChatArea } from './ChatArea/ChatArea';
import { JoinGuild } from './Dialog/JoinGuildDialog/JoinGuild';
import { GuildSettings } from './Dialog/GuildSettingsDialog/GuildSettings';
import { UserSettingsDialog } from './Dialog/UserSettingsDialog/UserSettingsDialog';
import { InstanceList } from './InstanceList/InstanceList';

const appStyles = makeStyles((theme: Theme) => ({
	root: {
		display: 'flex',
		height: '100%',
		flexDirection: 'column',
	},
	leftMenuBtn: {
		marginRight: theme.spacing(1),
	},
	title: {
		flexGrow: 1,
	},
	navFill: {
		...theme.mixins.toolbar,
	},
}));

export const App = React.memo(() => {
	const classes = appStyles();
	const dispatch = useDispatch();
	const { channels, currentGuild, currentChannel } = useSelector((state: RootState) => state.app);
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
			<ThemeDialog />
			<JoinGuild />
			<GuildSettings />
			<UserSettingsDialog />
			<HarmonyBar />
			<InstanceList />
			<div className={classes.navFill} /> {/* this fills the area where the navbar is*/}
			<ChatArea />
		</div>
	);
});
