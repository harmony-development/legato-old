import React, { useEffect } from 'react';
import { ListItem, ListItemAvatar, Avatar, ListItemText } from '@material-ui/core';
import { useSelector } from 'react-redux';

import { harmonySocket } from '../../../../Root';
import { RootState } from '../../../../redux/store';

interface Props {
	userid: string;
}

export const Member = (props: Props) => {
	const { users } = useSelector((state: RootState) => state.app);

	useEffect(() => {
		if (!users[props.userid]) {
			harmonySocket.sendGetUser(props.userid);
		}
	}, []);

	return (
		<ListItem>
			<ListItemAvatar>
				<Avatar alt={props.userid} src={props.userid && users[props.userid] ? users[props.userid].avatar : undefined} />
			</ListItemAvatar>
			<ListItemText
				primary={<>{props.userid && users[props.userid] ? users[props.userid].username : props.userid}</>}
			/>
		</ListItem>
	);
};
