import React, { useEffect } from 'react';
import { HarmonyBar } from './HarmonyBar/HarmonyBar';
import { ThemeDialog } from './Dialog/ThemeDialog';
import { useAppStyles } from './AppStyle';
import { ChatArea } from './ChatArea/ChatArea';
import { harmonySocket } from '../Root';
import { useHistory } from 'react-router';
import { IGuildData } from '../../types/socket';

export const App = () => {
    const classes = useAppStyles();
    const history = useHistory();

    useEffect(() => {
        if ((harmonySocket.conn.readyState !== WebSocket.OPEN && harmonySocket.conn.readyState !== WebSocket.CONNECTING) || typeof localStorage.getItem('token') !== 'string') {
            // bounce the user to the login screen if the socket is disconnected or there's no token
            history.push('/');
            return;
        }
        harmonySocket.events.addListener('getguilds', (raw: any) => {
            console.log(raw);
            if (Array.isArray(raw['guilds']) && raw['guilds'].length > 0) {
                const guildsList = raw['guilds'] as IGuildData[];
                guildsList.forEach((guild) => {
                    console.log(guild.guildid);
                });
            }
        });
        const guildReqID = setInterval(() => {
            if (harmonySocket.conn.readyState === WebSocket.OPEN) {
                harmonySocket.getGuilds();
                clearInterval(guildReqID);
            }
        }, 500);
    }, [history]);

    return (
        <div className={classes.root}>
            <ThemeDialog />
            <HarmonyBar />
            <div className={classes.navFill} /> {/* this fills the area where the navbar is*/}
            <ChatArea />
        </div>
    );
};
