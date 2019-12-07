import React, { useState, useRef } from 'react';
import { List, ListItem, ListItemText, ListItemIcon, Collapse, Tooltip, Input } from '@material-ui/core';
import SettingsIcon from '@material-ui/icons/Settings';
import ExpandMore from '@material-ui/icons/ExpandMore';
import ExpandLess from '@material-ui/icons/ExpandLess';
import LeaveIcon from '@material-ui/icons/ExitToApp';
import { ContextMenu, ContextMenuTrigger } from 'react-contextmenu';
import { useSelector, useDispatch } from 'react-redux';
import { IState } from '../../../../types/redux';
import { useChannelListStyle } from './ChannelListStyle';
import { harmonySocket } from '../../../Root';
import { ToggleGuildSettingsDialog, SetSelectedChannel } from '../../../../redux/Dispatches';

interface IChannelProps {
    channelid: string;
    channelname: string;
    setSelectedChannel: (value: string) => void;
}

const Channel = (props: IChannelProps) => {
    //const guildList = useSelector((state: IState) => state.guildList);
    const selectedGuild = useSelector((state: IState) => state.selectedGuild);
    const selectedChannel = useSelector((state: IState) => state.selectedChannel);
    const classes = useChannelListStyle();

    const handleDelete = () => {
        harmonySocket.sendDeleteChannel(selectedGuild, props.channelid);
    };

    return (
        <>
            <ContextMenuTrigger id={props.channelid}>
                <ListItem button key={props.channelid} className={props.channelid === selectedChannel ? classes.selectedChannel : undefined} onClick={() => props.setSelectedChannel(props.channelid)}>
                    <ListItemText secondary={`#${props.channelname}`} />
                </ListItem>
            </ContextMenuTrigger>
            <ContextMenu id={props.channelid}>
                <List>
                    <ListItem button onClick={handleDelete}>
                        <ListItemText primary='Delete Channel' />
                    </ListItem>
                </List>
            </ContextMenu>
        </>
    );
};

export const ChannelList = () => {
    const channels = useSelector((state: IState) => state.channels);
    const selectedGuild = useSelector((state: IState) => state.selectedGuild);
    const guildsList = useSelector((state: IState) => state.guildList);
    const [actionsExpanded, setActionsExpanded] = useState<boolean>(false);
    const [addingChannel, setAddingChannel] = useState<boolean>(false);
    const addChannelInput = useRef<HTMLInputElement | null>(null);
    const dispatch = useDispatch();
    const classes = useChannelListStyle();

    const leaveGuild = () => {
        harmonySocket.leaveGuild(selectedGuild);
    };

    const setSelectedChannel = (value: string) => {
        dispatch(SetSelectedChannel(value));
    };

    const toggleGuildSettings = () => {
        harmonySocket.sendGetInvites(selectedGuild);
        dispatch(ToggleGuildSettingsDialog());
    };

    const addChannelButtonClicked = () => {
        setAddingChannel(true);
    };

    const handleChannelNameFinish = (ev: React.KeyboardEvent<HTMLInputElement>) => {
        if (ev.key === 'Enter' && addChannelInput.current) {
            harmonySocket.sendAddChannel(selectedGuild, addChannelInput.current.value);
            setAddingChannel(false);
        }
    };

    return (
        <div>
            <List style={{ padding: 0 }}>
                {selectedGuild ? (
                    <>
                        <ListItem button onClick={() => setActionsExpanded(!actionsExpanded)}>
                            <ListItemText primary='Guild Options' />
                            {actionsExpanded ? <ExpandLess /> : <ExpandMore />}
                        </ListItem>
                        <Collapse in={actionsExpanded} timeout='auto' unmountOnExit>
                            <List component='div' disablePadding>
                                {guildsList[selectedGuild] && guildsList[selectedGuild].owner ? (
                                    <>
                                        <ListItem button className={classes.nested} onClick={toggleGuildSettings}>
                                            <ListItemIcon>
                                                <SettingsIcon />
                                            </ListItemIcon>
                                            <ListItemText primary='Guild Settings' />
                                        </ListItem>
                                    </>
                                ) : (
                                    undefined
                                )}
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
                <div className={classes.newChannelInput}>
                    {addingChannel ? (
                        <Input fullWidth autoFocus onKeyPress={handleChannelNameFinish} onBlur={() => setAddingChannel(false)} placeholder={'Channel Name'} inputRef={addChannelInput} />
                    ) : (
                        undefined
                    )}
                </div>
                {selectedGuild && guildsList[selectedGuild] && guildsList[selectedGuild].owner ? (
                    <Tooltip title='Create Channel'>
                        <ListItem button onClick={addChannelButtonClicked}>
                            <ListItemText style={{ textAlign: 'center' }} primary='+' />
                        </ListItem>
                    </Tooltip>
                ) : (
                    undefined
                )}
            </List>
        </div>
    );
};
