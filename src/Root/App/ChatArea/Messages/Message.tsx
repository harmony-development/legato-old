import React from 'react';
import { ListItem, ListItemAvatar, Avatar, ListItemText, Typography } from '@material-ui/core';

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
	return (
		<ListItem alignItems='flex-start'>
			<ListItemAvatar>
				<Avatar
					alt={props.userid}
					src={
						props.avatar
							? `http://${process.env.REACT_APP_HARMONY_SERVER_HOST}/filestore/${props.avatar}`
							: undefined
					}
				/>
			</ListItemAvatar>
			<ListItemText
				primary={
					<>
						{props.username || props.userid}
						<Typography component='span' variant='body1' color='textSecondary'>
							{UtcEpochToLocalDate(props.createdat)}
						</Typography>
					</>
				}
				secondary={props.message}
			/>
		</ListItem>
	);
};
