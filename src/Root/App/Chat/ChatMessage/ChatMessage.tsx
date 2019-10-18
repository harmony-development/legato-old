import React from 'react';
import { useStyles } from './styles';
import { Typography, ListItem, ListItemText, Divider } from '@material-ui/core';

interface IProps {
  message: string;
  user: string;
  files: string[];
  index: number;
}

const ChatMessage: React.FC<IProps> = (props: IProps) => {
  const classes = useStyles();

  return (
    <>
      <ListItem
        alignItems="flex-start"
        className={`${classes.message} ${
          props.index === 0 ? classes.messageLight : ''
        }`}
      >
        <ListItemText
          primary={
            <Typography
              component="span"
              variant="body2"
              className={classes.userText}
              color="textPrimary"
            >
              {props.user}
            </Typography>
          }
          secondary={
            <Typography
              component="span"
              variant="body2"
              className={classes.messageText}
              color="textSecondary"
            >
              {props.message}
            </Typography>
          }
        />
        <Divider variant="middle" />
        <div className={classes.fileSection}>
          {props.files.map(data => {
            return (
              <img key={data} className={classes.imageUpload} src={data}></img>
            );
          })}
        </div>
      </ListItem>
    </>
  );
};

export default ChatMessage;
