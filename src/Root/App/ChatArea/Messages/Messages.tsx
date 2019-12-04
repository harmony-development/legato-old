import React, { useEffect, useRef } from 'react';
import { List } from '@material-ui/core';
import { Message } from './Message';
import { useSelector } from 'react-redux';
import { IState } from '../../../../types/redux';

export const Messages = () => {
    const messages = useSelector((state: IState) => state.messages);
    const messageList = useRef<HTMLUListElement | undefined>(undefined);

    useEffect(() => {
        if (messageList.current) {
            console.log(messageList.current);
            messageList.current.scrollTop = messageList.current.scrollHeight;
        }
    }, [messages]);

    return (
        <List innerRef={messageList}>
            {messages.map((val) => {
                return <Message key={val.messageid} guild={val.guild} userid={val.userid} createdat={val.createdat} message={val.message} />;
            })}
        </List>
    );
};
