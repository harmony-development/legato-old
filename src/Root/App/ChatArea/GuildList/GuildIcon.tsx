import React from 'react';
import { ButtonBase, Tooltip } from '@material-ui/core';
import { useDispatch } from 'react-redux';
import { Actions } from '../../../../types/redux';
import { useGuildListStyle } from './GuildListStyle';

interface IProps {
    guildid: string;
    guildname: string;
    picture: string;
    selected: boolean;
}

export const GuildIcon = (props: IProps) => {
    const classes = useGuildListStyle();
    const dispatch = useDispatch();

    const onClick = () => {
        dispatch({
            type: Actions.SET_SELECTED_GUILD,
            payload: props.guildid
        });
    };

    return (
        <>
            <ButtonBase className={`${classes.guildiconroot} ${props.selected ? classes.selectedguildicon : undefined}`} key={props.guildid} onClick={onClick}>
                <Tooltip title={props.guildname} placement='right'>
                    <img className={classes.guildicon} alt='' src={props.picture} draggable={false} />
                </Tooltip>
            </ButtonBase>
        </>
    );
};
