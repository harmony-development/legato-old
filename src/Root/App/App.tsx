import React, { useEffect } from 'react';
import { HarmonyBar } from './HarmonyBar/HarmonyBar';
import { ThemeDialog } from './Dialog/ThemeDialog';
import { useAppStyles } from './AppStyle';
import { ChatArea } from './ChatArea/ChatArea';
import { harmonySocket } from '../Root';
import { useHistory } from 'react-router';
import { IGuildData } from '../../types/socket';
import { useDispatch, useSelector } from 'react-redux';
import { Actions, IState, IMessage } from '../../types/redux';
import { toast } from 'react-toastify';
import { JoinGuild } from './Dialog/JoinGuildDialog/JoinGuild';
import { GuildSettings } from './Dialog/GuildSettingsDialog/GuildSettings';
import {
    SetMessages,
    SetSelectedChannel,
    SetSelectedGuild,
    SetChannels,
    SetGuilds,
    AddMessage,
    ToggleGuildSettingsDialog,
    SetGuildPicture,
    SetInvites,
    SetGuildName,
    SetUsername
} from '../../redux/Dispatches';

export const App = () => {
    const classes = useAppStyles();
    const dispatch = useDispatch();
    const connected = useSelector((state: IState) => state.connected);
    const channels = useSelector((state: IState) => state.channels);
    const invites = useSelector((state: IState) => state.invites);
    const selectedGuild = useSelector((state: IState) => state.selectedGuild);
    const themeDialogOpen = useSelector((state: IState) => state.themeDialog);
    const joinDialogOpen = useSelector((state: IState) => state.joinGuildDialog);
    const guildSettingsDialogOpen = useSelector((state: IState) => state.guildSettingsDialog);
    const history = useHistory();
    let eventsBound = false;

    // event when the client has connected
    useEffect(() => {
        if (connected) {
            harmonySocket.getGuilds();
        }
    }, [connected]);

    useEffect(() => {
        if (connected) {
            dispatch(SetMessages([]));
            dispatch(SetSelectedChannel(undefined));
            harmonySocket.getMessages(selectedGuild);
            harmonySocket.getChannels(selectedGuild);
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [selectedGuild]);

    useEffect(() => {
        if (!eventsBound) {
            if ((harmonySocket.conn.readyState !== WebSocket.OPEN && harmonySocket.conn.readyState !== WebSocket.CONNECTING) || typeof localStorage.getItem('token') !== 'string') {
                // bounce the user to the login screen if the socket is disconnected or there's no token
                history.push('/');
                return;
            }

            harmonySocket.events.addListener('getguilds', (raw: any) => {
                let guildsList = raw['guilds'] as IGuildData[];
                if (Object.keys(guildsList).length === 0) {
                    dispatch(SetMessages([]));
                    dispatch(SetSelectedChannel(undefined));
                    dispatch(SetSelectedGuild(undefined));
                    dispatch(SetChannels({}));
                }
                dispatch(SetGuilds(guildsList));
            });
            harmonySocket.events.addListener('getmessages', (raw: any) => {
                if (raw['messages']) {
                    dispatch(SetMessages((raw['messages'] as IMessage[]).reverse()));
                }
            });
            harmonySocket.events.addListener('message', (raw: any) => {
                // prevent stupid API responses
                if (typeof raw['userid'] === 'string' && typeof raw['createdat'] === 'number' && typeof raw['guild'] === 'string' && typeof raw['message'] === 'string') {
                    dispatch(AddMessage(raw as IMessage));
                }
            });
            harmonySocket.events.addListener('getchannels', (raw: any) => {
                if (typeof raw === 'object') {
                    dispatch(SetChannels(raw['channels']));
                }
            });
            harmonySocket.events.addListener('deauth', () => {
                toast.warn('Your session has expired. Please login again');
                history.push('/');
                return;
            });
            harmonySocket.events.addListener('leaveguild', (raw: any) => {
                if (typeof raw['message'] === 'string') {
                    toast.error(raw['message']);
                    return;
                }
                harmonySocket.getGuilds();
            });
            harmonySocket.events.addListener('joinguild', (raw: any) => {
                if (!raw['message']) {
                    harmonySocket.getGuilds();
                    dispatch({ type: Actions.TOGGLE_JOIN_GUILD_DIALOG });
                }
            });
            harmonySocket.events.addListener('createguild', (raw: any) => {
                if (!raw['message']) {
                    harmonySocket.getGuilds();
                    dispatch({ type: Actions.TOGGLE_JOIN_GUILD_DIALOG });
                }
            });
            harmonySocket.events.addListener('updateguildpicture', (raw: any) => {
                if (raw['success'] === true && raw['picture'] && raw['guild']) {
                    dispatch(SetGuildPicture(raw['guild'], raw['picture']));
                    if (guildSettingsDialogOpen) {
                        dispatch(ToggleGuildSettingsDialog());
                    }
                } else {
                    toast.error('Error saving guild');
                }
            });
            harmonySocket.events.addListener('updateguildname', (raw: any) => {
                if (raw['success'] === true && raw['name'] && raw['guild']) {
                    dispatch(SetGuildName(raw['guild'], raw['name']));
                    if (guildSettingsDialogOpen) {
                        dispatch(ToggleGuildSettingsDialog());
                    }
                } else {
                    toast.error('Error saving guild');
                }
            });
            harmonySocket.events.addListener('getinvites', (raw: any) => {
                if (raw['invites'] && raw['guild']) {
                    dispatch(SetInvites(raw['invites']));
                }
            });
            harmonySocket.events.addListener('addguildchannel', (raw: any) => {
                if (raw['success'] === true && raw['guild'] && raw['channelname'] && raw['channelid']) {
                    dispatch(SetChannels({ ...channels, [raw['channelid']]: raw['channelname'] }));
                }
            });
            harmonySocket.events.addListener('deleteguildchannel', (raw: any) => {
                if (raw['success'] === true && raw['guild'] && raw['channelid']) {
                    const channelDeleted = {
                        ...channels
                    };
                    delete channelDeleted[raw['channelid']];
                    dispatch(SetChannels({ ...channelDeleted }));
                }
            });
            harmonySocket.events.addListener('deleteinvite', (raw: any) => {
                if (raw['success'] === true && raw['invite']) {
                    const invitesDeleted = {
                        ...invites
                    };
                    delete invitesDeleted[raw['invite']];
                    dispatch(SetInvites(invitesDeleted));
                }
            });
            harmonySocket.events.addListener('createinvite', (raw: any) => {
                if (raw['success'] === true && raw['invite']) {
                    const invitesDeleted = {
                        ...invites,
                        [raw['invite']]: 0
                    };
                    dispatch(SetInvites(invitesDeleted));
                }
            });
            harmonySocket.events.addListener('getusername', (raw: any) => {
                if (raw['userid'] && raw['username']) {
                    dispatch(SetUsername(raw['userid'], raw['username']));
                }
            });
            return () => {
                harmonySocket.events.removeAllListeners('getguilds');
                harmonySocket.events.removeAllListeners('getmessages');
                harmonySocket.events.removeAllListeners('message');
                harmonySocket.events.removeAllListeners('getchannels');
                harmonySocket.events.removeAllListeners('deauth');
                harmonySocket.events.removeAllListeners('leaveguild');
                harmonySocket.events.removeAllListeners('joinguild');
                harmonySocket.events.removeAllListeners('createguild');
                harmonySocket.events.removeAllListeners('updateguildpicture');
                harmonySocket.events.removeAllListeners('updateguildname');
                harmonySocket.events.removeAllListeners('getinvites');
                harmonySocket.events.removeAllListeners('addguildchannel');
                harmonySocket.events.removeAllListeners('deleteguildchannel');
                harmonySocket.events.removeAllListeners('deleteinvite');
                harmonySocket.events.removeAllListeners('createinvite');
            };
        }
    }, [history, dispatch, guildSettingsDialogOpen, eventsBound, channels, invites]);

    return (
        <div className={classes.root}>
            {themeDialogOpen ? <ThemeDialog /> : undefined}
            {joinDialogOpen ? <JoinGuild /> : undefined}
            {guildSettingsDialogOpen ? <GuildSettings /> : undefined}
            <HarmonyBar />
            <div className={classes.navFill} /> {/* this fills the area where the navbar is*/}
            <ChatArea />
        </div>
    );
};
