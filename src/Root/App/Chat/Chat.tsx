import React, { useState, useEffect, useRef } from 'react';

import { useStyles } from './styles';
import { Box } from '@material-ui/core';
import ChatMessage from './ChatMessage/ChatMessage';
import { useSelector } from 'react-redux';
import { IAppState } from '../../../store/types';
import { IMessage, IUserData, IGetTargetUserData } from '../../../types';
import ChatBox from './ChatBox/ChatBox';
import { socketServer } from '../../Root';
import { Events } from '../../../socket/socket';
import ImageDialog from './ImageDialog/ImageDialog';
import { TransitionGroup, CSSTransition } from 'react-transition-group';

let scrollTrigger = false;
let preScrollHeight = 0;
let preScrollTop = 0;
let gettingMessages = false;
const Chat: React.FC<{}> = () => {
  const classes = useStyles();

  const [messages, setMessages] = useState<IMessage[]>([]);
  const [userData, setUserData] = useState<IUserData>({});
  const [imageDialogOpen, setImageDialogOpen] = useState(false);
  const [previewImage, setPreviewImage] = useState('');
  const user = useSelector((state: IAppState) => state.user);
  const MessagesArea = useRef<HTMLDivElement>(null);

  const openImageDialog = (image: string): void => {
    setPreviewImage(image);
    setImageDialogOpen(true);
  };

  const onScroll = (event: React.UIEvent<HTMLDivElement>): void => {
    if (
      typeof localStorage.getItem('token') === 'string' &&
      messages[0] &&
      event.currentTarget &&
      !gettingMessages
    ) {
      if (event.currentTarget.scrollTop === 0) {
        scrollTrigger = true;
        preScrollHeight = event.currentTarget.scrollHeight;
        preScrollTop = event.currentTarget.scrollTop;
        socketServer.connection.emit(Events.LOAD_MESSAGES, {
          token: localStorage.getItem('token'),
          lastmessageid: messages[0].messageid || undefined
        });
        gettingMessages = true;
      }
    }
  };

  useEffect(() => {
    if (MessagesArea.current && !scrollTrigger) {
      MessagesArea.current.scrollTop = MessagesArea.current.scrollHeight;
    }
    scrollTrigger = false;
  }, [messages]);

  useEffect(() => {
    if (typeof localStorage.getItem('userData') === 'string') {
      setUserData(JSON.parse(localStorage.getItem('userData') as string));
    }

    if (typeof localStorage.getItem('token') === 'string') {
      socketServer.connection.emit(Events.LOAD_MESSAGES, {
        token: localStorage.getItem('token')
      });
    }

    socketServer.connection.on(Events.MESSAGE, (newMessage: IMessage) => {
      setMessages((oldMessages: IMessage[]) => [...oldMessages, newMessage]);
    });

    socketServer.connection.on(
      Events.LOAD_MESSAGES,
      (loadedMessages: IMessage[]) => {
        if (MessagesArea.current) {
          setMessages((oldMessages: IMessage[]) => [
            ...loadedMessages,
            ...oldMessages
          ]);
          if (MessagesArea.current) {
            MessagesArea.current.scrollTop =
              MessagesArea.current.scrollHeight -
              preScrollHeight +
              preScrollTop;
          }
          gettingMessages = false;
        }
      }
    );

    socketServer.connection.on(Events.LOAD_MESSAGES_ERROR, (err) =>
      console.log(err)
    );

    socketServer.connection.on(
      Events.GET_TARGET_USER_DATA,
      (data: IGetTargetUserData) => {
        if (data) {
          if (data.userid) {
            setUserData((oldUserData: IUserData) => ({
              ...oldUserData,
              [data.userid]: { avatar: data.avatar, username: data.username }
            }));
          }
        }
      }
    );

    return (): void => {
      socketServer.connection.removeListener(Events.MESSAGE);
      socketServer.connection.removeListener(Events.GET_TARGET_USER_DATA);
      socketServer.connection.removeListener(Events.LOAD_MESSAGES);
      socketServer.connection.removeListener(Events.LOAD_MESSAGES_ERROR);
    };
  }, []);

  useEffect(() => {
    localStorage.setItem('userData', JSON.stringify(userData));
  }, [userData]);

  return (
    <div className={classes.container}>
      <ImageDialog
        open={imageDialogOpen}
        setOpen={setImageDialogOpen}
        image={previewImage}
      />
      <div
        className={classes.chatBoxContainer}
        ref={MessagesArea}
        onScroll={onScroll}
      >
        {messages.map((message, index) => (
          <ChatMessage
            index={index % 2}
            key={index}
            userid={message.author}
            userData={userData}
            setUserData={setUserData}
            files={message.files}
            message={message.message}
            openImgDialog={openImageDialog}
          />
        ))}
      </div>
      <Box>
        <ChatBox name={user.username} />
      </Box>
    </div>
  );
};

export default Chat;
