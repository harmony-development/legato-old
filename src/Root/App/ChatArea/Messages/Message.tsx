import React, { useEffect } from 'react';
import { ListItem, ListItemAvatar, Avatar, ListItemText, Typography } from '@material-ui/core';
import { harmonySocket } from '../../../Root';

interface IProps {
    guild: string;
    userid: string;
    username: string;
    createdat: number;
    message: string;
    avatar?: string;
}

const UtcEpochToLocalDate = (time: number) => {
    const returnDate = new Date(0);
    returnDate.setUTCSeconds(time);
    return ` - ${returnDate.toDateString()} at ${returnDate.toLocaleTimeString()}`;
};

export const Message = (props: IProps) => {
    useEffect(() => {
        if (!props.username) {
            harmonySocket.sendGetUser(props.userid);
        }
    }, [props]);

    return (
        <ListItem alignItems='flex-start'>
            <ListItemAvatar>
                <Avatar alt={props.userid} src={props.avatar ? `http://localhost:2288/filestore/${props.avatar}` : undefined} />
            </ListItemAvatar>
            <ListItemText
                primary={
                    <>
                        {props.username || props.userid}
                        <Typography component='span' variant='body1' color='textSecondary'>
                            {UtcEpochToLocalDate(props.createdat)}
                        </Typography>
                    </>
                }
                secondary={props.message}
            />
        </ListItem>
    );
};
