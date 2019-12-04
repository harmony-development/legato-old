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

export const App = () => {
    const classes = useAppStyles();
    const dispatch = useDispatch();
    const connected = useSelector((state: IState) => state.connected);
    const selectedGuild = useSelector((state: IState) => state.selectedGuild);
    const history = useHistory();

    // event when the client has connected
    useEffect(() => {
        if (connected) {
            harmonySocket.getGuilds();
        }
    }, [connected]);

    useEffect(() => {
        if (connected) {
            console.log(selectedGuild);
            harmonySocket.getMessages(selectedGuild);
        }
    });

    useEffect(() => {
        if ((harmonySocket.conn.readyState !== WebSocket.OPEN && harmonySocket.conn.readyState !== WebSocket.CONNECTING) || typeof localStorage.getItem('token') !== 'string') {
            // bounce the user to the login screen if the socket is disconnected or there's no token
            history.push('/');
            return;
        }
        harmonySocket.events.addListener('getguilds', (raw: any) => {
            if (Object.keys(raw['guilds']).length) {
                const guildsList = raw['guilds'] as IGuildData[];
                dispatch({ type: Actions.SET_GUILDS, payload: guildsList });
            }
        });

        harmonySocket.events.addListener('message', (raw: any) => {
            if (Array.isArray(raw)) {
                dispatch({ type: Actions.SET_MESSAGES, payload: raw as IMessage[] });
            }
        });

        return () => {
            harmonySocket.events.removeAllListeners('getguilds');
            harmonySocket.events.removeAllListeners('message');
        };
    }, [history, dispatch]);

    return (
        <div className={classes.root}>
            <ThemeDialog />
            <HarmonyBar />
            <div className={classes.navFill} /> {/* this fills the area where the navbar is*/}
            <ChatArea />
        </div>
    );
};
