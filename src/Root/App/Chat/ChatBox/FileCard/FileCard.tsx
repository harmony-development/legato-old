import React from 'react';
import { Box, IconButton } from '@material-ui/core';
import { useStyles } from './styles';
import { Delete } from '@material-ui/icons';

interface IProps {
    image: string;
    removeFromQueue: (index: number) => void;
    index: number;
}

const FileCard: React.FC<IProps> = (props: IProps) => {
    const classes = useStyles();

    const handleDelete = (): void => {
        props.removeFromQueue(props.index);
    };

    return (
        <Box className={classes.root}>
            <div className={classes.overlay}>
                <IconButton className={classes.deleteButton} onClick={handleDelete}>
                    <Delete />
                </IconButton>
            </div>
            <img src={props.image} className={classes.thumbnail} alt=''></img>
        </Box>
    );
};

export default FileCard;
