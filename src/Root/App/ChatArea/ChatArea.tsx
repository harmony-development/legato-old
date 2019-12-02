import React from 'react';
import { useChatAreaStyles } from './ChatAreaStyle';
import { Messages } from './Messages/Messages';
import { Input } from './Input/Input';
import { ListItem, List, ListItemAvatar, Avatar, ButtonBase } from '@material-ui/core';

export const ChatArea = () => {
    const classes = useChatAreaStyles();

    return (
        <div className={classes.root}>
            <div className={classes.guildlist}>
                <ButtonBase className={classes.guildiconroot}>
                    <img className={classes.guildicon} alt='' src='https://gitlab.com/uploads/-/system/group/avatar/6322234/harmony-icon.png?width=64' />
                </ButtonBase>
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
