import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import axios from 'axios';
import { Dialog, DialogContent, AppBar, Toolbar, IconButton, Typography, Button, TextField } from '@material-ui/core';
import CloseIcon from '@material-ui/icons/Close';
import { toast } from 'react-toastify';

import { IState } from '../../../../types/redux';
import { harmonySocket } from '../../../Root';
import { AppDispatch } from '../../../../redux/store';
import { ToggleGuildSettingsDialog } from '../../../../redux/AppReducer';
import { ImagePicker } from '../ImagePicker';
import { useConfirmationContext } from '../../ConfirmationContext';

import { useGuildSettingsStyle } from './GuildSettingsStyle';
import { GuildInvites } from './GuildInvites';

export const GuildSettings = () => {
	const dispatch = useDispatch<AppDispatch>();
	const confirm = useConfirmationContext();
	const [open, currentGuild, inputStyle, guilds] = useSelector((state: IState) => [
		state.guildSettingsDialog,
		state.currentGuild,
		state.theme.inputStyle,
		state.guildList,
	]);
	const [guildName, setGuildName] = useState<string | undefined>(
		currentGuild ? (guilds[currentGuild] ? guilds[currentGuild].guildname : undefined) : undefined
	);
	const [guildIconFile, setGuildIconFile] = useState<File | null>(null);
	const [guildIcon, setGuildIcon] = useState<string | undefined>(
		currentGuild ? (guilds[currentGuild] ? guilds[currentGuild].picture : undefined) : undefined
	);
	const classes = useGuildSettingsStyle();

	const onSaveChanges = () => {
		if (currentGuild && guilds[currentGuild]) {
			if (guildIcon !== guilds[currentGuild].picture && guildIconFile) {
				const guildIconUpload = new FormData();
				guildIconUpload.append('token', localStorage.getItem('token') || 'none');
				guildIconUpload.append('guild', currentGuild);
				guildIconUpload.append('file', guildIconFile);
				axios
					.post(`http://${process.env.REACT_APP_HARMONY_SERVER_HOST}/api/rest/updateguildpicture`, guildIconUpload, {})
					.catch(err => {
						console.log(err);
						toast.error('Failed to update guild icon');
					});
			}
			if (guilds[currentGuild].guildname !== guildName && guildName) {
				harmonySocket.sendGuildNameUpdate(currentGuild, guildName);
			}
		}
	};

	const handleDeleteGuild = () => {
		confirm({
			title: 'Are you sure you would like to delete this guild?',
			description: 'This cannot be undone!',
		}).then(() => {
			if (currentGuild) {
				harmonySocket.sendDeleteGuild(currentGuild);
			}
		});
	};

	useEffect(() => {
		if (currentGuild) {
			harmonySocket.sendGetInvites(currentGuild);
		}
	}, [currentGuild]);

	return (
		<Dialog open={open} onClose={() => dispatch(ToggleGuildSettingsDialog())} fullScreen>
			<AppBar style={{ position: 'relative' }}>
				<Toolbar>
					<IconButton edge="start" color="inherit" onClick={() => dispatch(ToggleGuildSettingsDialog())}>
						<CloseIcon />
					</IconButton>
					<Typography variant="h6">Guild Settings</Typography>
				</Toolbar>
			</AppBar>
			<DialogContent>
				<div style={{ width: '33%' }}>
					<ImagePicker
						setImageFile={setGuildIconFile}
						setImage={setGuildIcon}
						image={
							guildIcon?.startsWith('data:image')
								? guildIcon
								: `http://${process.env.REACT_APP_HARMONY_SERVER_HOST}/filestore/${guildIcon}`
						}
					/>
					<TextField
						label="Guild Name"
						fullWidth
						variant={inputStyle as any}
						className={classes.menuEntry}
						value={guildName}
						onChange={(e: React.ChangeEvent<HTMLInputElement>) => setGuildName(e.currentTarget.value)}
					/>
					<Button variant="contained" color="secondary" className={classes.menuEntry} onClick={onSaveChanges}>
						Save Changes
					</Button>
					<Typography variant="h4" className={classes.menuEntry}>
						Join Codes
					</Typography>
					<GuildInvites />
					<Button variant="outlined" color="secondary" className={classes.menuEntry} onClick={handleDeleteGuild}>
						Delete Guild
					</Button>
				</div>
			</DialogContent>
		</Dialog>
	);
};
