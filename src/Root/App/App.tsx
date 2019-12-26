import React, { useEffect } from 'react';
import { HarmonyBar } from './HarmonyBar/HarmonyBar';
import { ThemeDialog } from './Dialog/ThemeDialog';
import { useAppStyles } from './AppStyle';
import { ChatArea } from './ChatArea/ChatArea';
import { harmonySocket } from '../Root';
import { useHistory, useParams } from 'react-router';
import { useDispatch, useSelector } from 'react-redux';
import { Actions, IState, IMessage } from '../../types/redux';
import { toast } from 'react-toastify';
import { JoinGuild } from './Dialog/JoinGuildDialog/JoinGuild';
import { GuildSettings } from './Dialog/GuildSettingsDialog/GuildSettings';
import { SetMessages, SetSelectedChannel, SetSelectedGuild } from '../../redux/Dispatches';
import { UserSettingsDialog } from './Dialog/UserSettingsDialog/UserSettingsDialog';
import { useSocketHandler } from '../SocketHandler';

export const App = () => {
    const classes = useAppStyles();
    const dispatch = useDispatch();
    const { selectedguildparam: selectedGuildParam, selectedchannelparam: selectedChannelParam } = useParams();
    const [
        connected,
        guilds,
        channels,
        invites,
        currentGuild,
        selectedChannel,
        themeDialogOpen,
        joinDialogOpen,
        guildSettingsDialogOpen,
        userSettingsDialogOpen
    ] = useSelector((state: IState) => [
        state.connected,
        state.guildList,
        state.channels,
        state.invites,
        state.selectedGuild,
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
            dispatch(SetSelectedGuild(selectedGuildParam));
        }
    }, [selectedGuildParam]);

    useEffect(() => {
        if (currentGuild) {
            history.push(`/app/${currentGuild}/${selectedChannel || ''}`);
            if (connected) {
                dispatch(SetMessages([]));
                dispatch(SetSelectedChannel(undefined));
                harmonySocket.getMessages(currentGuild);
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
