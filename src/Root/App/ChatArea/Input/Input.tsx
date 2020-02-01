import React, { useRef, useEffect } from 'react';
import { TextField, Button } from '@material-ui/core';
import { useSelector } from 'react-redux';
import { Add as AddIcon } from '@material-ui/icons';

import { IState } from '../../../../types/redux';
import { harmonySocket } from '../../../Root';

import { useInputStyles } from './InputStyle';

export const Input = () => {
	const [connected, inputStyle, guildID, channelID, focus] = useSelector((state: IState) => [
		state.connected,
		state.theme.inputStyle,
		state.currentGuild,
		state.currentChannel,
		state.chatInputFocus,
	]);
	const inputField = useRef<HTMLInputElement | undefined>();
	const classes = useInputStyles();

	const onKeyPress = (e: React.KeyboardEvent) => {
		if (e.key === 'Enter' && !e.shiftKey) {
			e.preventDefault();
			// does the input field exist and is it not blank
			if (inputField.current && !/^\s*$/.test(inputField.current.value) && channelID && guildID) {
				harmonySocket.sendMessage(guildID, channelID, inputField.current.value);
				inputField.current.value = '';
			}
		}
	};

	useEffect(() => {
		if (inputField.current) {
			inputField.current.focus();
		}
	}, [focus]);

	return (
		<div className={classes.inputRoot}>
			<Button>
				<AddIcon />
			</Button>
			<TextField
				label={connected ? 'Message' : 'Currently Offline'}
				variant={inputStyle as any}
				fullWidth
				multiline
				rowsMax={3}
				rows={3}
				onKeyPress={onKeyPress}
				inputRef={inputField}
				color="secondary"
			/>
		</div>
	);
};
