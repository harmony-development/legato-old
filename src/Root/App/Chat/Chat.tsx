import React, { useState, useEffect, useRef } from 'react';
import io from 'socket.io-client';

import { useStyles } from './styles';
import { TextField, List, IconButton, Grid, Tooltip } from '@material-ui/core';
import ChatMessage from './ChatMessage/ChatMessage';
import { useSelector } from 'react-redux';
import { IAppState, Events } from '../../../store/types';
import { Warning } from '@material-ui/icons';

interface IMessage {
  author: string;
  message: string;
}

const socket = io('http://localhost:4000');

const Chat = () => {
  const classes = useStyles(); // CSS styles

  const [messages, setMessages] = useState<IMessage[]>([]); // array of messages to render on screen
  const [draftMessage, setDraftMessage] = useState(''); // the text in the textbox
  const name = useSelector((state: IAppState) => state.name); // the username
  const firstRender = useRef(true); // whether it's the first render or not
  const [connected, setConnected] = useState(false);

  const chatBoxKeyEvent = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      event.preventDefault();
      if (!event.shiftKey && socket !== null) {
        let message = (event.target as HTMLInputElement).value;
        socket.emit(Events.MESSAGE, { message });
        (event.target as HTMLInputElement).selectionEnd = 0;
        setDraftMessage('');
      } else {
        setDraftMessage(draftMessage + '\n');
      }
    }
  };

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
      setMessages((prevMessages) => [...prevMessages, { author: response.author, message: response.message }]);
    });
  }, []);

  return (
    <div className={classes.container}>
      <div className={classes.chatBoxContainer}>
        <List>
          {messages.map((message, index) => (
            <ChatMessage index={index % 2} user={message.author} message={message.message} />
          ))}
        </List>
      </div>
      <Grid container className={classes.messageBoxContainer}>
        <TextField onKeyPress={chatBoxKeyEvent} value={draftMessage} onChange={(event) => setDraftMessage(event.target.value)} className={classes.chatBox} label={connected ? 'Send Message' : 'Currently Offline'} multiline rows='2' margin='normal' />
        {!connected ? (
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <Tooltip title='Currently Offline. You will not be able to send messages.'>
              <IconButton>
                <Warning />
              </IconButton>
            </Tooltip>
          </div>
        ) : (
          undefined
        )}
      </Grid>
    </div>
  );
};

export default Chat;
