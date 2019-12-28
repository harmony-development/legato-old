import { EventEmitter } from 'fbemitter';
import { IPacket } from '../types/socket';

export default class HarmonySocket {
    conn: WebSocket;
    events: EventEmitter;

    constructor() {
        this.conn = new WebSocket(`ws://${process.env.REACT_APP_HARMONY_SERVER_HOST}/api/socket`);
        this.events = new EventEmitter();
        this.conn.addEventListener('open', () => this.events.emit('open'));
        this.conn.addEventListener('close', () => this.events.emit('close'));
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
    }

    connect = () => {
        this.conn = new WebSocket(`ws://${process.env.REACT_APP_HARMONY_SERVER_HOST}/api/socket`);
        this.conn.addEventListener('open', () => this.events.emit('open'));
        this.conn.addEventListener('close', () => this.events.emit('close'));
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

    login(email: string, password: string) {
        this.emitEvent('login', {
            email,
            password
        });
    }

    register(email: string, username: string, password: string) {
        this.emitEvent('register', {
            email,
            username,
            password
        });
    }

    getGuilds() {
        this.emitEvent('getguilds', {
            token: localStorage.getItem('token')
        });
    }

    getMessages(guildID: string) {
        this.emitEvent('getmessages', {
            token: localStorage.getItem('token'),
            guild: guildID
        });
    }

    sendMessage(guildID: string, channelID: string, text: string) {
        this.emitEvent('message', {
            token: localStorage.getItem('token'),
            guild: guildID,
            channel: channelID,
            message: text
        });
    }

    getChannels(guildID: string) {
        this.emitEvent('getchannels', {
            token: localStorage.getItem('token'),
            guild: guildID
        });
    }

    joinGuild(inviteCode: string) {
        this.emitEvent('joinguild', {
            token: localStorage.getItem('token'),
            invite: inviteCode
        });
    }

    createGuild(guildName: string) {
        this.emitEvent('createguild', {
            token: localStorage.getItem('token'),
            guildname: guildName
        });
    }

    leaveGuild(guildID: string) {
        this.emitEvent('leaveguild', {
            token: localStorage.getItem('token'),
            guild: guildID
        });
    }

    sendGuildNameUpdate(guildID: string, newname: string) {
        this.emitEvent('updateguildname', {
            token: localStorage.getItem('token'),
            guild: guildID,
            name: newname
        });
    }

    sendGuildPictureUpdate(guildID: string, newpicture: string) {
        this.emitEvent('updateguildpicture', {
            token: localStorage.getItem('token'),
            guild: guildID,
            picture: newpicture
        });
    }

    sendGetInvites(guildID: string) {
        this.emitEvent('getinvites', {
            token: localStorage.getItem('token'),
            guild: guildID
        });
    }

    sendAddChannel(guildID: string, channelname: string) {
        this.emitEvent('addchannel', {
            token: localStorage.getItem('token'),
            guild: guildID,
            channel: channelname
        });
    }

    sendDeleteChannel(guildID: string, channelID: string) {
        this.emitEvent('deletechannel', {
            token: localStorage.getItem('token'),
            guild: guildID,
            channel: channelID
        });
    }

    sendDeleteInvite(invite: string, guild: string) {
        this.emitEvent('deleteinvite', {
            token: localStorage.getItem('token'),
            invite,
            guild
        });
    }

    sendCreateInvite(guild: string) {
        this.emitEvent('createinvite', {
            token: localStorage.getItem('token'),
            guild
        });
    }

    sendGetUser(userid: string) {
        this.emitEvent('getuser', {
            token: localStorage.getItem('token'),
            userid
        });
    }
}
