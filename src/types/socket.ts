export interface IPacket {
	type: string;
	data: unknown;
}

export enum PacketTypes {
	Token = 'token',
	GetGuilds = 'getguilds',
	GetChannels = 'getchannels',
	GetMessages = 'getmessages',
}
