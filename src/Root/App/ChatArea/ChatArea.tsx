import React, { useEffect, useRef } from 'react';
import { useChatAreaStyles } from './ChatAreaStyle';
import { Messages } from './Messages/Messages';
import { Input } from './Input/Input';
import { GuildList } from './GuildList/GuildList';
import { useSelector } from 'react-redux';
import { IState } from '../../../types/redux';
import { ChannelList } from './ChannelList/ChannelList';

export const ChatArea = () => {
    const classes = useChatAreaStyles();
    const messages = useSelector((state: IState) => state.messages);
    const messagesRef = useRef<HTMLDivElement | null>(null);
    const chatInput = useSelector((state: IState) => state.chatInput);

    useEffect(() => {
        if (messagesRef.current) {
            messagesRef.current.scrollTop = messagesRef.current.scrollHeight;
        }
    }, [messages]);

    const onKeyDown = (ev: React.KeyboardEvent<HTMLDivElement>) => {
        console.log('bruh');
        if (ev.key !== 'Tab' && chatInput) {
            chatInput.focus();
        }
    };

    return (
        <div className={classes.root}>
            <div className={classes.guildlist}>
                <GuildList />
            </div>
            <div className={classes.channellist}>
                <ChannelList />
            </div>
            <div className={classes.chatArea}>
                <div className={classes.messages} ref={messagesRef} onKeyDown={onKeyDown} tabIndex={-1}>
                    <Messages />
                </div>
                <div className={classes.input}>
                    <Input />
                </div>
            </div>
        </div>
    );
};
