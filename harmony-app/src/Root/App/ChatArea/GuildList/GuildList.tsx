import React, { useEffect } from 'react';
import GroupAdd from '@material-ui/icons/GroupAdd';
import { useDispatch, useSelector } from 'react-redux';
import { ButtonBase, Tooltip } from '@material-ui/core';

import { IState } from '../../../../types/redux';
import { ToggleGuildDialog } from '../../../../redux/AppReducer';
import { AppDispatch } from '../../../../redux/store';

import { useGuildListStyle } from './GuildListStyle';
import { GuildIcon } from './GuildIcon';

export const GuildList = () => {
	const classes = useGuildListStyle();
	const dispatch = useDispatch<AppDispatch>();
	const [guildList, currentGuild] = useSelector((state: IState) => [state.guildList, state.currentGuild]);

	return (
		<div className={classes.guildlist}>
			{Object.keys(guildList).map(key => {
				return (
					<GuildIcon
						guildid={key}
						key={key}
						selected={currentGuild === key}
						guildname={guildList[key].guildname}
						picture={guildList[key].picture}
					/>
				);
			})}
			<ButtonBase className={classes.joinGuild} onClick={() => dispatch(ToggleGuildDialog())}>
				<Tooltip title={'Join Or Create Guild'} placement="right">
					<GroupAdd />
				</Tooltip>
			</ButtonBase>
		</div>
	);
};
