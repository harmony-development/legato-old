export interface IPacket {
	type: string;
	data: unknown;
}

export interface IGuildsList {
	[guildid: string]: {
		guildname: string;
		picture: string;
		owner: boolean;
	};
}

export enum PacketTypes {
	Token = 'token',
	GetGuilds = 'getguilds',
	GetChannels = 'getchannels',
	GetMessages = 'getmessages',
}
