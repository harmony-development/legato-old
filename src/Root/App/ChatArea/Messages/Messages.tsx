import React, { useEffect, useRef } from 'react';
import { List } from '@material-ui/core';
import { Message } from './Message';
import { useSelector } from 'react-redux';
import { IState } from '../../../../types/redux';

export const Messages = () => {
    const messages = useSelector((state: IState) => state.messages);
    const selectedChannel = useSelector((state: IState) => state.selectedChannel);
    const usernames = useSelector((state: IState) => state.usernames);
    const messageList = useRef<HTMLUListElement | undefined>(undefined);

    useEffect(() => {
        if (messageList.current) {
            messageList.current.scrollTop = messageList.current.scrollHeight;
        }
    }, [messages]);

    return (
        <List innerRef={messageList}>
            {messages
                ? messages.map((val) => {
                      if (val.channel === selectedChannel) {
                          return <Message key={val.messageid} guild={val.guild} userid={val.userid} username={usernames[val.userid]} createdat={val.createdat} message={val.message} />;
                      } else {
                          return undefined;
                      }
                  })
                : undefined}
        </List>
    );
};
