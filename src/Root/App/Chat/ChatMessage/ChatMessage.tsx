import React from 'react';
import { useStyles } from './styles';
import { Typography, ListItem, ListItemText, Divider, Grid } from '@material-ui/core';

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
      <Grid item alignItems='flex-start' className={`${classes.message} ${props.index === 0 ? classes.messageLight : ''}`}>
        <ListItemText
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
          {props.files
            ? props.files.map((data) => {
                return <img key={data} className={classes.imageUpload} src={data}></img>;
              })
            : undefined}
        </div>
      </Grid>
    </>
  );
};

export default ChatMessage;
