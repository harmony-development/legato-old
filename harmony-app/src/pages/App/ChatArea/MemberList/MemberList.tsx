import React from 'react';
import { List } from '@material-ui/core';
import { useSelector } from 'react-redux';

import { RootState } from '../../../../redux/store';

import { Member } from './Member';

export const MemberList = () => {
	const { guildMembers, currentGuild } = useSelector((state: RootState) => state.app);

	return (
		<List>
			{currentGuild && guildMembers[currentGuild]
				? guildMembers[currentGuild].map(userid => {
						return <Member key={userid} userid={userid} />;
				  })
				: undefined}
		</List>
	);
};
