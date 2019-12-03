import React from 'react';
import { ButtonBase, Tooltip } from '@material-ui/core';
import { useGuildListStyle } from './GuildListStyle';
import { useSelector } from 'react-redux';
import { IState } from '../../../../types/redux';

export const GuildList = () => {
    const classes = useGuildListStyle();
    const guildList = useSelector((state: IState) => state.guildList);

    return (
        <div className={classes.guildlist}>
            {guildList.map((guild) => {
                return (
                    <ButtonBase className={classes.guildiconroot} key={guild.guildid}>
                        <Tooltip title={guild.guildname} placement='right'>
                            <img className={classes.guildicon} alt='' src={guild.picture} draggable={false} />
                        </Tooltip>
                    </ButtonBase>
                );
            })}
        </div>
    );
};
