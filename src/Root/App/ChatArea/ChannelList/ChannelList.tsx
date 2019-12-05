import React from 'react';
import { List, ListItem, ListItemText, Typography } from '@material-ui/core';

export const ChannelList = () => {
    return (
        <div>
            <List style={{ padding: 0 }}>
                <ListItem button>
                    <ListItemText secondary={'#general'} />
                </ListItem>
                <ListItem button>
                    <ListItemText secondary={'#media'} />
                </ListItem>
                <ListItem button>
                    <ListItemText secondary={'#bruh'} />
                </ListItem>
            </List>
        </div>
    );
};
