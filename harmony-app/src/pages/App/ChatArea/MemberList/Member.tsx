import React, { useEffect } from 'react';
import { ListItem, ListItemAvatar, Avatar, ListItemText } from '@material-ui/core';
import { useSelector } from 'react-redux';

import { IState } from '../../../../types/redux';
import { harmonySocket } from '../../../../Root';

interface Props {
	userid: string;
}

export const Member = (props: Props) => {
	const { users } = useSelector((state: IState) => state);

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
