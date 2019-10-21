import React from 'react';
import { useStyles } from './styles';
import { ListItemText, ListItemAvatar, Avatar, Box } from '@material-ui/core';

interface IProps {
  message: string;
  user: string;
  avatar: string | undefined;
  files: string[];
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
          {props.files.map((data) => {
            return <img key={data} className={classes.imageUpload} src={data} />;
          })}
        </div>
      ) : (
        undefined
      )}
    </>
  );
};

export default ChatMessage;
