import React, { useRef, useEffect, useState } from 'react';
import { TextField, Button, Box } from '@material-ui/core';
import { useSelector } from 'react-redux';
import { Add as AddIcon } from '@material-ui/icons';

import { IState } from '../../../../types/redux';
import { harmonySocket } from '../../../../Root';

import { useInputStyles } from './InputStyle';
import { FileCard } from './FileCard/FileCard';

export const Input = () => {
	const [connected, inputStyle, guildID, channelID, focus] = useSelector((state: IState) => [
		state.connected,
		state.theme.inputStyle,
		state.currentGuild,
		state.currentChannel,
		state.chatInputFocus,
	]);
	const inputField = useRef<HTMLInputElement | undefined>();
	const fileUploadRef = useRef<HTMLInputElement | null>(null);
	const classes = useInputStyles();
	const [fileQueue, setFileQueue] = useState<
		{
			file: File | undefined;
			preview: string;
		}[]
	>([]);

	const removeFromQueue = (index: number) =>
		setFileQueue([...fileQueue.slice(0, index), ...fileQueue.slice(index + 1)]);

	const onKeyPress = (e: React.KeyboardEvent) => {
		if (e.key === 'Enter' && !e.shiftKey) {
			e.preventDefault();
			// does the input field exist and is it not blank
			if (inputField.current && !/^\s*$/.test(inputField.current.value) && channelID && guildID) {
				if (fileQueue.length == 0) {
					harmonySocket.sendMessage(guildID, channelID, inputField.current.value);
				} else if (fileQueue[0].file instanceof File) {
					harmonySocket.sendMessageRest(guildID, channelID, inputField.current.value, fileQueue[0].file);
				}
				inputField.current.value = '';
				setFileQueue([]);
			}
		}
	};

	const onImageSelected = (event: React.ChangeEvent<HTMLInputElement>) => {
		const file = event.currentTarget.files?.[0];
		if (file) {
			if (file.type.startsWith('image/') && file.size < 33554432) {
				const fileReader = new FileReader();
				fileReader.readAsDataURL(file);
				fileReader.addEventListener('load', () => {
					if (typeof fileReader.result === 'string') {
						setFileQueue([
							...fileQueue,
							{
								file,
								preview: fileReader.result,
							},
						]);
					}
				});
			} else {
				setFileQueue([
					...fileQueue,
					{
						file,
						preview: '',
					},
				]);
			}
		}
	};

	useEffect(() => {
		if (inputField.current) {
			inputField.current.focus();
		}
	}, [focus]);

	return (
		<>
			<input type="file" id="file" ref={fileUploadRef} style={{ display: 'none' }} onChange={onImageSelected} />
			<div className={classes.fileQueue}>
				{fileQueue.map(entry => {
					return <FileCard key={entry.file?.name} image={entry.preview} removeFromQueue={removeFromQueue} index={0} />;
				})}
			</div>
			<div className={classes.inputRoot}>
				<Button onClick={() => fileUploadRef.current?.click()}>
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
		</>
	);
};
