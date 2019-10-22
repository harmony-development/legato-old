import React, { useState, useEffect, useRef } from 'react';

import { useStyles } from './styles';
import { Box } from '@material-ui/core';
import ChatMessage from './ChatMessage/ChatMessage';
import { useSelector } from 'react-redux';
import { IAppState } from '../../../store/types';
import { IMessage } from '../../../types';
import ChatBox from './ChatBox/ChatBox';
import { socketServer } from '../../Root';
import { Events } from '../../../socket/socket';
import ImageDialog from './ImageDialog/ImageDialog';

const Chat: React.FC<{}> = () => {
  const classes = useStyles();

  const [messages, setMessages] = useState<IMessage[]>([]);
  const [imageDialogOpen, setImageDialogOpen] = useState(false);
  const [previewImage, setPreviewImage] = useState('');
  const user = useSelector((state: IAppState) => state.user);
  const MessagesArea = useRef<HTMLDivElement>(null);

  const openImageDialog = (image: string): void => {
    setPreviewImage(image);
    setImageDialogOpen(true);
  };

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
      <ImageDialog open={imageDialogOpen} setOpen={setImageDialogOpen} image={previewImage} />
      <div className={classes.chatBoxContainer} ref={MessagesArea}>
        <Box>
          {messages.map((message, index) => (
            <ChatMessage key={index} index={index % 2} user={message.username} avatar={message.avatar || undefined} files={message.files} message={message.message} openImgDialog={openImageDialog} />
          ))}
        </Box>
      </div>
      <Box>
        <ChatBox name={user.username} />
      </Box>
    </div>
  );
};

export default Chat;
