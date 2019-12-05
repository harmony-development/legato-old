import React from 'react';
import { List, ListItem, ListItemText } from '@material-ui/core';
import { useSelector } from 'react-redux';
import { IState } from '../../../../types/redux';

export const ChannelList = () => {
    const channels = useSelector((state: IState) => state.channels);

    return (
        <div>
            <List style={{ padding: 0 }}>
                {channels
                    ? Object.keys(channels).map((key) => {
                          return (
                              <ListItem button key={key}>
                                  <ListItemText secondary={`#${channels[key]}`} />
                              </ListItem>
                          );
                      })
                    : undefined}
            </List>
        </div>
    );
};
