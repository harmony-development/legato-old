import React from 'react';
import { List } from '@material-ui/core';
import { Message } from './Message';
import { useSelector } from 'react-redux';
import { IState } from '../../../../types/redux';

export const Messages = () => {
    const messages = useSelector((state: IState) => state.messages);

    return (
        <List>
            {messages.map((val) => {
                return <Message guild={val.guild} userid={val.userid} createdat={val.createdat} message={val.message} />;
            })}
        </List>
    );
};
