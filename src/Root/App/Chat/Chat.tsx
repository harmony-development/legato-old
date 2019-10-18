import React, { useState, useEffect, useRef } from 'react';
import io from 'socket.io-client';

import { useStyles } from './styles';
import { TextField, Grid, IconButton, Tooltip } from '@material-ui/core';
import ChatMessage from './ChatMessage/ChatMessage';
import { useSelector } from 'react-redux';
import { IAppState } from '../../../store/types';
import { Warning, Image } from '@material-ui/icons';
import { IMessage, Events } from '../../../types';

const socket = io('http://localhost:4000');

const Chat = () => {
  const classes = useStyles(); // CSS styles

  const [messages, setMessages] = useState<IMessage[]>([]); // array of messages to render on screen
  const [draftMessage, setDraftMessage] = useState(''); // the text in the textbox
  const name = useSelector((state: IAppState) => state.name); // the username
  const firstRender = useRef(true); // whether it's the first render or not
  const inputFile = useRef<HTMLInputElement>(null);
  const [connected, setConnected] = useState(false);

  const chatBoxKeyEvent = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      event.preventDefault();
      if (!event.shiftKey && socket !== null) {
        let message = (event.target as HTMLInputElement).value;
        socket.emit(Events.MESSAGE, { author: name, message, files: [] });
        (event.target as HTMLInputElement).selectionEnd = 0;
        setDraftMessage('');
      } else {
        setDraftMessage(draftMessage + '\n');
      }
    }
  };

  const sendFile = () => {
    if (inputFile.current) {
      inputFile.current.click();
    }
  };

  const onFileSelected = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.currentTarget.files && event.currentTarget.files.length > 0) {
      const imageReader = new FileReader();
      imageReader.readAsDataURL(event.currentTarget.files[0]);
      imageReader.addEventListener('load', () => {
        socket.emit(Events.MESSAGE, {
          message: 'Image Upload',
          author: name,
          files: [imageReader.result]
        });
      });
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
      setMessages((prevMessages) => [
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
        <Grid>
          {messages.map((message, index) => (
            <ChatMessage index={index % 2} user={message.author} files={message.files} message={message.message} />
          ))}
        </Grid>
      </div>
      <div className={classes.messageBoxContainer}>
        <div className={classes.valign}>
          <input type='file' id='file' ref={inputFile} style={{ display: 'none' }} onChange={onFileSelected} />
          <Tooltip title='Send Image'>
            <IconButton onClick={sendFile}>
              <Image />
            </IconButton>
          </Tooltip>
        </div>
        <TextField onKeyPress={chatBoxKeyEvent} value={draftMessage} onChange={(event) => setDraftMessage(event.target.value)} className={classes.chatBox} label={connected ? 'Send Message' : 'Currently Offline'} multiline rows='2' margin='normal' />
        {!connected ? (
          <div className={classes.valign}>
            <Tooltip title='Currently Offline. You will not be able to send messages.'>
              <IconButton>
                <Warning />
              </IconButton>
            </Tooltip>
          </div>
        ) : (
          undefined
        )}
      </div>
    </div>
  );
};

export default Chat;
