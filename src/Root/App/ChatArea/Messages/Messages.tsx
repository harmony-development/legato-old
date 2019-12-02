import React from 'react';
import { List } from '@material-ui/core';
import { Message } from './Message';

export const Messages = () => {
    return (
        <List>
            <Message guild='harmony dev' userid='13jifrb' createdat={1575243386} message='Hi' />
        </List>
    );
};
