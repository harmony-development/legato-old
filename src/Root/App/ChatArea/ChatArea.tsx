import React from 'react';
import { useChatAreaStyles } from './ChatAreaStyle';
import { Messages } from './Messages/Messages';
import { Input } from './Input/Input';
import { GuildList } from './GuildList/GuildList';

export const ChatArea = () => {
    const classes = useChatAreaStyles();

    return (
        <div className={classes.root}>
            <div className={classes.guildlist}>
                <GuildList />
            </div>
            <div className={classes.chatArea}>
                <div className={classes.messages}>
                    <Messages />
                </div>
                <div className={classes.input}>
                    <Input />
                </div>
            </div>
        </div>
    );
};
