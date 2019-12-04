import React, { useEffect, useRef } from 'react';
import { useChatAreaStyles } from './ChatAreaStyle';
import { Messages } from './Messages/Messages';
import { Input } from './Input/Input';
import { GuildList } from './GuildList/GuildList';
import { useSelector } from 'react-redux';
import { IState } from '../../../types/redux';

export const ChatArea = () => {
    const classes = useChatAreaStyles();
    const messages = useSelector((state: IState) => state.messages);
    const messagesRef = useRef<HTMLDivElement | null>(null);

    useEffect(() => {
        if (messagesRef.current) {
            console.log(messagesRef.current);
            messagesRef.current.scrollTop = messagesRef.current.scrollHeight;
        }
    }, [messages]);

    return (
        <div className={classes.root}>
            <div className={classes.guildlist}>
                <GuildList />
            </div>
            <div className={classes.chatArea}>
                <div className={classes.messages} ref={messagesRef}>
                    <Messages />
                </div>
                <div className={classes.input}>
                    <Input />
                </div>
            </div>
        </div>
    );
};
