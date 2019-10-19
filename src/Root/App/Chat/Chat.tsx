import React, { useState, useEffect, useRef } from 'react';
import io from 'socket.io-client';

import { useStyles } from './styles';
import { Box } from '@material-ui/core';
import ChatMessage from './ChatMessage/ChatMessage';
import { useSelector } from 'react-redux';
import { IAppState } from '../../../store/types';
import { IMessage, Events } from '../../../types';
import ChatBox from './ChatBox/ChatBox';

const socket = io('http://0.0.0.0:4000');

const Chat: React.FC<{}> = () => {
  const classes = useStyles(); // CSS styles

  const [messages, setMessages] = useState<IMessage[]>([]); // array of messages to render on screen
  const [connected, setConnected] = useState(false);
  const user = useSelector((state: IAppState) => state.user); // the username
  const firstRender = useRef(true); // whether it's the first render or not
  const MessagesArea = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!firstRender.current) {
      socket.emit(Events.PROFILE_UPDATE, { name: user.name, icon: user.icon });
    }
  }, [user]);

  useEffect(() => {
    if (MessagesArea.current) {
      MessagesArea.current.scrollTop = MessagesArea.current.scrollHeight;
    }
  }, [messages]);

  useEffect(() => {
    firstRender.current = false;
    socket.on('connect', () => {
      setConnected(true);
      socket.emit(Events.LOGIN, { name: user.name });
    });
    socket.on('disconnect', () => {
      setConnected(false);
    });
    socket.on(Events.MESSAGE, (response: IMessage) => {
      setMessages((prevMessages) => [
        ...prevMessages,
        {
          author: response.author,
          icon: response.icon,
          message: response.message,
          files: response.files
        }
      ]);
    });
  }, []);

  return (
    <div className={classes.container}>
      <div className={classes.chatBoxContainer} ref={MessagesArea}>
        <Box>
          {messages.map((message, index) => (
            <ChatMessage key={index} index={index % 2} user={message.author} usericon={message.icon || undefined} files={message.files} message={message.message} />
          ))}
        </Box>
      </div>
      <ChatBox socket={socket} connected={connected} name={user.name} />
    </div>
  );
};

export default Chat;
