import React from 'react';
import { useGuildListStyle } from './GuildListStyle';
import { useSelector } from 'react-redux';
import { IState } from '../../../../types/redux';
import { GuildIcon } from './GuildIcon';

export const GuildList = () => {
    const classes = useGuildListStyle();
    const guildList = useSelector((state: IState) => state.guildList);

    return (
        <div className={classes.guildlist}>
            {Object.keys(guildList).map((key) => {
                return <GuildIcon guildid={key} key={key} guildname={guildList[key].guildname} picture={guildList[key].picture} />;
            })}
        </div>
    );
};
