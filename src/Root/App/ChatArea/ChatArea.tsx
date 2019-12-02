import React from 'react';
import { useChatAreaStyles } from './ChatAreaStyle';
import { Messages } from './Messages/Messages';
import { Input } from './Input/Input';

export const ChatArea = () => {
    const classes = useChatAreaStyles();

    return (
        <div className={classes.root}>
            <div className={classes.messages}>
                <Messages />
            </div>
            <div className={classes.input}>
                <Input />
            </div>
        </div>
    );
};
