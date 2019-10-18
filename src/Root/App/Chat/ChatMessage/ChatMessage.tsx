import React from 'react';
import { useStyles } from './styles';
import { Typography, ListItem, ListItemText, Divider } from '@material-ui/core';

interface IProps {
  message: string;
  user: string;
  files: File[];
  index: number;
}

export default function ChatMessage(props: IProps) {
  const classes = useStyles();

  return (
    <>
      <ListItem
        alignItems='flex-start'
        className={props.index === 0 ? classes.messageLight : ''}
      >
        <ListItemText
          primary={
            <Typography
              component='span'
              variant='body2'
              className={classes.userText}
              color='textPrimary'
            >
              {props.user}
            </Typography>
          }
          secondary={
            <Typography
              component='span'
              variant='body2'
              className={classes.messageText}
              color='textSecondary'
            >
              {props.message}
            </Typography>
          }
        />
        {
          props.files.map((data) => {
            <img
          })
        }
      </ListItem>
      <Divider variant='fullWidth' component='li' />
    </>
  );
}
