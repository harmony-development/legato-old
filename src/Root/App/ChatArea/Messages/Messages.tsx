import React, { useEffect, useRef } from 'react';
import { List } from '@material-ui/core';
import { useSelector } from 'react-redux';

import { IState } from '../../../../types/redux';
import { harmonySocket } from '../../../Root';

import { Message } from './Message';

export const Messages = () => {
	const [messages, users] = useSelector((state: IState) => [state.messages, state.users]);
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
			if (!users[val] || !users[val].avatar || !users[val].username) {
				harmonySocket.sendGetUser(val);
			}
		});
	}, [messages, users]);

	return (
		<List innerRef={messageList}>
			{messages
				? messages.map(val => {
						return (
							<Message
								key={val.messageid}
								userid={val.userid}
								messageid={val.messageid}
								username={users[val.userid] ? users[val.userid].username : ''}
								createdat={val.createdat}
								avatar={users[val.userid] ? users[val.userid].avatar : undefined}
								message={val.message}
							/>
						);
				  })
				: undefined}
		</List>
	);
};
