import React, { useEffect, useRef } from 'react';
import { List } from '@material-ui/core';
import { useSelector } from 'react-redux';

import { IState } from '../../../../types/redux';
import { harmonySocket } from '../../../Root';

import { Message } from './Message';

export const Messages = () => {
	const [messages, selectedChannel, users] = useSelector((state: IState) => [
		state.messages,
		state.currentChannel,
		state.users,
	]);
	const messageList = useRef<HTMLUListElement | undefined>(undefined);

	useEffect(() => {
		if (messageList.current) {
			messageList.current.scrollTop = messageList.current.scrollHeight;
			messageList.current.scrollLeft = 0;
		}
	}, [messages]);

	useEffect(() => {
		const userIDs = [...new Set(messages.map(val => val.userid))];
		userIDs.forEach(val => {
			if (!users[val]) {
				console.log(val);
				harmonySocket.sendGetUser(val);
			}
		});
	}, [messages]);

	return (
		<List innerRef={messageList}>
			{messages
				? messages.map(val => {
						if (val.channel === selectedChannel) {
							return (
								<Message
									key={val.messageid}
									guild={val.guild}
									userid={val.userid}
									username={users[val.userid] ? users[val.userid].username : ''}
									createdat={val.createdat}
									avatar={users[val.userid] ? users[val.userid].avatar : undefined}
									message={val.message}
								/>
							);
						} else {
							return undefined;
						}
				  })
				: undefined}
		</List>
	);
};
