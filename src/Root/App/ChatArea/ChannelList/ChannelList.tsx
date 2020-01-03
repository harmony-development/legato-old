import React, { useState, useRef } from 'react';
import { List, ListItem, ListItemText, ListItemIcon, Collapse, Tooltip, Input } from '@material-ui/core';
import SettingsIcon from '@material-ui/icons/Settings';
import ExpandMore from '@material-ui/icons/ExpandMore';
import ExpandLess from '@material-ui/icons/ExpandLess';
import LeaveIcon from '@material-ui/icons/ExitToApp';
import { ContextMenu, ContextMenuTrigger } from 'react-contextmenu';
import { useDispatch, useSelector } from 'react-redux';

import { IState } from '../../../../types/redux';
import { harmonySocket } from '../../../Root';
import { AppDispatch } from '../../../../redux/store';
import { SetCurrentChannel, ToggleGuildSettingsDialog } from '../../../../redux/AppReducer';

import { useChannelListStyle } from './ChannelListStyle';

interface IChannelProps {
	channelid: string;
	channelname: string;
	setSelectedChannel: (value: string) => void;
}

const Channel = (props: IChannelProps) => {
	const [currentGuild, selectedChannel, guildsList] = useSelector((state: IState) => [
		state.currentGuild,
		state.currentChannel,
		state.guildList,
	]);
	const classes = useChannelListStyle();

	const handleDelete = () => {
		if (currentGuild) {
			harmonySocket.sendDeleteChannel(currentGuild, props.channelid);
		}
	};

	return (
		<>
			<ContextMenuTrigger id={props.channelid}>
				<ListItem
					button
					key={props.channelid}
					className={props.channelid === selectedChannel ? classes.selectedChannel : undefined}
					onClick={() => props.setSelectedChannel(props.channelid)}
				>
					<ListItemText secondary={`#${props.channelname}`} />
				</ListItem>
			</ContextMenuTrigger>
			{currentGuild && guildsList[currentGuild] && guildsList[currentGuild].owner ? (
				<ContextMenu id={props.channelid}>
					<List>
						<ListItem button onClick={handleDelete}>
							<ListItemText primary="Delete Channel" />
						</ListItem>
					</List>
				</ContextMenu>
			) : (
				undefined
			)}
		</>
	);
};

export const ChannelList = () => {
	const dispatch = useDispatch<AppDispatch>();
	const [channels, currentGuild, guildsList] = useSelector((state: IState) => [
		state.channels,
		state.currentGuild,
		state.guildList,
	]);
	const [actionsExpanded, setActionsExpanded] = useState<boolean>(false);
	const [addingChannel, setAddingChannel] = useState<boolean>(false);
	const addChannelInput = useRef<HTMLInputElement | null>(null);
	const classes = useChannelListStyle();

	const leaveGuild = () => {
		if (currentGuild) {
			harmonySocket.leaveGuild(currentGuild);
		}
	};

	const setSelectedChannel = (value: string) => {
		dispatch(SetCurrentChannel(value));
	};

	const toggleGuildSettings = () => {
		if (currentGuild) {
			harmonySocket.sendGetInvites(currentGuild);
			dispatch(ToggleGuildSettingsDialog());
		}
	};

	const addChannelButtonClicked = () => {
		setAddingChannel(true);
	};

	const handleChannelNameFinish = (ev: React.KeyboardEvent<HTMLInputElement>) => {
		if (ev.key === 'Enter' && addChannelInput.current && currentGuild) {
			harmonySocket.sendAddChannel(currentGuild, addChannelInput.current.value);
			setAddingChannel(false);
		}
	};

	return (
		<div>
			<List style={{ padding: 0 }}>
				{currentGuild ? (
					<>
						<ListItem button onClick={() => setActionsExpanded(!actionsExpanded)}>
							<ListItemText primary="Guild Options" />
							{actionsExpanded ? <ExpandLess /> : <ExpandMore />}
						</ListItem>
						<Collapse in={actionsExpanded} timeout="auto" unmountOnExit>
							<List component="div" disablePadding>
								{guildsList[currentGuild] && guildsList[currentGuild].owner ? (
									<>
										<ListItem button className={classes.nested} onClick={toggleGuildSettings}>
											<ListItemIcon>
												<SettingsIcon />
											</ListItemIcon>
											<ListItemText primary="Guild Settings" />
										</ListItem>
									</>
								) : (
									undefined
								)}
								<ListItem button className={classes.nested} onClick={leaveGuild}>
									<ListItemIcon>
										<LeaveIcon />
									</ListItemIcon>
									<ListItemText primary="Leave Guild" />
								</ListItem>
							</List>
						</Collapse>
					</>
				) : (
					undefined
				)}
				{channels
					? Object.keys(channels).map(key => {
							return (
								<Channel
									key={key}
									channelid={key}
									channelname={channels[key]}
									setSelectedChannel={setSelectedChannel}
								/>
							);
					  })
					: undefined}
				<div className={classes.newChannelInput}>
					{addingChannel ? (
						<Input
							fullWidth
							autoFocus
							onKeyPress={handleChannelNameFinish}
							onBlur={() => setAddingChannel(false)}
							placeholder={'Channel Name'}
							inputRef={addChannelInput}
						/>
					) : (
						undefined
					)}
				</div>
				{currentGuild && guildsList[currentGuild] && guildsList[currentGuild].owner ? (
					<Tooltip title="Create Channel">
						<ListItem button onClick={addChannelButtonClicked}>
							<ListItemText style={{ textAlign: 'center' }} primary="+" />
						</ListItem>
					</Tooltip>
				) : (
					undefined
				)}
			</List>
		</div>
	);
};
