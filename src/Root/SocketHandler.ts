import { useEffect, useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import h from 'history';
import { toast } from 'react-toastify';
import HarmonySocket from '../socket/socket';
import { IGuildsList } from '../types/socket';
import {
    SetMessages,
    SetSelectedChannel,
    SetCurrentGuild,
    SetChannels,
    SetGuilds,
    AddMessage,
    SetGuildPicture,
    SetGuildName,
    SetInvites,
    SetUser
} from '../redux/Dispatches';
import { IMessage, Actions, IState } from '../types/redux';

export function useSocketHandler(socket: HarmonySocket, history: h.History<any>) {
    const dispatch = useDispatch();
    const { currentGuild, channels, invites } = useSelector((state: IState) => state);
    let firstDisconnect = useRef(true);

    useEffect(() => {
        socket.events.addListener('getguilds', (raw: any) => {
            console.log(raw);
            if (typeof raw['guilds'] === 'object') {
                const guildsList = raw['guilds'] as IGuildsList;
                if (Object.keys(guildsList).length === 0) {
                    dispatch(SetMessages([]));
                    dispatch(SetSelectedChannel(undefined));
                    dispatch(SetCurrentGuild(undefined));
                    dispatch(SetChannels({}));
                }
                dispatch(SetGuilds(guildsList));
            }
        });
        socket.events.addListener('getmessages', (raw: any) => {
            if (Array.isArray(raw['messages'])) {
                dispatch(SetMessages((raw['messages'] as IMessage[]).reverse()));
            }
        });
        socket.events.addListener('getchannels', (raw: any) => {
            if (typeof raw === 'object') {
                dispatch(SetChannels(raw['channels']));
            }
        });
        socket.events.addListener('message', (raw: any) => {
            if (
                typeof raw['userid'] === 'string' &&
                typeof raw['createdat'] === 'number' &&
                typeof raw['guild'] === 'string' &&
                typeof raw['message'] === 'string'
            ) {
                dispatch(AddMessage(raw as IMessage));
            }
        });
        socket.events.addListener('leaveguild', (raw: any) => {
            socket.getGuilds();
        });
        socket.events.addListener('joinguild', (raw: any) => {
            socket.getGuilds();
            dispatch({ type: Actions.TOGGLE_JOIN_GUILD_DIALOG });
        });
        socket.events.addListener('createguild', (raw: any) => {
            socket.getGuilds();
            dispatch({ type: Actions.TOGGLE_JOIN_GUILD_DIALOG });
        });
        socket.events.addListener('updateguildpicture', (raw: any) => {
            if (typeof raw['picture'] === 'string' && typeof raw['guild'] === 'string') {
                dispatch(SetGuildPicture(raw['guild'], raw['picture']));
            }
        });
        socket.events.addListener('updateguildname', (raw: any) => {
            if (typeof raw['name'] === 'string' && typeof raw['guild'] === 'string') {
                dispatch(SetGuildName(raw['guild'], raw['name']));
            }
        });
        socket.events.addListener('getinvites', (raw: any) => {
            if (typeof raw['invites'] === 'object') {
                dispatch(SetInvites(raw['invites']));
            }
        });
        socket.events.addListener('addchannel', (raw: any) => {
            if (
                typeof raw['guild'] === 'string' &&
                typeof raw['channelname'] === 'string' &&
                raw['channelid'] === 'string' &&
                raw['guild'] === currentGuild
            ) {
                dispatch(
                    SetChannels({
                        ...channels,
                        [raw['channelid']]: raw['name']
                    })
                );
            }
        });
        socket.events.addListener('deletechannel', (raw: any) => {
            if (typeof raw['guild'] === 'string' && typeof raw['channelid'] === 'string') {
                const deletedChannels = {
                    ...channels
                };
                delete deletedChannels[raw['channelid']];
                dispatch(SetChannels(deletedChannels));
            }
        });
        socket.events.addListener('createinvite', (raw: any) => {
            if (typeof raw['invite'] === 'string') {
                dispatch(
                    SetInvites({
                        ...invites,
                        [raw['invite']]: 0
                    })
                );
            }
        });
        socket.events.addListener('deleteinvite', (raw: any) => {
            if (typeof raw['invite'] === 'string') {
                const deletedInvites = {
                    ...invites
                };
                delete deletedInvites[raw['invite']];
                dispatch(
                    SetInvites({
                        ...deletedInvites
                    })
                );
            }
        });
        socket.events.addListener('getuser', (raw: any) => {
            if (
                typeof raw['userid'] === 'string' &&
                typeof raw['username'] === 'string' &&
                typeof raw['avatar'] === 'string'
            ) {
                dispatch(SetUser(raw['userid'], raw['username'], raw['avatar']));
            }
        });

        socket.events.addListener('deauth', (raw: any) => {
            toast.warn('Your session expired, please login again');
            history.push('/');
        });
        socket.events.addListener('error', (raw: any) => {
            if (typeof raw === 'object' && typeof raw['message'] === 'string') {
                toast.error(raw['message']);
            }
        });
        socket.events.addListener('close', () => {
            if (firstDisconnect.current) {
                firstDisconnect.current = false;
                toast.error('You have lost connection to the server');
            }
        });
        socket.events.addListener('open', () => {
            if (!firstDisconnect.current) {
                toast.success('You have reconnected to the server');
            }
            socket.getGuilds();
            firstDisconnect.current = true;
        });
        console.log('%cSocket Events Bound', 'font-size: x-large');
        if (socket.conn.readyState === WebSocket.OPEN) {
            socket.getGuilds();
        }
        return () => {
            socket.events.removeAllListeners('getguilds');
            socket.events.removeAllListeners('getmessages');
            socket.events.removeAllListeners('getchannels');
            socket.events.removeAllListeners('message');
            socket.events.removeAllListeners('leaveguild');
            socket.events.removeAllListeners('joinguild');
            socket.events.removeAllListeners('createguild');
            socket.events.removeAllListeners('updateguildpicture');
            socket.events.removeAllListeners('updateguildname');
            socket.events.removeAllListeners('getinvites');
            socket.events.removeAllListeners('createinvite');
            socket.events.removeAllListeners('deleteinvite');
            socket.events.removeAllListeners('getuser');
            socket.events.removeAllListeners('deauth');
            socket.events.removeAllListeners('error');
            socket.events.removeAllListeners('open');
            socket.events.removeAllListeners('close');
        };
    }, []);
}
