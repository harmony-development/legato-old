import React, { useState } from 'react';
import { List, ListItem, ListItemText, ListItemIcon, Collapse } from '@material-ui/core';
import SettingsIcon from '@material-ui/icons/Settings';
import ExpandMore from '@material-ui/icons/ExpandMore';
import ExpandLess from '@material-ui/icons/ExpandLess';
import LeaveIcon from '@material-ui/icons/ExitToApp';
import { useSelector, useDispatch } from 'react-redux';
import { IState, Actions } from '../../../../types/redux';
import { useChannelListStyle } from './ChannelListStyle';
import { harmonySocket } from '../../../Root';

interface IChannelProps {
    channelid: string;
    channelname: string;
    setSelectedChannel: (value: string) => void;
}

const Channel = (props: IChannelProps) => {
    const selectedChannel = useSelector((state: IState) => state.selectedChannel);
    const classes = useChannelListStyle();

    return (
        <ListItem button key={props.channelid} className={props.channelid === selectedChannel ? classes.selectedChannel : undefined} onClick={() => props.setSelectedChannel(props.channelid)}>
            <ListItemText secondary={`#${props.channelname}`} />
        </ListItem>
    );
};

export const ChannelList = () => {
    const channels = useSelector((state: IState) => state.channels);
    const selectedGuild = useSelector((state: IState) => state.selectedGuild);
    const [actionsExpanded, setActionsExpanded] = useState<boolean>(false);
    const dispatch = useDispatch();
    const classes = useChannelListStyle();

    const leaveGuild = () => {
        harmonySocket.leaveGuild(selectedGuild);
    };

    const setSelectedChannel = (value: string) => {
        dispatch({ type: Actions.SET_SELECTED_CHANNEL, payload: value });
    };

    return (
        <div>
            <List style={{ padding: 0 }}>
                {selectedGuild !== '' ? (
                    <>
                        <ListItem button onClick={() => setActionsExpanded(!actionsExpanded)}>
                            <ListItemIcon>
                                <SettingsIcon />
                            </ListItemIcon>
                            <ListItemText primary='Guild Options' />
                            {actionsExpanded ? <ExpandLess /> : <ExpandMore />}
                        </ListItem>
                        <Collapse in={actionsExpanded} timeout='auto' unmountOnExit>
                            <List component='div' disablePadding>
                                <ListItem button className={classes.nested} onClick={leaveGuild}>
                                    <ListItemIcon>
                                        <LeaveIcon />
                                    </ListItemIcon>
                                    <ListItemText primary='Leave Guild' />
                                </ListItem>
                            </List>
                        </Collapse>
                    </>
                ) : (
                    undefined
                )}
                {channels
                    ? Object.keys(channels).map((key) => {
                          return <Channel key={key} channelid={key} channelname={channels[key]} setSelectedChannel={setSelectedChannel} />;
                      })
                    : undefined}
            </List>
        </div>
    );
};
