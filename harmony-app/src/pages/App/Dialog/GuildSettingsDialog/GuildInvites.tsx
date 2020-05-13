import React from 'react';
import { toast } from 'react-toastify';
import { useSelector } from 'react-redux';
import copy from 'copy-to-clipboard';
import { Table, TableHead, TableRow, TableCell, TableBody, Tooltip, IconButton, Button } from '@material-ui/core';
import AddIcon from '@material-ui/icons/Add';
import DeleteIcon from '@material-ui/icons/Delete';
import ShareIcon from '@material-ui/icons/Share';

import { harmonySocket } from '../../../../Root';
import { RootState } from '../../../../redux/store';

export const GuildInvites: React.FC = () => {
	const [currentGuild, invites] = useSelector((state: RootState) => [state.app.currentGuild, state.app.invites]);

	const deleteInvite = (invite: string) => {
		if (currentGuild) {
			harmonySocket.sendDeleteInvite(invite, currentGuild);
		}
	};

	const createInviteLink = () => {
		if (currentGuild) {
			harmonySocket.sendCreateInvite(currentGuild);
		}
	};

	return (
		<>
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
												copy(`${window.location.origin}/invite/${key}`);
												toast.info('Successfully copied to clipboard!');
											}}
										>
											<ShareIcon />
										</IconButton>
									</Tooltip>
									<Tooltip title="Delete Invite Link">
										<IconButton onClick={() => deleteInvite(key)}>
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
		</>
	);
};
