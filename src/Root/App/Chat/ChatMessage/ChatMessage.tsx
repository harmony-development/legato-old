import React from 'react';
import { useStyles } from './styles';
import { Typography, ListItem, ListItemText, Divider, Grid } from '@material-ui/core';

interface IProps {
  message: string;
  user: string;
  files: string[];
  index: number;
}

export default function ChatMessage(props: IProps) {
  const classes = useStyles();

  return (
    <>
      <Grid item alignItems='flex-start' className={`${classes.message} ${props.index === 0 ? classes.messageLight : ''}`}>
        <ListItemText
          className={classes.fileSection}
          primary={
            <Typography component='span' variant='body2' className={classes.userText} color='textPrimary'>
              {props.user}
            </Typography>
          }
          secondary={
            <Typography component='span' variant='body2' className={classes.messageText} color='textSecondary'>
              {props.message}
            </Typography>
          }
        />
        <div className={classes.fileSection}>
          {props.files.map((data) => {
            return <img className={classes.imageUpload} src={data}></img>;
          })}
        </div>
      </Grid>
    </>
  );
}
