import React, { useEffect } from 'react';
import { useStyles } from './styles';
import { ListItemText, ListItemAvatar, Avatar, Box, ButtonBase } from '@material-ui/core';
import { socketServer } from '../../../Root';
import { Events } from '../../../../socket/socket';

const ChatMessage = (props) => {
    const classes = useStyles();

    useEffect(() => {
        socketServer.connection.emit(Events.GET_USER_DATA, {
            token: localStorage.getItem('token'),
            targetuser: props.userid
        });
    }, []);

    return (
        <>
            <Box display='flex' alignItems='center' className={`${classes.message} ${props.index === 0 ? classes.messageLight : ''}`}>
                <ListItemAvatar>
                    <Avatar alt={props.userData[props.userid] ? props.userData[props.userid].username : undefined} src={props.userData[props.userid] ? props.userData[props.userid].avatar : undefined} />
                </ListItemAvatar>
                <ListItemText primary={props.userData[props.userid] ? props.userData[props.userid].username : 'Loading Username...'} secondary={props.message} />
            </Box>
            {props.files && props.files.length > 0 ? (
                <div className={classes.fileSection}>
                    {props.files.map((data, index) => {
                        return (
                            <ButtonBase key={index}>
                                <img className={classes.imageUpload} src={data} onClick={() => props.openImgDialog(data)} alt='' />
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
