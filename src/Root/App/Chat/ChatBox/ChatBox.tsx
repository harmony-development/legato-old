import React, { useRef, useState } from 'react';
import { useStyles } from './styles';
import { Tooltip, IconButton, TextField } from '@material-ui/core';
import { Warning, Image } from '@material-ui/icons';
import { Events } from '../../../../types';

interface IProps {
  socket: SocketIOClient.Socket;
  connected: boolean;
  name: string;
}

const ChatBox: React.FC<IProps> = (props: IProps) => {
  const classes = useStyles();
  const [draftMessage, setDraftMessage] = useState('');
  const inputFile = useRef<HTMLInputElement>(null);

  const sendFile = (): void => {
    if (inputFile.current) {
      inputFile.current.click();
    }
  };

  const onFileSelected = (event: React.ChangeEvent<HTMLInputElement>): void => {
    if (event.currentTarget.files && event.currentTarget.files.length > 0) {
      const imageReader = new FileReader();
      imageReader.readAsDataURL(event.currentTarget.files[0]);
      imageReader.addEventListener('load', () => {
        props.socket.emit(Events.MESSAGE, {
          message: 'Image Upload',
          author: props.name,
          files: [imageReader.result]
        });
      });
    }
  };

  const chatBoxKeyEvent = (
    event: React.KeyboardEvent<HTMLInputElement>
  ): void => {
    if (event.key === 'Enter') {
      event.preventDefault();
      if (!event.shiftKey) {
        const message = (event.target as HTMLInputElement).value;
        props.socket.emit(Events.MESSAGE, {
          author: props.name,
          message,
          files: []
        });
        (event.target as HTMLInputElement).selectionEnd = 0;
        setDraftMessage('');
      } else {
        setDraftMessage(draftMessage + '\n');
      }
    }
  };

  return (
    <div className={classes.messageBoxContainer}>
      <div className={classes.valign}>
        <input
          type="file"
          id="file"
          ref={inputFile}
          style={{ display: 'none' }}
          onChange={onFileSelected}
        />
        <Tooltip title="Send Image">
          <IconButton onClick={sendFile}>
            <Image />
          </IconButton>
        </Tooltip>
      </div>
      <TextField
        onKeyPress={chatBoxKeyEvent}
        value={draftMessage}
        onChange={(event): void => setDraftMessage(event.target.value)}
        className={classes.chatBox}
        label={props.connected ? 'Send Message' : 'Currently Offline'}
        multiline
        rows="2"
        margin="normal"
      />
      {!props.connected ? (
        <div className={classes.valign}>
          <Tooltip title="Currently Offline. You will not be able to send messages.">
            <IconButton>
              <Warning />
            </IconButton>
          </Tooltip>
        </div>
      ) : (
        undefined
      )}
    </div>
  );
};

export default ChatBox;
