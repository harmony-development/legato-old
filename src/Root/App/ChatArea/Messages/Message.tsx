import React, { useState, useEffect } from 'react';
import {
	ListItem,
	ListItemAvatar,
	Avatar,
	ListItemText,
	Typography,
	ListItemSecondaryAction,
	IconButton,
} from '@material-ui/core';
import { MoreVert } from '@material-ui/icons';

interface IProps {
	userid: string;
	username: string;
	createdat: number;
	message: string;
	avatar?: string;
}

const UtcEpochToLocalDate = (time: number) => {
	const returnDate = new Date(0);
	returnDate.setUTCSeconds(time);
	return ` - ${returnDate.toDateString()} at ${returnDate.toLocaleTimeString()}`;
};

export const Message = (props: IProps) => {
	const [dropdownVisible, setDropdownVisible] = useState(false);

	return (
		<ListItem
			alignItems="flex-start"
			onMouseOver={() => setDropdownVisible(true)}
			onMouseLeave={() => setDropdownVisible(false)}
		>
			<ListItemAvatar>
				<Avatar alt={props.userid} src={props.avatar ? `${props.avatar}` : undefined} />
			</ListItemAvatar>
			<ListItemText
				primary={
					<>
						{props.username || props.userid}
						<Typography component="span" variant="body1" color="textSecondary">
							{UtcEpochToLocalDate(props.createdat)}
						</Typography>
					</>
				}
				secondary={props.message}
			/>

			{dropdownVisible ? (
				<IconButton edge="end" size="small" aria-label="message-options">
					<MoreVert />
				</IconButton>
			) : (
				undefined
			)}
		</ListItem>
	);
};
