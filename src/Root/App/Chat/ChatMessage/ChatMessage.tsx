import React from 'react';
import { useStyles } from './styles';
import { ListItemText, ListItemAvatar, Avatar, Box, ButtonBase } from '@material-ui/core';

interface IProps {
  message: string;
  user: string;
  avatar: string | undefined;
  files: string[];
  openImgDialog: (image: string) => void;
  index: number;
}

const ChatMessage: React.FC<IProps> = (props: IProps) => {
  const classes = useStyles();

  return (
    <>
      <Box display='flex' alignItems='center' className={`${classes.message} ${props.index === 0 ? classes.messageLight : ''}`}>
        <ListItemAvatar>
          <Avatar alt={props.user} src={props.avatar} />
        </ListItemAvatar>
        <ListItemText primary={props.user} secondary={props.message} />
      </Box>
      {props.files && props.files.length > 0 ? (
        <div className={classes.fileSection}>
          {props.files.map((data, index) => {
            return (
              <ButtonBase key={index}>
                <img className={classes.imageUpload} src={data} onClick={(): void => props.openImgDialog(data)} alt='' />
              </ButtonBase>
            );
          })}
        </div>
      ) : (
        undefined
      )}
    </>
  );
};

export default ChatMessage;
