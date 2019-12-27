import React from 'react';
import GroupAdd from '@material-ui/icons/GroupAdd';
import { useGuildListStyle } from './GuildListStyle';
import { useSelector, useDispatch } from 'react-redux';
import { IState, Actions } from '../../../../types/redux';
import { GuildIcon } from './GuildIcon';
import { ButtonBase, Tooltip } from '@material-ui/core';

export const GuildList = () => {
    const classes = useGuildListStyle();
    const [guildList, selectedGuild] = useSelector((state: IState) => [state.guildList, state.currentGuild]);
    const dispatch = useDispatch();

    return (
        <div className={classes.guildlist}>
            {Object.keys(guildList).map((key) => {
                return (
                    <GuildIcon
                        guildid={key}
                        key={key}
                        selected={selectedGuild === key}
                        guildname={guildList[key].guildname}
                        picture={guildList[key].picture}
                    />
                );
            })}
            <ButtonBase
                className={classes.joinGuild}
                onClick={() => dispatch({ type: Actions.TOGGLE_JOIN_GUILD_DIALOG })}
            >
                <Tooltip title={'Join Or Create Guild'} placement='right'>
                    <GroupAdd />
                </Tooltip>
            </ButtonBase>
        </div>
    );
};
