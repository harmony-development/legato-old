import React from 'react';
import { List, ListItem, ListItemText } from '@material-ui/core';
import { useSelector, useDispatch } from 'react-redux';
import { IState, Actions } from '../../../../types/redux';
import { useChannelListStyle } from './ChannelListStyle';

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
    const dispatch = useDispatch();

    const setSelectedChannel = (value: string) => {
        dispatch({ type: Actions.SET_SELECTED_CHANNEL, payload: value });
    };

    return (
        <div>
            <List style={{ padding: 0 }}>
                {channels
                    ? Object.keys(channels).map((key) => {
                          return <Channel key={key} channelid={key} channelname={channels[key]} setSelectedChannel={setSelectedChannel} />;
                      })
                    : undefined}
            </List>
        </div>
    );
};
