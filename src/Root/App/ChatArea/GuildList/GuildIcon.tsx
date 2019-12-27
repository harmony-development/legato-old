import React from 'react';
import { ButtonBase, Tooltip, List, ListItem, ListItemText } from '@material-ui/core';
import { useDispatch, useSelector } from 'react-redux';
import { Actions, IState } from '../../../../types/redux';
import { useGuildListStyle } from './GuildListStyle';
import { ContextMenuTrigger, ContextMenu, MenuItem } from 'react-contextmenu';
import { harmonySocket } from '../../../Root';
import { ToggleGuildSettingsDialog } from '../../../../redux/Dispatches';

interface IProps {
    guildid: string;
    guildname: string;
    picture: string;
    selected: boolean;
}

export const GuildIcon = (props: IProps) => {
    const [guildsList] = useSelector((state: IState) => [state.guildList]);
    const classes = useGuildListStyle();
    const dispatch = useDispatch();

    const onClick = () => {
        dispatch({
            type: Actions.SET_CURRENT_GUILD,
            payload: props.guildid
        });
    };

    const handleLeave = () => {
        harmonySocket.leaveGuild(props.guildid);
    };

    return (
        <>
            <ContextMenuTrigger id={props.guildid}>
                <ButtonBase
                    className={`${classes.guildiconroot} ${props.selected ? classes.selectedguildicon : undefined}`}
                    key={props.guildid}
                    onClick={onClick}
                >
                    <Tooltip title={props.guildname} placement='right'>
                        <img className={classes.guildicon} alt='' src={props.picture} draggable={false} />
                    </Tooltip>
                </ButtonBase>
            </ContextMenuTrigger>
            <ContextMenu id={props.guildid}>
                <List>
                    <MenuItem>
                        <ListItem button onClick={handleLeave}>
                            <ListItemText primary='Leave Guild' />
                        </ListItem>
                    </MenuItem>
                    {guildsList && guildsList[props.guildid].owner ? (
                        <>
                            <MenuItem>
                                <ListItem button onClick={() => dispatch(ToggleGuildSettingsDialog())}>
                                    <ListItemText primary='Guild Settings' />
                                </ListItem>
                            </MenuItem>
                        </>
                    ) : (
                        undefined
                    )}
                </List>
            </ContextMenu>
        </>
    );
};
