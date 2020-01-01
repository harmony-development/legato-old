import React, { useRef, useState } from 'react';
import { useSelector } from 'react-redux';
import axios from 'axios';
import {
	Dialog,
	DialogContent,
	AppBar,
	Toolbar,
	IconButton,
	Typography,
	Button,
	TextField,
	Avatar,
	ButtonBase,
	Table,
	TableHead,
	TableRow,
	TableCell,
	TableBody,
	Tooltip,
} from '@material-ui/core';
import copy from 'copy-to-clipboard';
import AddIcon from '@material-ui/icons/Add';
import CloseIcon from '@material-ui/icons/Close';
import DeleteIcon from '@material-ui/icons/Delete';
import ShareIcon from '@material-ui/icons/Share';
import { toast } from 'react-toastify';

import { IState } from '../../../../types/redux';
import { harmonySocket } from '../../../Root';
import { store } from '../../../../redux/store';
import { ToggleGuildSettingsDialog } from '../../../../redux/AppReducer';

import { useGuildSettingsStyle } from './GuildSettingsStyle';

export const GuildSettings = () => {
	const [open, currentGuild, inputStyle, guilds, invites] = useSelector((state: IState) => [
		state.guildSettingsDialog,
		state.currentGuild,
		state.theme.inputStyle,
		state.guildList,
		state.invites,
	]);
	const guildIconUpload = useRef<HTMLInputElement | null>(null);
	const [guildName, setGuildName] = useState<string | undefined>(
		currentGuild ? (guilds[currentGuild] ? guilds[currentGuild].guildname : undefined) : undefined
	);
	const [guildIconFile, setGuildIconFile] = useState<File | null>(null);
	const [guildIcon, setGuildIcon] = useState<string | undefined>(
		currentGuild ? (guilds[currentGuild] ? guilds[currentGuild].picture : undefined) : undefined
	);
	const classes = useGuildSettingsStyle();

	const deleteInviteLink = (invite: string) => {
		if (currentGuild) {
			harmonySocket.sendDeleteInvite(invite, currentGuild);
		}
	};

	const createInviteLink = () => {
		if (currentGuild) {
			harmonySocket.sendCreateInvite(currentGuild);
		}
	};

	const onSaveChanges = () => {
		if (currentGuild && guilds[currentGuild]) {
			if (guildIcon !== guilds[currentGuild].picture && guildIconFile) {
				const guildIconUpload = new FormData();
				guildIconUpload.append('file', guildIconFile);
				axios
					.post(`http://${process.env.REACT_APP_HARMONY_SERVER_HOST}/api/rest/fileupload`, guildIconUpload, {})
					.then(res => {
						if (res.data) {
							const uploadID = res.data;
							harmonySocket.sendGuildPictureUpdate(
								currentGuild,
								`http://${process.env.REACT_APP_HARMONY_SERVER_HOST}/filestore/${uploadID}`
							);
						}
					})
					.catch(() => {
						toast.error('Failed to update guild icon');
					});
			}
			if (guilds[currentGuild].guildname !== guildName && guildName) {
				harmonySocket.sendGuildNameUpdate(currentGuild, guildName);
			}
		}
	};

	const onGuildIconSelected = (event: React.ChangeEvent<HTMLInputElement>) => {
		if (event.currentTarget.files && event.currentTarget.files.length > 0) {
			const file = event.currentTarget.files[0];
			setGuildIconFile(file);
			if (file.type.startsWith('image/') && file.size < 33554432) {
				const fileReader = new FileReader();
				fileReader.readAsDataURL(file);
				fileReader.addEventListener('load', () => {
					if (typeof fileReader.result === 'string') {
						setGuildIcon(fileReader.result);
					}
				});
			}
		}
	};

	return (
		<Dialog open={open} onClose={() => store.dispatch(ToggleGuildSettingsDialog)} fullScreen>
			<AppBar style={{ position: 'relative' }}>
				<Toolbar>
					<IconButton edge="start" color="inherit" onClick={() => store.dispatch(ToggleGuildSettingsDialog)}>
						<CloseIcon />
					</IconButton>
					<Typography variant="h6">Guild Settings</Typography>
				</Toolbar>
			</AppBar>
			<DialogContent>
				<div style={{ width: '33%' }}>
					<input
						type="file"
						id="file"
						multiple
						ref={guildIconUpload}
						style={{ display: 'none' }}
						onChange={onGuildIconSelected}
					/>
					<ButtonBase
						style={{ borderRadius: '50%' }}
						onClick={() => {
							if (guildIconUpload.current) {
								guildIconUpload.current.click();
							}
						}}
					>
						<Avatar className={classes.guildIcon} src={guildIcon}></Avatar>
					</ButtonBase>
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
					<Table>
						<TableHead>
							<TableRow>
								<TableCell>Join Code</TableCell>
								<TableCell>Amount Used</TableCell>
								<TableCell>Actions</TableCell>
							</TableRow>
						</TableHead>
						<TableBody>
							{Object.keys(invites).map(key => {
								return (
									<TableRow key={key}>
										<TableCell component="th" scope="row">
											{key}
										</TableCell>
										<TableCell component="th" scope="row">
											{invites[key]}
										</TableCell>
										<TableCell component="td" scope="row">
											<Tooltip title="Copy Invite Link">
												<IconButton
													onClick={() => {
														copy(`http://${process.env.REACT_APP_HARMONY_SERVER_HOST}/invite/${key}`);
														toast.info('Successfully copied to clipboard!');
													}}
												>
													<ShareIcon />
												</IconButton>
											</Tooltip>
											<Tooltip title="Delete Invite Link">
												<IconButton onClick={() => deleteInviteLink(key)}>
													<DeleteIcon />
												</IconButton>
											</Tooltip>
										</TableCell>
									</TableRow>
								);
							})}
						</TableBody>
					</Table>
					<Button fullWidth onClick={createInviteLink}>
						<AddIcon />
					</Button>
				</div>
			</DialogContent>
		</Dialog>
	);
};
