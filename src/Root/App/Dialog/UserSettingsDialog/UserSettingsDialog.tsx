import React, { useState } from 'react';
import { Dialog, AppBar, Toolbar, IconButton, Typography, DialogContent, TextField, Button } from '@material-ui/core';
import CloseIcon from '@material-ui/icons/Close';
import { useDispatch, useSelector } from 'react-redux';
import { toast } from 'react-toastify';
import axios from 'axios';

import { IState } from '../../../../types/redux';
import { AppDispatch } from '../../../../redux/store';
import { ToggleUserSettingsDialog } from '../../../../redux/AppReducer';
import { ImagePicker } from '../ImagePicker';
import { harmonySocket } from '../../../Root';

import { useUserSettingsStyle } from './UserSettingsStyle';

export const UserSettingsDialog = () => {
	const dispatch = useDispatch<AppDispatch>();
	const [open, globalUsername, globalAvatar, inputStyle] = useSelector((state: IState) => [
		state.userSettingsDialog,
		state.self.username || undefined,
		state.self.avatar || undefined,
		state.theme.inputStyle,
	]);
	const [username, setUsername] = useState<string>(globalUsername || '');
	const [avatar, setAvatar] = useState<string | undefined>(globalAvatar);
	const [avatarFile, setAvatarFile] = useState<File | null>(null);
	const classes = useUserSettingsStyle();

	const onSaveChanges = () => {
		if (avatarFile && username) {
			const avatarFileUpload = new FormData();
			avatarFileUpload.append('token', localStorage.getItem('token') || 'none');
			avatarFileUpload.append('file', avatarFile);
			axios
				.post(`http://${process.env.REACT_APP_HARMONY_SERVER_HOST}/api/rest/fileupload`, avatarFileUpload, {})
				.then(res => {
					if (res.data) {
						const uploadID = res.data;
						harmonySocket.sendAvatarUpdate(`http://${process.env.REACT_APP_HARMONY_SERVER_HOST}/filestore/${uploadID}`);
					}
				})
				.catch(err => {
					console.log(err);
					toast.error('Failed to update avatar');
				});
		}
	};

	return (
		<Dialog open={open} onClose={() => dispatch(ToggleUserSettingsDialog())} fullScreen>
			<AppBar style={{ position: 'relative' }}>
				<Toolbar>
					<IconButton edge="start" onClick={() => dispatch(ToggleUserSettingsDialog())}>
						<CloseIcon />
					</IconButton>
					<Typography variant="h6">User Settings</Typography>
				</Toolbar>
			</AppBar>
			<DialogContent>
				<div style={{ width: '33%' }}>
					<ImagePicker setImageFile={setAvatarFile} setImage={setAvatar} image={avatar} />
					<TextField
						label="Username"
						fullWidth
						variant={inputStyle as any}
						className={classes.menuEntry}
						value={username}
						onChange={(e: React.ChangeEvent<HTMLInputElement>) => setUsername(e.currentTarget.value)}
					/>
					<Button variant="contained" color="secondary" className={classes.menuEntry} onClick={onSaveChanges}>
						Save Changes
					</Button>
				</div>
			</DialogContent>
		</Dialog>
	);
};
