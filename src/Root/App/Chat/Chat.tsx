import React, { useState, useEffect, useRef } from 'react';
import io from 'socket.io-client';

import { useStyles } from './styles';
import { TextField, List } from '@material-ui/core';
import ChatMessage from './ChatMessage/ChatMessage';
import { useSelector } from 'react-redux';
import { IAppState } from '../../../store/types';

interface IMessage {
  user: string;
  message: string;
}

interface IJoinResponse {
  userid: string;
}

const socket = io('http://localhost:4000');

const Chat = () => {
  const classes = useStyles(); // CSS styles

  const [messages, setMessages] = useState<IMessage[]>([]); // array of messages to render on screen
  const [draftMessage, setDraftMessage] = useState(''); // the text in the textbox
  const name = useSelector((state: IAppState) => state.name); // the username
  const firstRender = useRef(true); // whether it's the first render or not

  const chatBoxKeyEvent = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      event.preventDefault();
      if (!event.shiftKey && socket !== null) {
        let message = (event.target as HTMLInputElement).value;
        socket.emit('message', { message });
        (event.target as HTMLInputElement).selectionEnd = 0;
        setDraftMessage('');
      } else {
        setDraftMessage(draftMessage + '\n');
      }
    }
  };

  useEffect(() => {
    if (!firstRender.current) {
      socket.emit('UsernameUpdate', name);
    }
  }, [name]);

  useEffect(() => {
    firstRender.current = false;
    socket.on('connect', () => {
      socket.emit('ClientConnect', { name });
    });
    socket.on('ClientConnectEvent', (response: IJoinResponse) => {
      setMessages((prevMessages) => [...prevMessages, { user: response.userid, message: ' has joined the channel' }]);
    });
    socket.on('ClientDisconnectEvent', (response: IJoinResponse) => {
      setMessages((prevMessages) => [...prevMessages, { user: response.userid, message: ' has left the channel' }]);
    });
    socket.on('message', (response: IMessage) => {
      setMessages((prevMessages) => [...prevMessages, { user: response.user, message: response.message }]);
    });
  }, []);

  return (
    <div className={classes.container}>
      <div className={classes.chatBoxContainer}>
        <List>
          {messages.map((message, index) => (
            <ChatMessage index={index % 2} user={message.user} message={message.message} />
          ))}
        </List>
      </div>
      <div className={classes.messageBoxContainer}>
        <TextField onKeyPress={chatBoxKeyEvent} value={draftMessage} onChange={(event) => setDraftMessage(event.target.value)} className={classes.messageBox} label='Send Message' multiline rows='2' margin='normal' />
      </div>
    </div>
  );
};

export default Chat;
