import React, { useState } from 'react';
import {
	ListItem,
	ListItemAvatar,
	Avatar,
	ListItemText,
	Typography,
	IconButton,
	Menu,
	MenuItem,
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
	const [dropdownBtnVisible, setDropdownBtnVisible] = useState(false);
	const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

	const handleDropdownBtnClick = (event: React.MouseEvent<HTMLButtonElement>) => {
		setAnchorEl(event.currentTarget);
	};

	const handleClose = () => {
		setAnchorEl(null);
	};

	return (
		<ListItem
			alignItems="flex-start"
			onMouseOver={() => setDropdownBtnVisible(true)}
			onMouseLeave={() => setDropdownBtnVisible(false)}
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

			{dropdownBtnVisible ? (
				<>
					<IconButton edge="end" size="small" aria-label="message-options" onClick={handleDropdownBtnClick}>
						<MoreVert />
					</IconButton>
					<Menu open={Boolean(anchorEl)} onClose={handleClose} anchorEl={anchorEl}>
						<MenuItem onClick={handleClose}>Edit</MenuItem>
						<MenuItem onClick={handleClose}>Delete</MenuItem>
					</Menu>
				</>
			) : (
				undefined
			)}
		</ListItem>
	);
};
