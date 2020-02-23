import { EventEmitter } from 'fbemitter';
import axios from 'axios';

import { IPacket } from '../types/socket';

export default class HarmonySocket {
	token: string;
	conn: WebSocket;
	events: EventEmitter;
	userFetchQueue: {
		[key: string]: boolean;
	};

	constructor() {
		this.token = 'NONE';
		// eslint-disable-next-line no-undef
		this.conn = new WebSocket(`ws://${process.env.REACT_APP_HARMONY_SERVER_HOST}/api/socket`);
		// eslint-disable-next-line no-undef
		this.events = new EventEmitter();
		this.userFetchQueue = {};
		this.bindConnect();
	}

	connect = () => {
		// eslint-disable-next-line no-undef
		this.conn = new WebSocket(`ws://${process.env.REACT_APP_HARMONY_SERVER_HOST}/api/socket`);
		this.userFetchQueue = {};
		this.bindConnect();
	};

	bindConnect = () => {
		this.conn.addEventListener('open', () => this.events.emit('open'));
		this.conn.addEventListener('close', () => {
			setTimeout(this.connect, 3000);
			this.events.emit('close');
		});
		this.conn.addEventListener('error', () => this.events.emit('error'));
		this.conn.onmessage = (e: MessageEvent) => {
			const unprocessed = JSON.parse(e.data);
			if (typeof unprocessed['type'] === 'string' && typeof unprocessed['data'] === 'object') {
				const packet: IPacket = unprocessed;
				this.events.emit(packet.type, packet.data);
			} else {
				console.warn(`Unsupported packet received`);
				console.log(unprocessed);
			}
		};
	};

	exec(fn: () => void) {
		if (this.conn.readyState === WebSocket.OPEN) {
			fn();
		} else {
			this.events.addListener('open', () => {
				fn();
				this.events.removeCurrentListener();
			});
		}
	}

	emitEvent(type: string, data: unknown) {
		// choke all packets if connection is not working
		if (this.conn.readyState === WebSocket.OPEN) {
			this.conn.send(JSON.stringify({ type, data }));
		}
	}

	refreshToken() {
		this.token = localStorage.getItem('token') || 'NONE';
	}

	login(email: string, password: string) {
		this.emitEvent('login', {
			email,
			password,
		});
	}

	register(email: string, username: string, password: string) {
		this.emitEvent('register', {
			email,
			username,
			password,
		});
	}

	getGuilds() {
		this.emitEvent('getguilds', {
			token: this.token,
		});
	}

	getMessages(guildID: string, channelID: string) {
		this.emitEvent('getmessages', {
			token: this.token,
			guild: guildID,
			channel: channelID,
		});
	}

	getOldMessages(guildID: string, channelID: string, lastMessageID: string) {
		this.emitEvent('getmessages', {
			token: this.token,
			guild: guildID,
			channel: channelID,
			lastmessage: lastMessageID,
		});
	}

	sendMessage(guildID: string, channelID: string, text: string) {
		this.emitEvent('message', {
			token: this.token,
			guild: guildID,
			channel: channelID,
			message: text,
		});
	}

	sendMessageRest(guildID: string, channelID: string, text: string, attachment: File) {
		const uploadData = new FormData();
		uploadData.append('file', attachment);
		uploadData.append('token', this.token);
		uploadData.append('guild', guildID);
		uploadData.append('channel', channelID);
		uploadData.append('message', text);
		return axios.post(`http://${process.env.REACT_APP_HARMONY_SERVER_HOST}/api/rest/message`, uploadData, {});
	}

	sendDeleteMessage(guildID: string, channelID: string, messageID: string) {
		this.emitEvent('deletemessage', {
			token: this.token,
			guild: guildID,
			channel: channelID,
			message: messageID,
		});
	}

	getChannels(guildID: string) {
		this.emitEvent('getchannels', {
			token: this.token,
			guild: guildID,
		});
	}

	getMembers(guildID: string) {
		this.emitEvent('getmembers', {
			token: this.token,
			guild: guildID,
		});
	}

	getSelf() {
		this.emitEvent('getself', {
			token: this.token,
		});
	}

	joinGuild(inviteCode: string) {
		this.emitEvent('joinguild', {
			token: this.token,
			invite: inviteCode,
		});
	}

	createGuild(guildName: string) {
		this.emitEvent('createguild', {
			token: this.token,
			guildname: guildName,
		});
	}

	leaveGuild(guildID: string) {
		this.emitEvent('leaveguild', {
			token: this.token,
			guild: guildID,
		});
	}

	sendGuildNameUpdate(guildID: string, newname: string) {
		this.emitEvent('updateguildname', {
			token: this.token,
			guild: guildID,
			name: newname,
		});
	}

	sendGuildPictureUpdate(guildID: string, newpicture: string) {
		this.emitEvent('updateguildpicture', {
			token: this.token,
			guild: guildID,
			picture: newpicture,
		});
	}

	sendGetInvites(guildID: string) {
		this.emitEvent('getinvites', {
			token: this.token,
			guild: guildID,
		});
	}

	sendAddChannel(guildID: string, channelname: string) {
		this.emitEvent('addchannel', {
			token: this.token,
			guild: guildID,
			channel: channelname,
		});
	}

	sendDeleteChannel(guildID: string, channelID: string) {
		this.emitEvent('deletechannel', {
			token: this.token,
			guild: guildID,
			channel: channelID,
		});
	}

	sendDeleteInvite(invite: string, guild: string) {
		this.emitEvent('deleteinvite', {
			token: this.token,
			invite,
			guild,
		});
	}

	sendCreateInvite(guild: string) {
		this.emitEvent('createinvite', {
			token: this.token,
			guild,
		});
	}

	sendGetUser(userid: string) {
		if (!this.userFetchQueue[userid]) {
			this.userFetchQueue[userid] = true;
			this.emitEvent('getuser', {
				token: this.token,
				userid,
			});
		}
	}

	sendAvatarUpdate(avatar: string) {
		this.emitEvent('avatarupdate', {
			token: this.token,
			avatar,
		});
	}

	sendUsernameUpdate(username: string) {
		this.emitEvent('usernameupdate', {
			token: this.token,
			username,
		});
	}

	sendDeleteGuild(guildid: string) {
		this.emitEvent('deleteguild', {
			token: this.token,
			guild: guildid,
		});
	}

	sendPong() {
		this.emitEvent('ping', null);
	}
}
