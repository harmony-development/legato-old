import React from 'react';
import { Paper, CardMedia, Grid, Box, IconButton } from '@material-ui/core';
import { useStyles } from './styles';
import { Remove, Delete } from '@material-ui/icons';

interface IProps {
  image: string;
  removeFromQueue: (index: number) => void;
  index: number;
}

const FileCard = (props: IProps) => {
  const classes = useStyles();

  const handleDelete = () => {
    props.removeFromQueue(props.index);
  };

  return (
    <Box className={classes.root}>
      <div className={classes.overlay}>
        <IconButton className={classes.deleteButton} onClick={handleDelete}>
          <Delete />
        </IconButton>
      </div>
      <img src={props.image} className={classes.thumbnail}></img>
    </Box>
  );
};

export default FileCard;
