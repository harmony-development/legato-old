import React, { useState, useEffect, useRef } from 'react';

import { useStyles } from './styles';
import { Box } from '@material-ui/core';
import ChatMessage from './ChatMessage/ChatMessage';
import { useSelector } from 'react-redux';
import { IAppState } from '../../../store/types';
import { IMessage, Events } from '../../../types';
import ChatBox from './ChatBox/ChatBox';
import { socketServer } from '../../Root';

const Chat: React.FC<{}> = () => {
  const classes = useStyles(); // CSS styles

  const [messages, setMessages] = useState<IMessage[]>([]); // array of messages to render on screen
  const user = useSelector((state: IAppState) => state.user); // the username
  const MessagesArea = useRef<HTMLDivElement>(null);

  useEffect(() => {
    socketServer.connection.on(Events.MESSAGE, (newMessage: IMessage) => {
      setMessages((oldMessages: IMessage[]) => [...oldMessages, newMessage]);
    });

    return (): void => {
      socketServer.connection.removeListener(Events.MESSAGE);
    };
  }, []);

  useEffect(() => {
    if (MessagesArea.current) {
      MessagesArea.current.scrollTop = MessagesArea.current.scrollHeight;
    }
  }, [messages]);

  return (
    <div className={classes.container}>
      <div className={classes.chatBoxContainer} ref={MessagesArea}>
        <Box>
          {messages.map((message, index) => (
            <ChatMessage
              key={index}
              index={index % 2}
              user={message.author}
              usericon={message.avatar || undefined}
              files={message.files}
              message={message.message}
            />
          ))}
        </Box>
      </div>
      <ChatBox name={user.name} />
    </div>
  );
};

export default Chat;
