import React, { useEffect } from 'react';
import { HarmonyBar } from './HarmonyBar/HarmonyBar';
import { ThemeDialog } from './Dialog/ThemeDialog';
import { useAppStyles } from './AppStyle';
import { ChatArea } from './ChatArea/ChatArea';
import { harmonySocket } from '../Root';
import { useHistory, useParams } from 'react-router';
import { useDispatch, useSelector } from 'react-redux';
import { IState } from '../../types/redux';
import { JoinGuild } from './Dialog/JoinGuildDialog/JoinGuild';
import { GuildSettings } from './Dialog/GuildSettingsDialog/GuildSettings';
import { SetMessages, SetSelectedChannel, SetCurrentGuild } from '../../redux/Dispatches';
import { UserSettingsDialog } from './Dialog/UserSettingsDialog/UserSettingsDialog';
import { useSocketHandler } from '../SocketHandler';

export const App = () => {
    const classes = useAppStyles();
    const dispatch = useDispatch();
    const { selectedguildparam: selectedGuildParam, selectedchannelparam: selectedChannelParam } = useParams();
    const [
        connected,
        channels,
        currentGuild,
        selectedChannel,
        themeDialogOpen,
        joinDialogOpen,
        guildSettingsDialogOpen,
        userSettingsDialogOpen
    ] = useSelector((state: IState) => [
        state.connected,
        state.channels,
        state.currentGuild,
        state.selectedChannel,
        state.themeDialog,
        state.joinGuildDialog,
        state.guildSettingsDialog,
        state.userSettingsDialog
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
            if (!selectedChannelParam) {
                harmonySocket.events.addListener('open', () => {
                    harmonySocket.getChannels(selectedGuildParam);
                    harmonySocket.events.removeCurrentListener();
                });
            }
        }
    }, [selectedGuildParam]);

    useEffect(() => {
        if (currentGuild) {
            history.push(`/app/${currentGuild}/${selectedChannel || ''}`);
            if (connected) {
                dispatch(SetMessages([]));
                dispatch(SetSelectedChannel(undefined));
                harmonySocket.getChannels(currentGuild);
            }
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
