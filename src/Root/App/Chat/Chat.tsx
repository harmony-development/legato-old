import React, { useState, useEffect, useRef } from 'react';
import io from 'socket.io-client';

import { useStyles } from './styles';
import { List } from '@material-ui/core';
import ChatMessage from './ChatMessage/ChatMessage';
import { useSelector } from 'react-redux';
import { IAppState } from '../../../store/types';
import { IMessage, Events } from '../../../types';
import ChatBox from './ChatBox/ChatBox';

const socket = io('http://localhost:4000');

const Chat: React.FC<{}> = () => {
  const classes = useStyles(); // CSS styles

  const [messages, setMessages] = useState<IMessage[]>([]); // array of messages to render on screen
  const [connected, setConnected] = useState(false);
  const name = useSelector((state: IAppState) => state.name); // the username
  const firstRender = useRef(true); // whether it's the first render or not

  useEffect(() => {
    if (!firstRender.current) {
      socket.emit(Events.USERNAME_UPDATE, { name });
    }
  }, [name]);

  useEffect(() => {
    firstRender.current = false;
    socket.on('connect', () => {
      setConnected(true);
      socket.emit(Events.LOGIN, { name });
    });
    socket.on('disconnect', () => {
      setConnected(false);
    });
    socket.on(Events.MESSAGE, (response: IMessage) => {
      setMessages(prevMessages => [
        ...prevMessages,
        {
          author: response.author,
          message: response.message,
          files: response.files
        }
      ]);
    });
  }, []);

  return (
    <div className={classes.container}>
      <div className={classes.chatBoxContainer}>
        <List>
          {messages.map((message, index) => (
            <ChatMessage
              key={index}
              index={index % 2}
              user={message.author}
              files={message.files}
              message={message.message}
            />
          ))}
        </List>
      </div>
      <ChatBox socket={socket} connected={connected} name={name} />
    </div>
  );
};

export default Chat;
