import React, { useEffect, useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Paper } from '@material-ui/core';

import { IState } from '../../../types/redux';
import { AppDispatch } from '../../../redux/store';
import { FocusChatInput } from '../../../redux/AppReducer';
import { DisconnectedStatus } from '../DisconnectedStatus';

import { useChatAreaStyles } from './ChatAreaStyle';
import { Messages } from './Messages/Messages';
import { Input } from './Input/Input';
import { GuildList } from './GuildList/GuildList';
import { ChannelList } from './ChannelList/ChannelList';
import { MemberList } from './MemberList/MemberList';

export const ChatArea = () => {
	const classes = useChatAreaStyles();
	const dispatch = useDispatch<AppDispatch>();
	const [messages, connected] = useSelector((state: IState) => [state.messages, state.connected]);
	const messagesRef = useRef<HTMLDivElement | null>(null);

	useEffect(() => {
		if (messagesRef.current) {
			messagesRef.current.scrollTop = messagesRef.current.scrollHeight;
		}
	}, [messages]);

	const onKeyDown = (ev: React.KeyboardEvent<HTMLDivElement>) => {
		if (ev.key !== 'Tab') {
			dispatch(FocusChatInput());
		}
	};

	return (
		<div className={classes.root}>
			<Paper elevation={2} className={classes.guildlist} square>
				<div>
					<GuildList />
				</div>
			</Paper>
			<Paper elevation={2} className={classes.channellist} square>
				<ChannelList />
			</Paper>
			<div className={classes.chatArea}>
				<div className={classes.messages} ref={messagesRef} onKeyDown={onKeyDown} tabIndex={-1}>
					{!connected ? <DisconnectedStatus /> : undefined}
					<Messages />
				</div>
				<div className={classes.input}>
					<Input />
				</div>
			</div>
			<Paper elevation={2} className={classes.memberlist} square>
				<MemberList />
			</Paper>
		</div>
	);
};
