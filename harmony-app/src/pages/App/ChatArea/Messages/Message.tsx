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
	Tooltip,
} from '@material-ui/core';
import { MoreVert, PlayArrow } from '@material-ui/icons';
import { useSelector } from 'react-redux';

import { harmonySocket } from '../../../../Root';
import { IState } from '../../../../types/redux';

interface IProps {
	userid: string;
	messageid: string;
	username: string;
	createdat: number;
	message: string;
	avatar?: string;
	attachment?: string;
}

const UtcEpochToLocalDate = (time: number) => {
	const returnDate = new Date(0);
	returnDate.setUTCSeconds(time);
	return ` - ${returnDate.toDateString()} at ${returnDate.toLocaleTimeString()}`;
};

export const Message = (props: IProps) => {
	const [dropdownBtnVisible, setDropdownBtnVisible] = useState(false);
	const [output, setOutput] = useState('');
	const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
	const { currentGuild, currentChannel, self } = useSelector((state: IState) => state);

	const handleDropdownBtnClick = (event: React.MouseEvent<HTMLButtonElement>) => {
		setAnchorEl(event.currentTarget);
	};

	const handleClose = () => {
		setAnchorEl(null);
	};

	const handleDelete = () => {
		if (currentChannel && currentGuild) {
			harmonySocket.sendDeleteMessage(currentGuild, currentChannel, props.messageid);
		}
	};

	return (
		<>
			<ListItem
				alignItems="flex-start"
				onMouseOver={() => setDropdownBtnVisible(true)}
				onMouseLeave={() => setDropdownBtnVisible(false)}
			>
				<ListItemAvatar>
					<Avatar alt={props.userid} src={props.avatar ? props.avatar : undefined} />
				</ListItemAvatar>
				<ListItemText
					primary={
						<Typography>
							{props.username || props.userid}{' '}
							<Typography component="span" variant="body1" color="textSecondary">
								{UtcEpochToLocalDate(props.createdat)}
							</Typography>
						</Typography>
					}
					disableTypography
					secondary={
						<>
							{props.message.startsWith('```') && props.message.endsWith('```') ? (
								<>
									<div style={{ backgroundColor: 'rgba(0, 0, 0, 0.2)', paddingLeft: '8px' }}>
										<Typography style={{ fontFamily: 'monospace', whiteSpace: 'break-spaces' }} display="inline">
											{props.message.substring(3, props.message.length - 3)}
										</Typography>
									</div>
									{localStorage.getItem('developerCodeExecution') === 'true' ? (
										<>
											<Tooltip title="Run Code Clientside">
												<IconButton
													size="small"
													onClick={() => {
														try {
															setOutput(eval(props.message.substring(3, props.message.length - 3)));
														} catch (e) {
															if (e instanceof Error) {
																setOutput(`ERROR : ${e.message}`);
															}
														}
													}}
												>
													<PlayArrow />
												</IconButton>
											</Tooltip>
											{output ? (
												<div style={{ backgroundColor: 'rgba(0, 0, 0, 0.1)', paddingLeft: '8px' }}>
													<Typography style={{ fontFamily: 'monospace', whiteSpace: 'break-spaces' }} display="inline">
														{'Output Console'}
														<br />
														{output}
													</Typography>
												</div>
											) : (
												undefined
											)}
										</>
									) : (
										undefined
									)}
								</>
							) : (
								<Typography>{props.message}</Typography>
							)}
							{props.attachment ? (
								<div style={{ display: 'flex', width: '100%', flex: '0 0 100%' }}>
									<img
										src={`http://${process.env.REACT_APP_HARMONY_SERVER_HOST}/filestore/${props.attachment}`}
										style={{ maxHeight: '300px', maxWidth: '500px' }}
									/>
								</div>
							) : (
								undefined
							)}
						</>
					}
				/>

				{dropdownBtnVisible && self.userid === props.userid ? (
					<>
						<IconButton edge="end" size="small" aria-label="message-options" onClick={handleDropdownBtnClick}>
							<MoreVert />
						</IconButton>
						<Menu open={Boolean(anchorEl)} onClose={handleClose} anchorEl={anchorEl}>
							<MenuItem onClick={handleClose}>Edit</MenuItem>
							<MenuItem onClick={handleDelete}>Delete</MenuItem>
						</Menu>
					</>
				) : (
					undefined
				)}
			</ListItem>
		</>
	);
};
