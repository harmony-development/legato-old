import React, { useEffect, useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Paper } from '@material-ui/core';

import { AppDispatch, RootState } from '../../../redux/store';
import { FocusChatInput } from '../../../redux/AppReducer';
import { DisconnectedStatus } from '../DisconnectedStatus';
import { harmonySocket } from '../../../Root';

import { useChatAreaStyles } from './ChatAreaStyle';
import { Messages } from './Messages/Messages';
import { Input } from './Input/Input';
import { GuildList } from './GuildList/GuildList';
import { ChannelList } from './ChannelList/ChannelList';
import { MemberList } from './MemberList/MemberList';

let preScrollHeight = 0;
let preScrollTop = 0;
let scrollTrigger = false;

export const ChatArea = () => {
	const classes = useChatAreaStyles();
	const dispatch = useDispatch<AppDispatch>();
	const { messages, connected, currentGuild, currentChannel } = useSelector((state: RootState) => state.app);
	const messagesRef = useRef<HTMLDivElement | null>(null);

	useEffect(() => {
		if (messagesRef.current) {
			if (!scrollTrigger) {
				messagesRef.current.scrollTop = messagesRef.current.scrollHeight;
			} else {
				messagesRef.current.scrollTop = messagesRef.current.scrollHeight - preScrollHeight + preScrollTop;
				scrollTrigger = false;
			}
		}
	}, [messages]);

	const onKeyDown = (ev: React.KeyboardEvent<HTMLDivElement>) => {
		if (ev.key !== 'Tab') {
			dispatch(FocusChatInput());
		}
	};

	const onScroll = (event: React.UIEvent<HTMLDivElement>) => {
		if (event.currentTarget.scrollTop === 0 && currentGuild && currentChannel) {
			scrollTrigger = true;
			preScrollHeight = event.currentTarget.scrollHeight;
			preScrollTop = event.currentTarget.scrollTop;
			harmonySocket.getOldMessages(currentGuild, currentChannel, messages[0].messageid);
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
				<div className={classes.messages} ref={messagesRef} onKeyDown={onKeyDown} tabIndex={-1} onScroll={onScroll}>
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
