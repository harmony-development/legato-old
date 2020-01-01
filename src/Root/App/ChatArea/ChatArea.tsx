import React, { useEffect, useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { IState } from '../../../types/redux';
import { AppDispatch } from '../../../redux/store';
import { FocusChatInput } from '../../../redux/AppReducer';

import { useChatAreaStyles } from './ChatAreaStyle';
import { Messages } from './Messages/Messages';
import { Input } from './Input/Input';
import { GuildList } from './GuildList/GuildList';
import { ChannelList } from './ChannelList/ChannelList';

export const ChatArea = () => {
	const classes = useChatAreaStyles();
	const dispatch = useDispatch<AppDispatch>();
	const [messages] = useSelector((state: IState) => [state.messages]);
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
			<div className={classes.guildlist}>
				<GuildList />
			</div>
			<div className={classes.channellist}>
				<ChannelList />
			</div>
			<div className={classes.chatArea}>
				<div className={classes.messages} ref={messagesRef} onKeyDown={onKeyDown} tabIndex={-1}>
					<Messages />
				</div>
				<div className={classes.input}>
					<Input />
				</div>
			</div>
		</div>
	);
};
